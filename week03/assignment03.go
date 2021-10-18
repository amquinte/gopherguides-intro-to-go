package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Length int
	Name   string
	plays  int
	views  int
	rating float32
}

type Theater struct {
	movies []Movie
}

func main() {

}

func (m *Movie) Rate(r float32) error {

	if m.plays == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	}

	m.rating = m.rating + r
	return nil
}

func (m *Movie) Play(v int) {

	m.views = m.views + v //Increase number of viewers for that movie
	m.plays += 1          //Increases number of plays for that movie by one
}

func (m Movie) Viewers() int {
	return m.views
}

func (m Movie) Plays() int {
	return m.plays
}

func (m Movie) Rating() float64 {
	r := m.rating / float32(m.plays)
	return float64(r)
}

func (m Movie) String() string {
	movieLen := strconv.Itoa(m.Length)
	movieRating := strconv.FormatFloat(float64(m.rating), 'f', 1, 64)
	movieString := m.Name + " (" + movieLen + "m) " + movieRating + "%"
	return movieString
}

type CritiqueFn func(m *Movie) (float32, error)

func (t *Theater) Play(v int, m ...Movie) error {
	if m == nil {
		return fmt.Errorf("no movies to play")
	}
	return nil
}
