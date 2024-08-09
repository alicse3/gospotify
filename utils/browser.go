package utils

import (
	"errors"
	"os/exec"
	"runtime"

	"github.com/alicse3/gospotify/consts"
)

// CommandExecutor interface defines the methods for executing system commands.
type CommandExecutor interface {
	Run(command string, args ...string) error
}

// DefaultCommandExectutor is a struct which implements CommandExecutor interface.
type DefaultCommandExectutor struct{}

// Run executes a system command with the provided arguments.
// It returns an error if the command execution fails.
func (dce *DefaultCommandExectutor) Run(command string, args ...string) error {
	return exec.Command(command, args...).Run()
}

// BrowserOpener defines an interface for opening a browser.
type BrowserOpener interface {
	Open(url string) error
}

// DefaultBrowserOpener is a struct which implements BrowserOpener interface.
type DefaultBrowserOpener struct {
	commandExecutor CommandExecutor
}

// NewDefaultBrowserOpener creates a new DefaultBrowserOpener with the provided executor.
func NewDefaultBrowserOpener(commandExecutor CommandExecutor) *DefaultBrowserOpener {
	return &DefaultBrowserOpener{commandExecutor}
}

// Open opens a URL in the default browser.
func (boi *DefaultBrowserOpener) Open(url string) error {
	var err error

	switch runtime.GOOS {
	case "darwin":
		err = boi.commandExecutor.Run("open", url)
	case "windows":
		err = boi.commandExecutor.Run("cmd", "/c", "start", url)
	case "linux":
		err = boi.commandExecutor.Run("xdg-open", url)
	default:
		err = errors.New(consts.MsgUnsupportedPlatform)
	}

	if err != nil {
		return err
	}

	return nil
}
