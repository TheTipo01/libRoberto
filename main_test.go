package libroberto

import (
	"os"
	"testing"
	"time"
)

func TestGenAudioMp3(t *testing.T) {
	uuid := GenAudioMp3("Ciao, mondo!", time.Second*30)
	if _, err := os.Stat("./temp/" + uuid + ".mp3"); os.IsNotExist(err) {
		t.Error("File not generated")
	} else {
		_ = os.Remove("./temp/" + uuid + ".mp3")
	}
}
