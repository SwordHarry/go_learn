package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if Balance() < amount {
		return false
	}
	Deposit(-amount)
	return true
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	go teller()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		res := Withdraw(200)
		if !res {
			fmt.Println("取款失败")
		}
	}()
	wg.Wait()
	b := Balance()
	fmt.Println(b)
}
