package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	var err error

	htmlFilePath := "./dash/index.html"

	// Ensure the path is absolute, because some systems require an absolute path to open the file
	absolutePath, pathErr := filepath.Abs(htmlFilePath)
	if pathErr != nil {
		fmt.Println("Error getting absolute path:", pathErr)
		return
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", absolutePath).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", absolutePath).Start()
	case "darwin":
		err = exec.Command("open", absolutePath).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Print(err)
	}
}
