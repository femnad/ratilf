package main

import (
	"flag"
	"github.com/femnad/mare"
	"github.com/femnad/ratilf/pkg/ratfilter"
	"os/exec"
	"sort"
	"strings"
)

func ensureExecutable(executable string) {
	if executable == "" {
		panic("No executable provided")
	}
}

func getClassFromExecutable(executable string) string {
	return strings.Title(executable)
}

func runOrRaiseWindow(executable string, class string) {
	if class == "" {
		class = getClassFromExecutable(executable)
	}
	windows := ratfilter.GetWindowsOfClass(class)
	sort.Sort(ratfilter.DescByLastAccess(windows))
	for _, window := range windows {
		ratfilter.FocusWindowWithNumber(&window)
		return
	}
	cmd := exec.Command(executable)
	err := cmd.Run()
	mare.PanicIfErr(err)
}

func main() {
	executable := flag.String("exec", "", "executable of window to run or raise")
	class := flag.String("class", "", "class of window to run or raise")
	flag.Parse()
	ensureExecutable(*executable)
	runOrRaiseWindow(*executable, *class)
}
