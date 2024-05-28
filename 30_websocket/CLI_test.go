package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/riyoth/learn-go-with-tests-exercises/28_command_line"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartedWith  int
	StartCalled  bool
	FinishedWith string
}

func (p *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	p.StartCalled = true
	p.StartedWith = numberOfPlayers
}

func (p *GameSpy) Finish(winner string) {
	p.FinishedWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of player and start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
	})
	t.Run("record chris win from user input", func(t *testing.T) {
		in := userSends("1", "Chris wins")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameFinishWith(t, game, "Chris")
	})
	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := userSends("1", "Cleo wins")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameFinishWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Pies")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("Game should not have started")
		}

		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrMsg

		assertMessagesSentToUser(t, stdout, wantPrompt)
	})
}

func assertGameStartedWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	if game.StartedWith != want {
		t.Errorf("wanted Start called with %d but got %d", want, game.StartedWith)
	}
}
func assertGameFinishWith(t testing.TB, game *GameSpy, want string) {
	t.Helper()
	if game.FinishedWith != want {
		t.Errorf("wanted Finish called with %v but got %v", want, game.FinishedWith)
	}
}
func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
