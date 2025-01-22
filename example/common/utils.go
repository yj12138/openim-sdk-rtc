package common

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func TimeStampToStr(timestampSec int64) string {
	t := time.Unix(timestampSec, 0)
	return t.Format(time.DateTime)
}

func OpenBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url) // Linux
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", url) // Windows
	case "darwin":
		cmd = exec.Command("open", url) // macOS
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}
