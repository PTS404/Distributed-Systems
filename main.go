package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id                  int
	timesEaten          int
	leftFork, rightFork chan bool //is the fork on the table or not
}

func (p *Philosopher) eat() {
	p.timesEaten++
	fmt.Printf("Philosopher %d is eating (count: %d)\n", p.id, p.timesEaten)

	//Drops both forks
	p.leftFork <- false
	p.rightFork <- false
	fmt.Printf("Philosopher %d drops forks\n", p.id)

	//Makes sure the same person can't eat twice in a row
	time.Sleep(time.Millisecond * 100)
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking \n", p.id)
}

func (p *Philosopher) table(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; p.timesEaten < 3; i++ {
		p.think()

		/* Avoids deadlock by making sure, that the last philosopher will try to
		pick up the right fork first, making the philosopher to his left, able to
		pick up both */
		if p.id == 5 {
			<-p.rightFork
			<-p.leftFork
		} else {
			<-p.leftFork
			<-p.rightFork
		}
		p.eat()
	}
}

func main() {
	forks := make([]chan bool, 5)           //Channel slice of the five forks
	philosophers := make([]*Philosopher, 5) //Slice of all five philosophers

	//https://gobyexample.com/waitgroups
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			forks[index] = make(chan bool, 1)
			forks[index] <- false //Places forks on the table
			defer wg.Done()
		}(i)
	}
	/* All forks must be initiliazed before making the philosophers.
	This also avoids a deadlock, since otherwise it would be an empty fork channel. */
	wg.Wait()

	//Makes the philosophers
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:         i + 1,
			timesEaten: 0,
			leftFork:   forks[i],
			rightFork:  forks[(i+1)%5], //Will reset i to avoid index out of range exception
		}
	}

	//runs for each philosopher
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].table(&wg)
	}

	wg.Wait()

	//Prints how many times each philosopher has eaten in total
	fmt.Println()
	fmt.Println("Summary")
	for i := 0; i < 5; i++ {
		fmt.Printf("Philosopher %d has eaten %d times\n", philosophers[i].id, philosophers[i].timesEaten)
	}
}
