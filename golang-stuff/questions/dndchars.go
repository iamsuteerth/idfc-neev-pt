package questions

import (
	"math"
	"math/rand"
	"sort"
)

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
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	diceRolls := make([]int, 4)
	for i := 0; i < 4; i++ {
		diceRolls[i] = rand.Intn(6) + 1
	}
	sort.Ints(diceRolls)
	return diceRolls[1] + diceRolls[2] + diceRolls[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	strength := Ability()
	dexterity := Ability()
	constitution := Ability()
	intelligence := Ability()
	wisdom := Ability()
	charisma := Ability()

	hitpoints := 10 + Modifier(constitution)

	return Character{
		Strength:     strength,
		Dexterity:    dexterity,
		Constitution: constitution,
		Intelligence: intelligence,
		Wisdom:       wisdom,
		Charisma:     charisma,
		Hitpoints:    hitpoints,
	}
}
