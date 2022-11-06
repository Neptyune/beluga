package main

import (
	"fmt"
	commandExecuter "github.com/neptyune/beluga/utils"
)

func main() {
	//tui.StartTea()
	fmt.Println(commandExecuter.ContainerListAsSlice())

}
