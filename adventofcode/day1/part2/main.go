package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

	every_elves_individual_calories_total := make([]int, len(elf))

	for i := 0; i < len(elf); i++ {
		elf_calories := 0
		elf_foods := strings.Split(elf[i], "\n")
		for j := 0; j < len(elf_foods); j++ {
			x, _ := strconv.ParseInt(elf_foods[j], 10, 64)
			elf_calories += int(x)
		}
		every_elves_individual_calories_total[i] = elf_calories
		if elf_calories > max_calories {
			max_calories = elf_calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(every_elves_individual_calories_total)))
	total := 0
	for i := 0; i < 3; i++ {
		total += every_elves_individual_calories_total[i]
	}

	fmt.Println("The total calories of the 3 elves with the most calories is: ", total)
}
