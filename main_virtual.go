package libRoberto

import (
	"crypto/sha1"
	"encoding/base32"
	"github.com/blang/vfs/memfs"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	FS = memfs.Create()
)

// GenAudioVirtual generates a mp3 file (in a virtual filesystem) from a string, returning its UUID (aka SHA1 hash of the text)
// Remember to delete the file when you are done, as it sits in RAM!
func GenAudioMp3Virtual(text string, timeOut time.Duration) string {
	h := sha1.New()
	h.Write([]byte(text))
	uuid := strings.ToUpper(base32.HexEncoding.EncodeToString(h.Sum(nil)))

	genMp3Memfs(text, uuid, timeOut)

	return uuid
}

func genMp3Memfs(text string, uuid string, timeOut time.Duration) {
	var tts, ffmpeg *exec.Cmd
	// Create a channel so that we can wait a maximum of 30 second before killing the processes
	c := make(chan int)
	// File
	f, _ := FS.OpenFile(uuid+".mp3", os.O_CREATE, 0666)
	defer f.Close()

	go func() {
		tts = exec.Command("balcon", "-i", "-o", "-enc", "utf8", "-n", "Roberto")
		tts.Stdin = strings.NewReader(text)
		ttsOut, _ := tts.StdoutPipe()
		_ = tts.Start()

		ffmpeg = exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "2", "-f", "pipe:1")
		ffmpeg.Stdin = ttsOut
		ffmpeg.Stdout = f

		_ = ffmpeg.Run()
		_ = tts.Wait()
		c <- 1
	}()

	timer := time.NewTimer(timeOut)
	// If we get a response from the channel in a reasonable time, we just exit normally
	select {
	case <-c:
		timer.Stop()
		break
	case <-timer.C:
		// After 30 second, we can assume the process failed in some manner
		if tts != nil && tts.Process != nil {
			_ = tts.Process.Kill()
		}

		if ffmpeg != nil && ffmpeg.Process != nil {
			_ = ffmpeg.Process.Kill()
		}
	}
}

// GenAudioVirtual generates a dca file (in a virtual filesystem) from a string, returning its UUID (aka SHA1 hash of the text)
// Remember to delete the file when you are done, as it sits in RAM!
func GenAudioDcaVirtual(text string, timeOut time.Duration) string {
	h := sha1.New()
	h.Write([]byte(text))
	uuid := strings.ToUpper(base32.HexEncoding.EncodeToString(h.Sum(nil)))

	genDCAMemfs(text, uuid, timeOut)

	return uuid
}

// Generates a virtual DCA file
func genDCAMemfs(text string, uuid string, timeOut time.Duration) {
	var tts, ffmpeg, dca *exec.Cmd
	// Create a channel so that we can wait a maximum of 30 second before killing the processes
	c := make(chan int)
	// File
	f, _ := FS.OpenFile(uuid+".dca", os.O_CREATE, 0666)
	defer f.Close()

	go func() {
		// Starts TTS generation
		tts = exec.Command("balcon", "-i", "-o", "-enc", "utf8", "-n", "Roberto")
		tts.Stdin = strings.NewReader(text)
		ttsOut, _ := tts.StdoutPipe()
		_ = tts.Start()

		// Starts middle step
		ffmpeg = exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "2", "pipe:1")
		ffmpeg.Stdin = ttsOut
		ffmpegOut, _ := ffmpeg.StdoutPipe()

		dca = exec.Command("dca")
		dca.Stdin = ffmpegOut
		dca.Stdout = f

		_ = ffmpeg.Run()
		_ = tts.Wait()
		c <- 1
	}()

	timer := time.NewTimer(timeOut)
	// If we get a response from the channel in a reasonable time, we just exit normally
	select {
	case <-c:
		timer.Stop()
		break
	case <-timer.C:
		// After 30 second, we can assume the process failed in some manner
		if tts != nil && tts.Process != nil {
			_ = tts.Process.Kill()
		}

		if ffmpeg != nil && ffmpeg.Process != nil {
			_ = ffmpeg.Process.Kill()
		}

		if dca != nil && dca.Process != nil {
			_ = dca.Process.Kill()
		}
	}
}
