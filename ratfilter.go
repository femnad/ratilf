package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/femnad/mare"
)

// Window Struct for representing Ratpoison windows
type Window struct {
	number       int
	class        string
	status       string
	title        string
	lastAccessed int
}

// WindowList List of Ratpoison windows
type WindowList []Window

const (
	windowFormat = "%n %l %c %s %t"
)

type byLastAccess WindowList

func (w byLastAccess) Len() int {
	return len(w)
}

func (w byLastAccess) Less(i, j int) bool {
	return w[i].lastAccessed < w[j].lastAccessed
}

func (w byLastAccess) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func parseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	mare.PanicIfErr(err)
	return int(i)
}

func runRatpoisonWindowCommand() WindowList {
	argument := fmt.Sprintf("windows %s", windowFormat)
	cmd := exec.Command("ratpoison", []string{"-c", argument}...)
	output, err := cmd.Output()
	mare.PanicIfErr(err)
	var windows WindowList
	outputString := strings.TrimSpace(string(output))
	for _, line := range strings.Split(outputString, "\n") {
		fields := strings.Fields(line)
		windowNumber := parseInt(fields[0])
		lastAccessed := parseInt(fields[1])
		windowTitle := strings.Join(fields[3:], " ")
		window := Window{number: windowNumber, lastAccessed: lastAccessed, class: fields[2], status: fields[3],
			title: windowTitle}
		windows = append(windows, window)
	}
	return windows
}

func getWindowsWithClass(class string) WindowList {
	windows := runRatpoisonWindowCommand()
	var matchedWindows WindowList
	for _, window := range windows {
		if window.class == class {
			matchedWindows = append(matchedWindows, window)
		}
	}
	return matchedWindows
}

func getCurrentWindowClass() (string, error) {
	windows := runRatpoisonWindowCommand()
	for _, window := range windows {
		if window.status == "*" {
			return window.class, nil
		}
	}
	return "", errors.New("unable to find current window")
}

func getWindowWithNumber(number int) (*Window, error) {
	windows := runRatpoisonWindowCommand()
	for _, window := range windows {
		if window.number == number {
			return &window, nil
		}
	}
	return nil, fmt.Errorf("unable to find window with number %d", number)
}

func printWindowsOfSameClass() {
	currentWindowClass, err := getCurrentWindowClass()
	mare.PanicIfErr(err)
	windows := getWindowsWithClass(currentWindowClass)
	sort.Sort(byLastAccess(windows))
	for _, window := range windows {
		fmt.Printf("%d %s %s\n", window.number, window.status, window.title)
	}
}

func parseWindowNumber(windowLine string) int {
	return parseInt(strings.Fields(windowLine)[0])
}

func selectWindow(windowNumber int) {
	selection := fmt.Sprintf("select %d", windowNumber)
	cmd := exec.Command("ratpoison", []string{"-c", selection}...)
	err := cmd.Run()
	mare.PanicIfErr(err)
}

func focusMatchingWindow(windowLine string) {
	windowNumber := parseWindowNumber(windowLine)
	window, err := getWindowWithNumber(windowNumber)
	mare.PanicIfErr(err)
	selectWindow(window.number)
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
