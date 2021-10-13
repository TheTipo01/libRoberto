package libRoberto

import (
	"embed"
	"github.com/bwmarrin/lit"
	"github.com/forPelevin/gomoji"
	"github.com/goccy/go-json"
	"strings"
)

//go:embed emoji.json
var emojiFile embed.FS

var (
	// Emoji string replacer, replacing every emoji with it's description
	emoji *strings.Replacer
)

func init() {
	emoji = emojiReplacer()
}

func emojiReplacer() *strings.Replacer {
	var (
		emojiJSON Emoji
		args      []string
	)

	// Load JSON file
	jsonFile, err := emojiFile.Open("emoji.json")
	if err != nil {
		lit.Error("Error opening file: %s", err)
		return nil
	}

	_ = json.NewDecoder(jsonFile).Decode(&emojiJSON)
	_ = jsonFile.Close()

	// Create the replacer
	for _, e := range emojiJSON {
		args = append(args, e.Emoji, e.Descrizione)
	}

	return strings.NewReplacer(args...)
}

func EmojiToDescription(str string) string {
	if gomoji.ContainsEmoji(str) {
		str = emoji.Replace(str)
	}

	return str
}
