package main
import "fmt"

//format package
func main() {

	fmt.Println("Hello From Go!")
	
	// var a ,b int
	var name string
	
	fmt.Println("Enter Creator hash : " ,name)
	fmt.Scanln(&name)
	for i := 0; i < 3; i++ {
		taskcreate()
	}
	// fmt.Println(runMod(a,b))

	//runnig maths go lib func
	//bcoz it also belongs ot main packg 
	//so we can call it directly



}


//if you create a var , using it is important
