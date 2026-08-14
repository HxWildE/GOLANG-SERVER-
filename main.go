package main
import (
	"fmt"
	)


//format package
func main() {

	fmt.Println("Hello From Go!")
	
	// var a ,b int
	var name string

	 var tasks tasklist
	 tasks = writetask(tasks)
     readtask(tasks)

	fmt.Println("Enter Creator hash : " ,name)
	fmt.Scanln(&name)
	fmt.Println("Enter Creator hash : " ,name)	

	var a int = 0
	var b int = 0

	fmt.Println("Enter first number :")
	fmt.Scanln(&a)
	fmt.Println("Enter Second number : ")
	fmt.Scanln(&b)
	
	fmt.Println(runAdd(a , b))
	fmt.Println(runMod(a ,b))
	fmt.Println(runMul(a , b))
	
	taskcreate()


	
	// fmt.Println(runMod(a,b))

	//runnig maths go lib func
	//bcoz it also belongs ot main packg 
	//so we can call it directly



}


//if you create a var , using it is important
