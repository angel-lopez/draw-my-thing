package game

import (
	"testing"
)

func TestGameInitialization(t *testing.T) {
	game := Game{}
	err := game.StartNewRound("", &Player{})
	if err == nil {
		t.Error("a round should not have started without at least two players in the game")
	}
	_ = game.Join()
	_ = game.Join()
	err = game.StartNewRound("", &Player{})
	if err != nil {
		t.Errorf("a round should have started with at least two players in the game")
	}
}

func TestGuessRestrictions(t *testing.T) {
	game := Game{}
	player := game.Join()
	wordToGuess := "cat"
	_, err := player.Guess("preemptive_guess")
	if err == nil {
		t.Errorf("guessing should not be allowed without starting a round")
	}
	_ = game.Join()
	game.StartNewRound(wordToGuess, &Player{})
	_, err = player.Guess(wordToGuess)
	if err != nil {
		t.Errorf("guessing should be allowed after starting round")
	}
	_, err = player.Guess("unnecessary_guess")
	if err == nil {
		t.Error("player should not be able to guess again in a round where they have already guessed correctly")
	}
	game.EndRound()
	_, err = player.Guess("")
	if err == nil {
		t.Error("guessing should not be allowed after ending a round")
	}
}

func TestArtistGuessRestriction(t *testing.T) {
	game := Game{}
	artist := game.Join()
	_ = game.Join()
	game.StartNewRound("", artist)
	_, err := artist.Guess("")
	if err == nil {
		t.Error("the current artist should not be allowed to guess")
	}
}

func TestGuessCorrectness(t *testing.T) {
	wordToGuess := "cat"
	game := Game{}
	player := game.Join()
	_ = game.Join()
	game.StartNewRound(wordToGuess, &Player{})
	wasCorrect, _ := player.Guess("incorrect_guess")
	if wasCorrect {
		t.Error("players guess should not have been correct")
	}
	wasCorrect, _ = player.Guess(wordToGuess)
	if !wasCorrect {
		t.Errorf("players guess should have been correct")
	}
}

func TestGuessOutcomes(t *testing.T) {
	game := Game{}
	player1 := game.Join()
	player2 := game.Join()
	wordToGuess := "cat"
	game.StartNewRound(wordToGuess, &Player{})
	_, _ = player1.Guess(wordToGuess)
	_, _ = player2.Guess("incorrect_guess")

	if player1.GetScore() != 100 {
		t.Error("player should have 100 score after one correct guess")
	}

	if player2.GetScore() != 0 {
		t.Errorf("player should have 0 score after never guessing correctly")
	}
}

func TestIsArtist(t *testing.T) {
	game := Game{}
	player1 := game.Join()
	player2 := game.Join()
	game.StartNewRound("", player1)
	if !player1.IsArtist() {
		t.Error("player set as artist not showing as artist")
	}
	if player2.IsArtist() {
		t.Error("player not set as artist should not show up as artist")
	}
}

func TestArtistScoring(t *testing.T) {
	game := Game{}
	artist := game.Join()
	player1 := game.Join()
	player2 := game.Join()
	wordToGuess := "cat"
	game.StartNewRound(wordToGuess, artist)
	player1.Guess(wordToGuess)
	player2.Guess(wordToGuess)
	if artist.GetScore() != 20 {
		t.Errorf("current artist should receive 10 score for each correct guess during the round, expected %v, got %v", 20, artist.GetScore())
	}
}
