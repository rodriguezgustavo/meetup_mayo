package account4

import "sync"

var (
	mu   sync.RWMutex
	balance int //convention always after mutex
)

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func Withdraw(amount int) bool{
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if(balance<0){
		deposit(amount)
		return false // not enough funds
	}
	return true
}


func deposit (amount int){
	balance+=amount
}