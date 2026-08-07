package main
import "fmt"

//format package
func main() {

	fmt.Println("Hello From Go!")
	
	var a ,b int
	var name string
	fmt.Scanln(&name)
	fmt.Println("Creator : " ,name)
	fmt.Println("Running Maths mod func ! give 2 inputs :")
	
	fmt.Println(runMod(a,b))
	//runnig maths go lib func
	//bcoz it also belongs ot main packg 
	//so we can call it directly

}


//if you create a var , using it is important
