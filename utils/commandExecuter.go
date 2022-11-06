package utils

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	//"strings"
)

func VolumePrune() string {
	return CreateTerminalOutput("/C", "docker", "volume", "prune", "-f")
}

func VolumeCreate() string { //haven't tested because idk what removing volumes is
	return CreateTerminalOutput("volume", "create")
}

func VolumeInspect() string {
	return CreateTerminalOutput("/C", "docker", "volume", "inspect")
}

func VolumeList() string {
	return CreateTerminalOutput("/C", "docker", "volume", "list")
}

func ContainerListAsString() string {
	return CreateTerminalOutput("/C", "docker", "images", "list")

}

func ContainerListAsSlice() []string {
	containersAsString := ContainerListAsString()
	re := regexp.MustCompile(" +")
	split := re.Split(containersAsString, -1)
	fmt.Println(split)
	return split
}

func ImagesListAsSlice() []string {
	imagesAsString := CreateTerminalOutput("/C", "docker", "images", "list")
	re := regexp.MustCompile(" +")
	split := re.Split(imagesAsString, -1)
	fmt.Println(split)
	return split

}

func removeFromSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func ImagesSearch() {
	//todo implement
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
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return string(out)
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}
