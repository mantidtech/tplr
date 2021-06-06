package functions

import (
	"os"
	"text/template"

	"golang.org/x/sys/unix"
)

// MiscellaneousFunctions are functions that don't have a more specific home
func MiscellaneousFunctions() template.FuncMap {
	return template.FuncMap{
		"terminalWidth": TerminalWidth,
	}
}

// TerminalWidth returns the number of columns that the terminal currently has,
// or 0 if the program isn't run in one, or it can't otherwise be determined
func TerminalWidth() int {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0
	}
	return int(ws.Col)
}
