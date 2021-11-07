package libroberto

import (
	"crypto/sha1"
	"encoding/base32"
	"os/exec"
	"strings"
)

// GenUUID generates the UUID for the given text
func GenUUID(text string) string {
	h := sha1.New()
	h.Write([]byte(text))

	return strings.ToUpper(base32.HexEncoding.EncodeToString(h.Sum(nil)))
}

// CmdsStart starts all the exec.Cmd inside the slice
func CmdsStart(cmds []*exec.Cmd) {
	for _, cmd := range cmds {
		_ = cmd.Start()
	}
}

// CmdsWait waits for all the exec.Cmd inside the slice to finish processing, to free up resources
func CmdsWait(cmds []*exec.Cmd) {
	for _, cmd := range cmds {
		_ = cmd.Wait()
	}
}
