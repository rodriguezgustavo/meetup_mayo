package account5

import "testing"

func TestGetUser(t *testing.T) {
	done := make(chan bool, 2)

	go func() {
		for i:=0; i < 100; i++ {
			GetUser(i)
		}

		done <- true
	}()

	go func() {
		for i:=0; i < 100; i++ {
			GetUser(i)
		}

		done <- true
	}()

	<- done
	<- done
	
	if got, expected := len(users), 100; got != expected {
		t.Errorf("Got: %d, expected: %d", got, expected)	
	}
}