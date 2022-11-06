package utils

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	//"strings"
)

func VolumePrune() string {
	return CreateTerminalOutput("volume", "prune", "-f")
}

func VolumeCreate() string { //haven't tested because idk what removing volumes is
	return CreateTerminalOutput("volume", "create")
}

func VolumeInspect() string {
	return CreateTerminalOutput("volume", "inspect", JSONFormat)
}

func VolumeList() string {
	return CreateTerminalOutput("volume", "list")
}

func ContainerListAsString() string {
	return CreateTerminalOutput("container", "list", "-a")

}

func ContainerListAsSlice() [][]string {
	matrix := [][]string{}
	re := regexp.MustCompile(" +")
	containersAsArray := strings.Split(ContainerListAsString(), "\n")
	for _, line := range containersAsArray {
		split := re.Split(line, -1)
		matrix = append(matrix, split)
	}
	//fmt.Println(matrix)
	//for i, slice := range matrix {
	//	matrix[i] = removeFromSlice(slice, 2)
	//}
	return matrix
}

func ImagesListAsSlice() []string {
	imagesAsString := CreateTerminalOutput("images", "list")
	re := regexp.MustCompile(" +")
	split := re.Split(imagesAsString, -1)
	fmt.Println(split)
	return split

}

func ImagesSearch() {
	//todo implement
}

func removeFromSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// retrieves stats
func StatsCommand() string {
	return CreateTerminalOutput("stats", "--no-stream", "--format", JSONFormat)
}

func GetDockerInfo() string {
	return CreateTerminalOutput("info", "--format", JSONFormat)
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

func longestWord(slice1 []string) int {
	largestLength := 0
	for _, word := range slice1 {
		if len(word) > largestLength {
			largestLength = len(word)
		}
	}
	return largestLength
}
