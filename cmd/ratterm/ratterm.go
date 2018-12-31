package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/femnad/ratilf/pkg/ratfilter"
)

func printWindowsOfTerminalClass(terminalClass string) {
	windows := ratfilter.GetWindowsOfClass(terminalClass)
	sort.Sort(ratfilter.AscByLastAccess(windows))
	for _, window := range windows {
		fmt.Println(window)
	}
	os.Exit(1)
}

func main() {
	class := flag.String("class", "", "class of the terminal emulator")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		printWindowsOfTerminalClass(*class)
	} else {
		ratfilter.FocusMatchingWindow(args[0])
	}
}