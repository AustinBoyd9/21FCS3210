/*
 * CS3210 - Principles of Programming Languages - Fall 2021
 * Instructor: Thyago Mota
 * Student's Name:
 * Description: Homework 09 - PIN Cracking
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const PIN_SIZE = 8
const DIGITS = "0123456789"

var wg sync.WaitGroup

func crack(pin string) {
	defer wg.Done()
	for {
		attempt := ""
		for i := 0; i < PIN_SIZE; i++ {
			attempt += string(DIGITS[rand.Intn(len(DIGITS))])
		}
		if attempt == pin {
			break
		}
	}
}

func main() {

	// corrected!
	rand.Seed(int64(time.Now().Nanosecond()))

	// TODO: generate a random PIN with size PIN_SIZE to crack

	// begin timing
	start := time.Now()

	// TODO: crack the PIN by launching go routines (hint: start with 4; it should take ~5m)

	// TODO: figure it out how many go routines main has to wait

	// show how long it took
	fmt.Println("PIN was cracked in", time.Since(start))
}
