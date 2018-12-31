package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/femnad/mare"
	"github.com/femnad/ratilf/pkg/ratfilter"
)


func getCurrentWindowClass() (string, error) {
	windows := ratfilter.RunRatpoisonWindowCommand()
	for _, window := range windows {
		if window.Status == "*" {
			return window.Class, nil
		}
	}
	return "", errors.New("unable to find current window")
}


func printWindowsOfSameClass() {
	currentWindowClass, err := getCurrentWindowClass()
	mare.PanicIfErr(err)
	windows := ratfilter.GetWindowsOfClass(currentWindowClass)
	sort.Sort(ratfilter.AscByLastAccess(windows))
	for _, window := range windows {
		fmt.Println(window)
	}
}

func focusMatchingWindow(windowLine string) {
	ratfilter.FocusMatchingWindow(windowLine)
	os.Exit(1)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		printWindowsOfSameClass()
	} else {
		focusMatchingWindow(args[0])
	}
}
