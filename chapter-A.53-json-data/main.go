package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "John Doe",
		Age:  30,
	}

	// Convert struct to JSON
	var jsonData []byte
	jsonData, _ = json.Marshal(user)
	fmt.Println(string(jsonData))

	// Convert JSON to struct
	var user2 User
	var err = json.Unmarshal(jsonData, &user2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user2)

	// Convert JSON to map
	var userMap map[string]interface{}
	json.Unmarshal(jsonData, &userMap)
	fmt.Println(userMap)

	// Decode Array JSON ke Array obj
	jsonStr := `[
	{"name":"John Doe","age":30},
	{"name":"Jane Doe","age":25}
	]`

	var data []User
	var err2 = json.Unmarshal([]byte(jsonStr), &data)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user 1:", data[0].Name)
	fmt.Println("user 2:", data[1])

	// Encode obj ke JSON string
	var obj = []User{{"Jhon cena", 23}, {"Jhon doe", 21}}
	var jsonData2, err3 = json.Marshal(obj)
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}
	var jsonString = string(jsonData2)
	fmt.Println(jsonString)
}
