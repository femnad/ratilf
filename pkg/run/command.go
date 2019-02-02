package run

import (
	"os/exec"
	"strings"
	"github.com/femnad/mare"
)

func Command(command string) {
	tokens := strings.Fields(command)
	exe := tokens[0]
	args := tokens[1:]
	cmd := exec.Command(exe, args...)
	err := cmd.Start()
	mare.PanicIfErr(err)
}