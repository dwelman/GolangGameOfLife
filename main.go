package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func fill_map(array [][]bool) [][]bool {
	for y := range array {
		for x := range array[y] {
			if rand.Intn(100) > 40 {
				array[y][x] = true
			} else {
				array[y][x] = false
			}
		}
	}
	return array
}

func print_map(array [][]bool) {
	for y := range array {
		for x := range array[y] {
			if array[y][x] == false {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}

func process_map(a [][]bool) [][]bool {
	height := len(a)
	width := len(a[0])

	//Clone array for return
	array := make([][]bool, height)
	for i := range a {
		array[i] = make([]bool, width)
	}

	for y := range a {
		for x := range a[0] {
			//Check rules for game of life on original array, record results on new array

			//Create the ranges we need to check neighbour cells
			x_left := x - 1
			if x_left < 0 {
				x_left = 0
			}
			x_right := x + 1
			if x_right >= width {
				x_right = width - 1
			}
			y_up := y - 1
			if y_up < 0 {
				y_up = 0
			}
			y_down := y + 1
			if y_down >= height {
				y_down = height - 1
			}

			neighbour_count := 0
			for yy := y_up; yy <= y_down; yy++ {
				for xx := x_left; xx <= x_right; xx++ {
					if y == yy && x == xx {
						continue
					}
					if a[yy][xx] == true {
						neighbour_count++
					}
				}
			}

			if a[y][x] == true {
				if neighbour_count < 2 {
					//Dies from unerpopulation
					array[y][x] = false
				} else if neighbour_count == 2 || neighbour_count == 3 {
					//Lives on
					array[y][x] = true
				} else {
					//Dies from overpopulation
					array[y][x] = false
				}
			} else {
				if neighbour_count == 3 {
					//Lives by repopulation
					array[y][x] = true
				}
			}
		}
	}

	return array
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Must have 2 arguments for width and height, with an option third option of speed (default 1)")
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if width <= 4 {
		log.Fatal("Width must be greater than 4")
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if height <= 4 {
		log.Fatal("Height must be greater than 4")
	}
	speed := 1
	if len(os.Args) > 3 {
		speed, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
	}
	if speed <= 0 {
		log.Fatal("Speed must be greater than 0")
	}

	fmt.Println("width:", width)
	fmt.Println("height:", height)

	array := make([][]bool, height)
	for i := range array {
		array[i] = make([]bool, width)
	}

	array = fill_map(array)

	for true {
		array = process_map(array)
		fmt.Printf("\033[0;0H")
		print_map(array)
		time.Sleep(time.Second / time.Duration(speed))
	}
}
