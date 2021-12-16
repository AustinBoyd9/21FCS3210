/*
 * CS3210 - Principles of Programming Languages - Fall 2021
 * Instructor: Thyago Mota
 * Description: Prg04 - Publish Subscribe Simulation
 * Student(s) Name(s):
 */

package main

import (
	"sync"
)

type PubSub struct {
	mu     sync.Mutex
	topics map[string][]chan string
}

var wg sync.WaitGroup

// TODO: create and return a new channel on a given topic, updating the PubSub struct; this is a helper method used by subscriber goroutines
func (ps PubSub) subscribe(topic string) chan string {

}

// TODO: write the given message on all the channels associated with the given topic; if the topic does not exist, do nothing; this is a helper method used by publisher goroutines
func (ps PubSub) publish(topic string, msg string) {

}

// TODO: send messages taken from a given array of message, one at a time and at random intervals, to all topic subscribers; this is the code template used by publisher goroutines
func publisher(ps *PubSub, topic string, msgs []string) {

}

// TODO: read and display all messages received from a particular topic; this is the code template used by subscriber goroutines
func subscriber(ps *PubSub, name string, topic string) {

}

func main() {
	// TODO: create the ps struct

	// TODO: create the arrays of messages to be sent on each topic

	// TODO: set wait group to 2 (# of publishers)

	// TODO: create the publisher goroutines

	// TODO: create the subscriber goroutines
	// John is only interested in "bees"

	// Mary is interested in "bees" and "colorado"

	// TODO: wait for all publishers to be done
}
