package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct{
	ID   int
	Title string
	status string
	queued bool
}

type tasklist []Task
//SLICE NOT AN ARRAY (KINDA VECTOR CPP)

func writetask(tsk tasklist) tasklist{
	
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter ID OF task")
	scanner.Scan()
	id,_ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter name OF task")
	scanner.Scan()
	name := scanner.Text()
	
	fmt.Println("Enter status task")
	scanner.Scan()
	status := scanner.Text()
	
	fmt.Println("Enter if its queued or not ? ")
	scanner.Scan()
	queued,_ := strconv.ParseBool(scanner.Text())

	tsk = append(tsk, Task{
		ID : id,
	    Title : name,
	    status : status ,
		queued : queued,
	})

	return tsk

}

func readtask(tsk tasklist){
	for i := 0; i < len(tsk); i++ {
		fmt.Println("ID:", tsk[i].ID)
		fmt.Println("Title:", tsk[i].Title)
		fmt.Println("Status:", tsk[i].status)
		fmt.Println("Queued:", tsk[i].queued)
	}
}


