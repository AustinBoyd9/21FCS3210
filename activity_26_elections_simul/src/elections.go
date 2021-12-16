/*
 * CS3210 - Principles of Programming Languages - Fall 2021
 * Instructor: Thyago Mota
 * Description: Activity 21 - elections
 */

package main

import (
	"sync"
)

const NUM_PROPOSALS = 3
const POPULATION = 1000000

var wg sync.WaitGroup

/*
TODO: create a ballot (an array of int with size = NUM_PROPOSALS) and assign random values of 0 or 1 to each proposal; sleep for a random amount of time (but for no longer than 10s); finally, write the ballot to the ballotBox channel.
*/
func vote(ballotBox chan<- [NUM_PROPOSALS]int) {

}

/*
TODO: create an array of int with size = NUM_PROPOSALS to tally the ballots from the given ballotBox channel; return the results array.
*/
func getResults(ballotBox <-chan [NUM_PROPOSALS]int) [NUM_PROPOSALS]int {
}

func main() {
	// TODO: create a ballotBox channel of ballots with a capacity to hold POPULATION ballots

	// TODO: start POPULATION goroutines using function "vote"

	// TODO: when ALL goroutines finish, show the results of the election!

}
