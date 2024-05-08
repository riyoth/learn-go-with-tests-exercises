package clockface

import (
	"math"
	"time"
)

const (
	secondsInClock     = 60
	secondsInHalfClock = secondsInClock / 2
	minutesInClock     = 60
	minutesInHalfClock = minutesInClock / 2
	hoursInClock       = 12
	hoursInHalfClock   = hoursInClock / 2
)

type Point struct {
	X float64
	Y float64
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func HoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}
func secondHandPoint(t time.Time) Point {
	point := angleToPoint(secondsInRadians(t))
	return point
}

func minuteHandPoint(t time.Time) Point {
	point := angleToPoint(minutesInRadians(t))
	return point
}
func hourHandPoint(t time.Time) Point {
	point := angleToPoint(HoursInRadians(t))
	return point
}
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}

	return p
}
