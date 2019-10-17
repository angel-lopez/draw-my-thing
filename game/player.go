package game

import "fmt"

// Player ...
type Player struct {
	attachedGame     *Game
	score            int
	guessedCorrectly bool
	isArtist         bool
}

// Guess ...
func (p *Player) Guess(word string) (isCorrectGuess bool, err error) {
	if !p.attachedGame.roundIsActive {
		return false, fmt.Errorf("guessing not allowed, there is no active round")
	}
	if p.guessedCorrectly {
		return false, fmt.Errorf("guessing not allowed, player has already guessed correctly for the current round")
	}
	if p.IsArtist() {
		return false, fmt.Errorf("the current artist is not allowed to guess")
	}
	isCorrectGuess = word == p.attachedGame.wordToGuess
	if isCorrectGuess {
		p.guessedCorrectly = true
		p.score += 100
		p.attachedGame.currentArtist.score += 10
	}
	return isCorrectGuess, nil
}

// GetScore ...
func (p *Player) GetScore() int {
	return p.score
}

// IsArtist ...
func (p *Player) IsArtist() bool {
	return p == p.attachedGame.currentArtist
}
