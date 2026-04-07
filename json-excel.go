package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Label string  `json:"label"`
	Conf  float64 `json:"conf"`
}

func main() {

	file, err := os.Open("test_0.json")
	if err != nil {
		fmt.Println("Ошибка при открытии файла", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var config Config
		err := json.Unmarshal([]byte(scanner.Text()), &config)
		if err != nil {
			fmt.Println("Ошибка при разборе файла", err)
			continue
		}
		fmt.Println("Название:", config.Label)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
