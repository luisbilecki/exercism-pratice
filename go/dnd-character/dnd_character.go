package dndcharacter

import (
	"math"
	"math/rand"
)

const DEFAULT_HIT_POINTS = 10
const DEFAULT_ATTEMPTS = 4
const ABILITY_MAX = 15
const ABILITY_MIN = 3

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score - 10)) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(ABILITY_MAX) + ABILITY_MIN
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	char.Hitpoints = DEFAULT_HIT_POINTS + Modifier(char.Constitution)
	return char
}
