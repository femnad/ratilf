package main

import (
	"flag"
	"fmt"
	"github.com/femnad/ratilf/pkg/ratfilter"
	"os"
)

func listWindows() {
	windows := ratfilter.GetWindows().SortByLastAccessAsc()
	fmt.Print(windows.Output())
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		listWindows()
	} else {
		ratfilter.FocusMatchingWindow(args[0])
		os.Exit(1)
	}
}
