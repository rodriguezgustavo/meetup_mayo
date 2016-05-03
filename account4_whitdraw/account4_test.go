package account4

import "testing"

func TestAccount(t *testing.T) {
	done := make(chan struct{})

	deposits:=0
	deposits2:=0
	whithdraws:=0

	go func() {
		for i:=0;i<1000;i++ {
			Deposit(200)
			deposits+=200
		}
		done <- struct{}{}
	}()

	go func() {
		for i:=0;i<1000;i++ {
			Deposit(200)
			deposits2+=200
		}
		done <- struct{}{}
	}()


	go func() {
		for i:=0;i<1000;i++ {
			if(Withdraw(200)){
				whithdraws+=200
			}
		}
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if got, want := Balance(), deposits+deposits2-whithdraws; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

