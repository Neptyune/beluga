package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func VolumePrune() string {
	return CreateTerminalOutput("volume", "prune", "-f")
}

func VolumeCreate() string { //haven't tested because idk what removing volumes is
	return CreateTerminalOutput("volume", "create")
}

func CreateTerminalOutput(args ...string) string {
	var out []byte
	var err error
	if isWindows() {
		winArgs := make([]string, 1, 4)
		winArgs = []string{"/C", "docker"}
		winArgs = append(winArgs, args...)
		out, err = exec.Command("cmd", winArgs...).Output()
	} else {
		out, err = exec.Command("docker", args...).Output()
	}
	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}
	return string(out)
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
