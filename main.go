package libroberto

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

// Generates audio from a string. Checks if it already exist before generating it
func genMp3(text string, uuid string, timeOut time.Duration) {
	const audioExtension = ".mp3"
	_, err := os.Stat("./temp/" + uuid + audioExtension)

	if err != nil {
		var tts, ffmpeg *exec.Cmd
		// Create a channel so that we can wait a maximum of 30 second before killing the processes
		c := make(chan int)

		go func() {
			tts = exec.Command("balcon", "-i", "-o", "-enc", "utf8", "-n", "Roberto")
			tts.Stdin = strings.NewReader(text)
			ttsOut, _ := tts.StdoutPipe()
			_ = tts.Start()

			ffmpeg = exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "2", "-f", "mp3", "./temp/"+uuid+audioExtension)
			ffmpeg.Stdin = ttsOut
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
}

// GenAudioMp3 generates a mp3 file from a string, returning its UUID (aka SHA1 hash of the text)
func GenAudioMp3(text string, timeOut time.Duration) string {
	uuid := GenUUID(text)

	genMp3(text, uuid, timeOut)

	return uuid
}

// GenDCA returns a slice of exec.Cmd with commands to start. The stdout of the last element will contain the DCA stream
func GenDCA(text string) []*exec.Cmd {
	tts := exec.Command("balcon", "-i", "-o", "-enc", "utf8", "-n", "Roberto")
	tts.Stdin = strings.NewReader(text)
	ttsOut, _ := tts.StdoutPipe()

	ffmpeg := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "s16le", "-ar", "48000", "-ac", "pipe:1")
	ffmpeg.Stdin = ttsOut
	ffmpegOut, _ := ffmpeg.StdoutPipe()

	dca := exec.Command("dca")
	dca.Stdin = ffmpegOut

	return []*exec.Cmd{tts, ffmpeg, dca}
}
