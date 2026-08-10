package main

import (
	"bufio"
	"fmt"
	"os"
)

func taskcreate() {
	task := []string{}
	cnfrm := "y"

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your Name:")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Hello", name)

	for cnfrm == "y" {
		fmt.Println("Want to enter a new task? (y/n)")
		scanner.Scan()
		cnfrm = scanner.Text()

		if cnfrm == "n" {
			break
		}

		fmt.Println("Enter new task:")
		scanner.Scan()
		newtask := scanner.Text()

		task = append(task, newtask)

		fmt.Printf("New task: %s\n", newtask)
	}

	fmt.Println("===================")
	fmt.Println("TASK CREATED")
	fmt.Println("======= LIST =======")

	for query, tasks := range task {
		fmt.Println(query, tasks)
	}
}