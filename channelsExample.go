package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Lock struct {
	password int
}

func (l *Lock)Check(passwordToCheck int) bool {
	time.Sleep(time.Second / 1000)
	if l.password == passwordToCheck {
		return true
	} 
	return false
}

func NewLock(min, max int) Lock {
	passwd := rand.Intn(max - min) + min
	return Lock{
		password: passwd,
	}
}

func (l *Lock)Unlock(from int, to int, ch chan int) bool {
	for guess := from; guess <= to; guess++ {
		if l.Check(guess){
			ch <- guess
			close(ch)
			return true
		}
	}
	return false
}

func (lock *Lock) Hacker(min int, max int, amount int) int {
	ch := make(chan int)

	nmin, nmax := min, max / amount
	for i := 1; i <= amount; i++ {
		go lock.Unlock(nmin, nmax, ch)
		nmin = nmax + 1
		nmax = max / amount * (i + 1)
	}  
	go lock.Unlock(nmin, max, ch)
	return <- ch
}

func (l *Lock) BruteForce(min, max int) int {
	for guess := min; guess <= max; guess++ {
		if l.Check(guess){
			return guess
		}
	}
	return 0
}

func main() {
	min, max := 1000, 9999
	lock1 := NewLock(min, max)

	// Hacker Block
	tn := time.Now()
	passwd := lock1.Hacker(min, max, 10)
	tp :=  time.Now()
	fmt.Println("[Hacker] Password:", passwd, "Time:", tp.Sub(tn))
	//--------

	// Brute
	tn = time.Now()
	passwd = lock1.BruteForce(min, max)
	tp =  time.Now()
	fmt.Println("[Brute] Password:", passwd, "Time:", tp.Sub(tn))
	//--------
}
