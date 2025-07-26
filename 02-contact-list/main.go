package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	Name  string
	Email string
}

func main() {
	var myContacts []Contact
	running := true
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Contact Lists App!")
	fmt.Println("You can add and view your contact here")

	scanner.Buffer(make([]byte, 1024), 1024)
	for running {
		fmt.Println("Menu list: ")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contact")
		fmt.Println("3. Remove Contact")
		fmt.Println("4. Exit")

		scanner.Scan()

		input := scanner.Text()

		switch input {
		case "1":
			fmt.Println("Input your contact name")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println("Input your contact email")
			scanner.Scan()
			email := scanner.Text()
			myContacts = addContact(name, email, myContacts)
		case "2":
			viewContacts(myContacts)
		case "3":
			fmt.Println("Input your contact name or email that you want to delete")
			scanner.Scan()
			keyword := scanner.Text()
			deleteContact(keyword, &myContacts)
		case "4":
			return
		default:
			fmt.Println("Invalid input")
		}
	}
}

func viewContacts(contacts []Contact) {
	if contacts == nil {
		fmt.Println("No contacts to show")
		return
	}

	fmt.Println("Your contact lists:")
	for key, item := range contacts {
		fmt.Printf("%d. Name: %s Email:%s\n", key+1, item.Name, item.Email)
	}
}

func addContact(name, email string, contacts []Contact) []Contact {
	contacts = append(contacts, Contact{Name: name, Email: email})

	fmt.Printf("Success adding %s to contact\n", name)
	return contacts
}

func deleteContact(keyword string, contacts *[]Contact) {
	foundIndex := -1

	for key, value := range *contacts {
		if value.Email == keyword || value.Name == keyword {
			foundIndex = key
			break
		}
	}

	if foundIndex < 0 {
		fmt.Println("Can't delete, no contact with email/name given found")
		return
	}

	*contacts = append((*contacts)[:foundIndex], (*contacts)[foundIndex+1:]...)
}
