package main

import "fmt"

func main() {
	fmt.Println("\n&& (and)")
	fmt.Println(" true && true\t:", true && true)   //true
	fmt.Println(" true && false\t:", true && false) //false

	fmt.Println("\n|| (or)")
	fmt.Println(" true || true\t:", true || true)    //true
	fmt.Println(" true || false\t:", true || false)  //true
	fmt.Println("false || false\t:", false || false) //true

	fmt.Println("\n! (not/negation)")
	fmt.Println("!true\t\t:", !true)   //false
	fmt.Println("!false\t\t:", !false) //true

	fmt.Println("\nMultiple logic operators")
	fmt.Println("(27 <= 32) && (32 > 48)\t:", (27 <= 32) && (48 > 32)) //true
	fmt.Println("(27 <= 32) && (32 > 48)\t:", (27 <= 32) && (48 < 32)) //false
	fmt.Println("(27 <= 32) && (32 > 48)\t:", (27 <= 32) || (48 > 32)) //true
	fmt.Println("(27 <= 32) && (32 < 48)\t:", (27 <= 32) || (48 < 32)) //true
}
