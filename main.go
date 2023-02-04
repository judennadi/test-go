package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

type Zenith struct {
	balance int
}

type Fidelity struct {
	balance int
}

func NewZenith() *Zenith {
	return &Zenith{
		balance: 10,
	}
}

func (z *Zenith) GetBalance() int {
	return z.balance
}

func (z *Zenith) Deposit(amount int) {
	z.balance += amount
}

func (z *Zenith) Withdraw(amount int) error {
	newBalance := z.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	z.balance = newBalance
	return nil
}

func NewFidelity() *Fidelity {
	return &Fidelity{
		balance: 0,
	}
}

func (f *Fidelity) GetBalance() int {
	return f.balance
}

func (f *Fidelity) Deposit(amount int) {
	f.balance += amount
}

func (f *Fidelity) Withdraw(amount int) error {
	newBalance := f.balance - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	f.balance = newBalance
	return nil
}

func main() {
	var waitGroup sync.WaitGroup
	// accounts := []BankAccount{
	// 	NewFidelity(),
	// 	NewZenith(),
	// }

	// for _, acct := range accounts {
	// 	acct.Deposit(500)
	// 	acct.Deposit(200)

	// 	balance := acct.GetBalance()
	// 	fmt.Printf("%T bank balance is: $%v \n", acct, balance)
	// }

	// if err := accounts[1].Withdraw(800); err != nil {
	// 	fmt.Println(err)
	// }

	// balance := accounts[1].GetBalance()
	// fmt.Printf("%T bank balance is now: $%v \n", accounts[1], balance)
	ch1 := make(chan int, 1)

	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()

		for i := 1; i <= 5; i++ {
			if i == 1 {
				fmt.Println("G1 sending 10, 5")
				c <- 10
				c <- 5
				// close(c)
			}
			time.Sleep(time.Millisecond * 500)
			fmt.Println(i, "Goroutine1")
		}
	}(ch1)

	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		// fmt.Println("G2:", <-c)
		for i := 1; i <= 5; i++ {
			// if i == 1 {
			// 	fmt.Println("G3 sending 5")
			// 	c <- 5
			// 	close(c)
			// }
			time.Sleep(time.Millisecond * 500)
			fmt.Println(i, "Goroutine3")
		}
	}(ch1)
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		// fmt.Println("G2:", <-c)
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 500)
			if i == 3 {
				fmt.Println("G4 receiving", <-c)
			}
			fmt.Println(i, "Goroutine4")
		}
	}(ch1)

	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 500)
			if i == 3 {
				fmt.Println("G2 receiving", <-c)
				fmt.Println("G2 sending 20")
				c <- 20
				close(c)
			}
			fmt.Println(i, "Goroutine2")
		}
	}(ch1)

	// fmt.Println("G1:", <-ch1)
	// close(ch1)

	// v, ok := <-ch1
	// if ok {
	// 	fmt.Println("ch1 is Open")
	// 	fmt.Println("val is ", v)
	// } else {
	// 	fmt.Println("ch1 is Closed")
	// 	fmt.Println("val is ", v)
	// }

	time.Sleep(time.Second * 7)
	fmt.Println("G11:", <-ch1)
	waitGroup.Wait()
	fmt.Println("Exiting...")
}
