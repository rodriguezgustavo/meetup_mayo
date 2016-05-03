package account5

import(
	"fmt"
//	"sync"
)

var users map[int]string // users map id/name

//var loadOne sync.Once

// Load users
func loadUsers() {
	users = make(map[int]string)

	for i := 1; i <= 100; i++ {
		users[i] = fmt.Sprintf("name%d", i)
	}
}

// Get and user by id
func GetUser(id int) string {
	if users == nil {
		loadUsers()
	}
	
	/*
	loadOne.Do(loadUsers)
	*/

	return users[id]
}