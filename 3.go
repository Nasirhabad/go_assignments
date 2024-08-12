package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// User represents the structure of the user data from the API.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	} `json:"name"`
}

func fetchUsers(url string, wg *sync.WaitGroup, ch chan<- []User) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	ch <- users
}

func main() {
	var wg sync.WaitGroup
	userChannel := make(chan []User)

	url := "https://fakestoreapi.com/users"

	wg.Add(1)
	go fetchUsers(url, &wg, userChannel)

	go func() {
		wg.Wait()
		close(userChannel)
	}()

	for users := range userChannel {
		for _, user := range users {
			fmt.Println("====")
			fmt.Printf("ID: %d\n", user.ID)
			fmt.Printf("Username: %s\n", user.Username)
			fmt.Printf("Email: %s\n", user.Email)
			fmt.Printf("First name: %s\n", user.Name.FirstName)
			fmt.Printf("Last name: %s\n", user.Name.LastName)
			fmt.Println("====")
		}
	}
}
