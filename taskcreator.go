package main
import "fmt"

func taskcreate(){

	name := ""
	task := []string {}
	cnfrm := "y"

	fmt.Println("Enter your Name :")
	fmt.Scan(&name)

	for cnfrm == "y" {
		fmt.Println("WAnt to Enter new task (y/n)")
		fmt.Scanln(&cnfrm)
		if cnfrm == "n" {
			break
		}

		fmt.Println("Enter new task ")
		newtask := ""
		fmt.Scanln(&newtask)
		task = append(task, newtask)
		fmt.Printf("New task  : %s\n", newtask)
	}

	fmt.Println("===================")
	fmt.Println("TASK CREATED")
	fmt.Println("======	LIST==============")

	for query,tasks := range(task){
		fmt.Println(query , tasks)
	}
}



