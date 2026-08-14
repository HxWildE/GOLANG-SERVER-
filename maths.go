package main
import ("fmt")

func runAdd(a ,b int) int {
	
	fmt.Println("Adding ...")
	return a + b

}

func runMul(a , b int)  int{
	
	fmt.Println("Multiplying...")
	return b * a
}

func runMod(a, b int) int {
	
	fmt.Println("Modding up ...")
	return a % b
}




