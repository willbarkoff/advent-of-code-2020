package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort
on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there
have a little picture of a starfish; the locals just call them stars. None of the currency
exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by
the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent
calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star.
Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle
input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two
numbers together.
*/

const goal = 2020

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	previousNumbers := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		for _, v := range previousNumbers {
			if v+num == goal {
				fmt.Println(v * num)
				return
			}
		}

		previousNumbers = append(previousNumbers, num)
	}
}
