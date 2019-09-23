package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(e error) {
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
}

func loadfile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getAlienCount() (uint, error) {
	var number int
	fmt.Println("Number of Aliens : ")

	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		return 0, err
	}

	return uint(number), nil
}

func main() {
	wc := NewWorldMap()

	input, err := loadfile("input.txt")
	checkError(err)

	err = wc.AddCities(string(input))
	checkError(err)

	var alienCount uint
	alienCount, err = getAlienCount()
	checkError(err)

	err = wc.AddAliens(alienCount)
	checkError(err)

	fmt.Println("\nWorld at genesis!")
	fmt.Println(wc)

	fmt.Println()
	fmt.Println("Starting Simulation ***")
	fmt.Println()

	wc.Simulate(10000)
	fmt.Println("Starting Ended ***")
	fmt.Println()

	fmt.Println("After Simulation. Picture of World:")
	fmt.Println(wc)
}
