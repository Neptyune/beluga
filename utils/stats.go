package Utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Container struct {
	BlockIO   string `json:"BlockIO"`
	CPUPerc   string `json:"CPUPerc"`
	Container string `json:"Container"`
	ID        string `json:"ID"`
	MemPerc   string `json:"MemPerc"`
	MemUsage  string `json:"MemUsage"`
	Name      string `json:"Name"`
	NetIO     string `json:"NetIO"`
	PIDs      string `json:"PIDs"`
}

// function to create a container
func CreateContainer() {
	rawData := StatsCommand()
	//instance of Container
	container1 := Container{}
	data := strings.TrimLeft(rawData, "\"")
	data = data[:(len(data) - 2)]
	//fmt.Println(data)
	err := json.Unmarshal([]byte(data), &container1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Struct is: ", container1)
}

// retrieves stats
func StatsCommand() string {
	fmt.Println("hello")
	format := "{{ json . }}"
	quotedFormat := fmt.Sprintf("%q", format)
	out := CreateTerminalOutput("/C", "docker", "stats", "--no-stream", "--format", quotedFormat)
	return out
}
