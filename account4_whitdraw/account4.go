package account4

import "sync"

var (
	mu   sync.RWMutex
	balance int //convention always after mutex
)

func Deposit(amount int) {
	mu.Lock()
	balance=balance+amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool{
	Deposit(-amount)
	if(Balance()<0){
		Deposit(amount)
		return false // not enough funds
	}
	return true
}
