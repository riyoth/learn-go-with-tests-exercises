package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q,  want %q", store.winCalls[0], winner)

	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func AssertScheduledAlert(t testing.TB, got ScheduledAlert, want ScheduledAlert) {
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

type GameSpy struct {
	StartedWith int
	StartCalled bool
	BlindAlert  []byte

	FinishedCalled bool
	FinishedWith   string
}

func (p *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	p.StartCalled = true
	p.StartedWith = numberOfPlayers
	alertsDestination.Write(p.BlindAlert)
}

func (p *GameSpy) Finish(winner string) {
	p.FinishedWith = winner
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == want
	})
	if !passed {
		t.Errorf("wanted Started called with %v but got %v", want, game.StartedWith)
	}
}
func assertGameFinishWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == want
	})
	if !passed {
		t.Errorf("wanted Finish called with %v but got %v", want, game.FinishedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
