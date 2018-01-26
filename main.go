package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	Quiz()
}

// Quiz function runs a math problem command line game
func Quiz() {
	inputTime := flag.Int("time", 30, "an int")
	flag.Parse()
	problems, _ := os.Open("problems.csv")
	// TODO: bonus- shuffle problems
	reader := csv.NewReader(bufio.NewReader(problems))
	score, answered := 0, 0
	var ready string
	fmt.Printf("Ready? Enter Y to start the %v second timer.", *inputTime)
	fmt.Scanln(&ready)
	if ready == "Y" || ready == "y" {
		start := time.Now()
		for {
			line, error := reader.Read()
			elapsed := time.Since(start)
			if time.Duration.Seconds(elapsed) > float64(*inputTime) {
				break
			}
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			var answer int
			fmt.Println("Calculate: ", line[0])
			fmt.Scanln(&answer) // TODO: End function if timer is up, don't wait for answer
			answered = answered + 1
			if strconv.Itoa(answer) == line[1] {
				score = score + 1
			}
		}
	}
	if answered == 12 {
		fmt.Printf("You made %v calculations and got %v correct in %v seconds. You worked through all 12 calculations", answered, score, *inputTime) //TODO	make total possible calculations dynamic
	} else {
		fmt.Printf("You made %v calculations out of 12, and got %v correct in %v seconds.", answered, score, *inputTime) //TODO	make total possible calculations dynamic
	}

}
