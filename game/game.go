package game

import "fmt"

// Game ...
type Game struct {
	numOfPlayers  int
	roundIsActive bool
	wordToGuess   string
	currentArtist *Player
}

// Join ...
func (g *Game) Join() (p *Player) {
	p = &Player{}
	p.attachedGame = g
	g.numOfPlayers++
	return p
}

// StartNewRound ...
func (g *Game) StartNewRound(wordToGuess string, newArtist *Player) error {
	if g.numOfPlayers < 2 {
		return fmt.Errorf("a round cannot start without at least two players in the game")
	}
	g.wordToGuess = wordToGuess
	g.currentArtist = newArtist
	g.roundIsActive = true
	return nil
}

// EndRound ...
func (g *Game) EndRound() {
	g.roundIsActive = false
}
