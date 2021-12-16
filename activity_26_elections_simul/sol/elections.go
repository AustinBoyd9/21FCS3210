/*
 * CS3210 - Principles of Programming Languages - Fall 2021
 * Instructor: Thyago Mota
 * Description: Activity 21 - elections
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUM_PROPOSALS = 3
const POPULATION = 1000000

var wg sync.WaitGroup

/*
TODO: create a ballot (an array of int with size = NUM_PROPOSALS) and assign random values of 0 or 1 to each proposal; sleep for a random amount of time (but for no longer than 10s); finally, write the ballot to the ballotBox channel.
*/
func vote(ballotBox chan<- [NUM_PROPOSALS]int) {
	defer wg.Done()
	var ballot [NUM_PROPOSALS]int
	for i := 0; i < NUM_PROPOSALS; i++ {
		ballot[i] = rand.Intn(2)
	}
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	ballotBox <- ballot
}

/*
TODO: create an array of int with size = NUM_PROPOSALS to tally the ballots from the given ballotBox channel; return the results array.
*/
func getResults(ballotBox <-chan [NUM_PROPOSALS]int) [NUM_PROPOSALS]int {
	var results [NUM_PROPOSALS]int
	for i := 0; i < POPULATION; i++ {
		ballot := <-ballotBox
		for j := 0; j < NUM_PROPOSALS; j++ {
			results[j] += ballot[j]
		}
	}
	return results
}

func main() {
	// TODO: create a ballotBox channel of ballots with a capacity to hold POPULATION ballots
	ballotBox := make(chan [NUM_PROPOSALS]int, POPULATION)

	// TODO: start POPULATION goroutines using function "vote"
	wg.Add(POPULATION)
	fmt.Println("Voters casting their votes now...")
	for i := 0; i < POPULATION; i++ {
		go vote(ballotBox)
	}
	wg.Wait()

	// TODO: when ALL goroutines finish, show the results of the election!
	fmt.Println("Done!")
	fmt.Println("Election Results:")
	fmt.Println(getResults(ballotBox))
	fmt.Println("Batman approves the election results!")
}
