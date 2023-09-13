package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

// Philosopher represents a philosopher with an ID, left fork, right fork, and eat count.
type Philosopher struct {
	id                  int
	leftFork, rightFork chan bool
	eatCount            int
}

// eat simulates a philosopher eating, increments the eat count, and releases forks.
func (p *Philosopher) eat() {
	p.eatCount++
	fmt.Printf("Philosopher %d is eating (count: %d)\n", p.id, p.eatCount)
	time.Sleep(time.Millisecond * 100)
	p.leftFork <- true
	p.rightFork <- true
}

// think simulates a philosopher thinking.
func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Millisecond * 100)
}

// dine simulates a philosopher's dining behavior.
func (p *Philosopher) dine(eatCount int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < eatCount; i++ {
		p.think()
		<-p.leftFork
		<-p.rightFork

		p.eat()
	}
}

func main() {
	// Create forks as channels with initial availability.
	forks := make([]chan bool, numPhilosophers)
	philosophers := make([]*Philosopher, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		forks[i] = make(chan bool, 1)
		forks[i] <- true // Initialize forks as available.
	}

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
			eatCount:  0,
		}
	}

	var wg sync.WaitGroup

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].dine(3, &wg) // Each philosopher dines 3 times.
	}

	wg.Wait()
}
