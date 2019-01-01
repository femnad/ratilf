package ratfilter

import (
	"fmt"
	"github.com/femnad/mare"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// Window type for representing Ratpoison windows
type Window struct {
	Number       int
	Class        string
	Status       WindowStatus
	Title        string
	LastAccessed int
}

func (w Window) String() string {
	return fmt.Sprintf("%d %s %s", w.Number, w.Status, w.Title)
}

// WindowList List of Ratpoison windows
type WindowList []Window

// WindowStatus window status
type WindowStatus string

const (
	windowFormat = "%n %l %c %s %t"
)

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	mare.PanicIfErr(err)
	return int(i)
}

func RunRatpoisonWindowCommand() WindowList {
	argument := fmt.Sprintf("windows %s", windowFormat)
	cmd := exec.Command("ratpoison", []string{"-c", argument}...)
	output, err := cmd.Output()
	mare.PanicIfErr(err)
	var windows WindowList
	outputString := strings.TrimSpace(string(output))
	for _, line := range strings.Split(outputString, "\n") {
		fields := strings.Fields(line)
		windowNumber := ParseInt(fields[0])
		lastAccessed := ParseInt(fields[1])
		windowStatus := WindowStatus(fields[3])
		windowTitle := strings.Join(fields[4:], " ")
		window := Window{Number: windowNumber, LastAccessed: lastAccessed, Class: fields[2], Status: windowStatus,
			Title: windowTitle}
		windows = append(windows, window)
	}
	return windows
}

func getWindowsWithClass(class string) WindowList {
	windows := RunRatpoisonWindowCommand()
	var matchedWindows WindowList
	for _, window := range windows {
		if window.Class == class {
			matchedWindows = append(matchedWindows, window)
		}
	}
	return matchedWindows
}

func GetWindows() WindowList {
	return RunRatpoisonWindowCommand()
}

func (w WindowList) Output() string {
	var sb strings.Builder
	for _, window := range w {
		sb.WriteString(window.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (w WindowList) SortByLastAccessAsc() WindowList {
	sort.Sort(AscByLastAccess(w))
	return w
}

type AscByLastAccess WindowList

func (w AscByLastAccess) Len() int {
	return len(w)
}

func (w AscByLastAccess) Less(i, j int) bool {
	return w[i].LastAccessed < w[j].LastAccessed
}

func (w AscByLastAccess) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type DescByLastAccess WindowList

func (w DescByLastAccess) Len() int {
	return len(w)
}

func (w DescByLastAccess) Less(i, j int) bool {
	return w[i].LastAccessed > w[j].LastAccessed
}

func (w DescByLastAccess) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func GetWindowsOfClass(class string) []Window {
	return getWindowsWithClass(class)
}

func selectWindow(windowNumber int) {
	selection := fmt.Sprintf("select %d", windowNumber)
	cmd := exec.Command("ratpoison", []string{"-c", selection}...)
	err := cmd.Run()
	mare.PanicIfErr(err)
}

func parseWindowNumber(windowLine string) int {
	return ParseInt(strings.Fields(windowLine)[0])
}

func getWindowWithNumber(number int) (*Window, error) {
	windows := RunRatpoisonWindowCommand()
	for _, window := range windows {
		if window.Number == number {
			return &window, nil
		}
	}
	return nil, fmt.Errorf("unable to find window with number %d", number)
}

func FocusWindowWithNumber(window *Window) {
	selectWindow(window.Number)
}

func FocusMatchingWindow(windowLine string) {
	windowNumber := parseWindowNumber(windowLine)
	window, err := getWindowWithNumber(windowNumber)
	mare.PanicIfErr(err)
	FocusWindowWithNumber(window)
}
