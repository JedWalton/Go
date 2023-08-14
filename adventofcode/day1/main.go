package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func read_file() []byte {
	filename := "adventofcode.com_2022_day_1_input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
	}
	return data

}
func main() {
	data := read_file()
	string_data := string(data)
	elf := strings.Split(string_data, "\n\n")
	max_calories := 0

	for i := 0; i < len(elf); i++ {
		elf_calories := 0
		elf_foods := strings.Split(elf[i], "\n")
		for j := 0; j < len(elf_foods); j++ {
			x, _ := strconv.ParseInt(elf_foods[j], 10, 64)
			elf_calories += int(x)
		}
		if elf_calories > max_calories {
			max_calories = elf_calories
		}
	}
	fmt.Println(max_calories)
}
