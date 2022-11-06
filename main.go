package main

import (
	"fmt"
	commandExecuter "github.com/neptyune/beluga/utils"
)

func main() {
	// Here check for host os and any other issues
	//tui.StartTea()
	fmt.Println(commandExecuter.VolumeCreate())

}
