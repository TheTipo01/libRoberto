package libroberto

import (
	"os"
	"testing"
	"time"
)

func TestGenAudioMp3(t *testing.T) {
	uuid := GenAudio("Ciao, mondo!", "ogg", time.Second*30)
	if _, err := os.Stat("./temp/" + uuid + ".ogg"); os.IsNotExist(err) {
		t.Error("File not generated")
	} else {
		_ = os.Remove("./temp/" + uuid + ".ogg")
	}
}
