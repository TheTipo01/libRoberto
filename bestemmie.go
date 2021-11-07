package libroberto

import (
	"embed"
	"math/rand"
	"strings"
)

//go:embed parole.txt
var adjectivesFile embed.FS

var (
	// Gods
	gods       = []string{"Dio", "Gesù", "Madonna"}
	adjectives []string
)

func init() {
	initializeAdjectives()
}

// Reads adjectives
func initializeAdjectives() {
	foo, _ := adjectivesFile.ReadFile("parole.txt")
	adjectives = strings.Split(string(foo), "\n")
}

// Bestemmia generates a bestemmia
func Bestemmia() string {
	s1 := gods[rand.Intn(len(gods))]

	s := s1 + " " + adjectives[rand.Intn(len(adjectives))]

	if s1 == gods[2] {
		s = s[:len(s)-2] + "a"
	}

	return s
}
