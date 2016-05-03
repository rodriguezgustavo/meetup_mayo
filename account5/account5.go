package account5

import(
	"fmt"
//	"sync"
)

var users map[int]string

//var loadOne sync.Once

func loadUsers() {
	users = make(map[int]string)

	for i := 1; i <= 100; i++ {
		users[i] = fmt.Sprintf("name%d", i)
	}
}

func GetUser(id int) string {
	if users == nil {
		loadUsers()
	}
	
	/*
	loadOne.Do(loadUsers)
	*/

	return users[id]
}