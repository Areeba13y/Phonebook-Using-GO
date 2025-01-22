package main

import "fmt"

func main() {
	phonebook := make(map[string]string)

	phonebook["John"] = "123-456-7890"
	phonebook["Alice"] = "987-654-3210"
	phonebook["Bob"] = "456-789-1234"

	fmt.Println("Phonebook:", phonebook)

	name := "Alice"
	if phone, exists := phonebook[name]; exists {
		fmt.Printf("Phone number of %s: %s\n", name, phone)
	} else {
		fmt.Printf("Contact %s not found!\n", name)
	}

	phonebook["Bob"] = "111-222-3333"
	fmt.Println("Updated Phonebook:", phonebook)

	delete(phonebook, "John")
	fmt.Println("Phonebook after deletion:", phonebook)
}
