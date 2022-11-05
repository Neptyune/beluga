package Utils

import (
	"log"
	"os/exec"
)

func CreateTerminalOutput(command string, args ...string) string {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
