package account2

import "testing"


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