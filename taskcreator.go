package main

import ("fmt")


func taskcreate(){
	
    var priority int = 1
	name := ""
	task := ""
	age := 0
	
	fmt.Println("Enter your Name :")
	fmt.Scan(&name)

	fmt.Println("Enter your Task :")
	fmt.Scan(&task)

	fmt.Println("Enter your Age :")
	fmt.Scan(&age)

	fmt.Println("===================")
	fmt.Println("TASK CREATED")
	fmt.Println("====================")
	fmt.Printf("Task : %s  allocated by Author %s of age %d years",task,name,age)
	fmt.Printf("Priority : %d",priority)
	priority = priority + 1 

}


