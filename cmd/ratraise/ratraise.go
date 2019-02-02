package main

import (
	"flag"
	"github.com/femnad/ratilf/pkg/ratfilter"
	"github.com/femnad/ratilf/pkg/run"
	"strings"
)

func ensureExecutable(executable string) {
	if executable == "" {
		panic("No executable provided")
	}
}

func getClassFromCommand(command string) string {
	exe := strings.Fields(command)[0]
	return strings.Title(exe)
}

func runOrRaiseWindow(command string, class string) {
	if class == "" {
		class = getClassFromCommand(command)
	}
	windows := ratfilter.GetWindowsOfClass(class).SortByLastAccessDesc()
	if len(windows) > 0 {
		ratfilter.FocusWindowWithNumber(&windows[0])
	} else {
		run.Command(command)
	}
}

func main() {
	executable := flag.String("exec", "", "executable of window to run or raise")
	class := flag.String("class", "", "class of window to run or raise")
	flag.Parse()
	ensureExecutable(*executable)
	runOrRaiseWindow(*executable, *class)
}
