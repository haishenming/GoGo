package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Name string
		Age int
		Addr string
	}

	user := User{Name:"leo", Age:12, Addr:"China"}

	userJson, _ := json.Marshal(user)

	fmt.Println(string(userJson))

	var newUser User
	json.Unmarshal(userJson, &newUser)

	fmt.Println(newUser)


}
