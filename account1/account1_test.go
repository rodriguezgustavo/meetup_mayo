package meetup

import (
	"testing"
)

func TestAccount(t *testing.T) {
	done := make(chan struct{})

	go func() {
		for i:=0;i<1000;i++ {
			Deposit(200)
		}
		done <- struct{}{}
	}()

	go func() {
		for i:=0;i<1000;i++ {
			Deposit(100)
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 1000*200+1000*100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}


/*
	var x []int

	go func() { x=make([]int,10)}
	go func() { x=make([]int,1000000)}

	x[999999]=1 //possible memory corruption

	three parts of slice -> pointer, length, capacity
 */


