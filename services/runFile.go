package services

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunFile(fileName, extension string) error {
	dir := "file_manager"
	filePath := filepath.Join(dir, fileName+extension)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("File does not exist")
	}

	switch extension {
	case ".sh":
		return executeCommand("bash", filePath)
	case ".py":
		return executeCommand("python3", filePath)
	case ".js":
		return executeCommand("node", filePath)
	case ".go":
		return executeCommand("go", filePath)
	default:
		return openInEditor(filePath)
	}
}

func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func openInEditor(filePath string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("notepad", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	default:
		cmd = exec.Command("xdg-open", filePath)
	}
	return cmd.Run()
}
