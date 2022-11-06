package Utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func VolumePrune() string {
	return CreateTerminalOutput("/C", "docker", "volume", "prune", "-af")
}

func VolumeCreate() string { //haven't tested because idk what removing volumes is
	return CreateTerminalOutput("/C", "docker", "volume", "create")
}

func CreateTerminalOutput(args ...string) string {
	out, err := exec.Command(getTerminalType(), args...).Output()
	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}
	return string(out)
}

func getTerminalType() string {
	if runtime.GOOS == "windows" {
		return "cmd"
	} else {
		return "bash"
	}
}
