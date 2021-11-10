package libroberto

import (
	"embed"
	"math/rand"
	"strings"
)

//go:embed parole.txt
var adjectivesFile embed.FS

var (
	// Gods holds, well, the gods
	Gods = []string{"Dio", "Gesù", "Madonna"}
	// Adjectives holds a list of adjectives in italian
	Adjectives []string
)

func init() {
	initializeAdjectives()
}

// Reads Adjectives
func initializeAdjectives() {
	foo, _ := adjectivesFile.ReadFile("parole.txt")
	Adjectives = strings.Split(strings.ReplaceAll(string(foo), "\r\n", "\n"), "\n")
}

// Bestemmia generates a bestemmia
func Bestemmia() string {
	s1 := Gods[rand.Intn(len(Gods))]

	s := s1 + " " + Adjectives[rand.Intn(len(Adjectives))]

	if s1 == Gods[2] {
		s = s[:len(s)-1] + "a"
	}

	return s
}
