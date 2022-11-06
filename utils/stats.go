package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

const JSONFormat string = "\"{{ json . }}\""

// structs
type Volume struct {
	CreatedAt  string
	Driver     string
	Labels     string
	Mountpoint string
	Name       string
	Options    string
	Scope      string
}

type LiveContainer struct {
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

type DockerInfo struct {
	Containers        int
	ContainersRunning int
	ContainersPaused  int
	ContainersStopped int
	Images            int
	ServerVersion     string
	KernelVersion     string
	OperatingSystem   string
	OSType            string
	Architecture      string
	NCPU              int
	MemTotal          int
}

// function to create DockerInfo struct
func DockerInfoStruct() DockerInfo {
	rawData := GetDockerInfo()
	data := TrimJSON(rawData, 2)
	InfoStruct := DockerInfo{}
	err := json.Unmarshal([]byte(data), &InfoStruct)
	if err != nil {
		fmt.Println(err)
	}
	return InfoStruct
}

// function to create a container
func GetLiveContainer() []LiveContainer {
	fmt.Println(StatsCommand())
	rawData := strings.Split(StatsCommand(), "\n")
	containers := []LiveContainer{}
	for i, data := range rawData {
		fmt.Println(data)
		if i == len(rawData)-1 {
			break
		} else {
			data = TrimJSON(data, 1)
		}

		container := LiveContainer{}
		err := json.Unmarshal([]byte(data), &container)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("Struct is: ", container)
		containers = append(containers, container)
	}
	return containers
}

func TrimJSON(rawData string, length int) string {
	data := strings.TrimLeft(rawData, "\"")
	data = data[:(len(data) - length)]
	//fmt.Println(data)
	return data
}

func GetVolume() {
	data := VolumeInspect()
	data = TrimJSON(data, 2)
	fmt.Println(data)
}
