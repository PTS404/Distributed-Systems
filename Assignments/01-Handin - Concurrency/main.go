package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id                  int
	timesEaten          int
	leftFork, rightFork *Fork //is the fork on the table or not
}

type Fork struct {
	id        int
	timesUsed int
	onTable   chan bool //is the fork on the table or not
	holder    chan int  //the philosopher who picks up the fork
}

func (p *Philosopher) eat() {
	p.timesEaten++
	fmt.Printf("Philosopher %d is eating (count: %d)\n", p.id, p.timesEaten)

	//Drops both forks
	p.leftFork.onTable <- true
	p.rightFork.onTable <- true
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
			<-p.rightFork.onTable
			<-p.leftFork.onTable
		} else {
			<-p.leftFork.onTable
			<-p.rightFork.onTable
		}
		p.leftFork.holder <- p.id
		p.rightFork.holder <- p.id
		p.eat()
	}
}

func (f *Fork) forksTable(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; f.timesUsed < 6; i++ {
		// philosopher := <-f.holder
		<-f.holder
		f.timesUsed++
		
		//fmt.Printf("Fork %d has been picked up by Philosopher %d \n", f.id, philosopher)
	}
}

func main() {
	forks := make([]*Fork, 5)               //Slice of all five forks
	philosophers := make([]*Philosopher, 5) //Slice of all five philosophers

	//https://gobyexample.com/waitgroups
	var wg sync.WaitGroup

	//Makes the forks
	for i := 0; i < 5; i++ {
		forks[i] = &Fork{
			id:        i + 1,
			timesUsed: 0,
			onTable:   make(chan bool, 1),
			holder:    make(chan int, 1),
		}
		forks[i].onTable <- true //Places forks on the table
	}

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

		wg.Add(1)
		go forks[i].forksTable(&wg)
	}

	wg.Wait()

	//Prints how many times each philosopher has eaten in total
	fmt.Println()
	fmt.Println("Summary:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Philosopher %d has eaten %d times\n", philosophers[i].id, philosophers[i].timesEaten)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("Fork %d has been used %d times\n", forks[i].id, forks[i].timesUsed)
	}
}
