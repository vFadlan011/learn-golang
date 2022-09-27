package main

import "fmt"

func main() {
	// dalam operasi perbandingan, tipe data yang dibandingkan harus sama
	fmt.Println("\n==")
	fmt.Println(" abc == def\t:", "abc" == "def")   // false
	fmt.Println(" abc == abc\t:", "abc" == "abc")   // true
	fmt.Println("true == true\t:", true == true)    // true
	fmt.Println("false== false\t:", false == false) // true

	fmt.Println("\n!=")
	fmt.Println("   a != A\t:", "a" != "A")   // true
	fmt.Println("   a != a\t:", "a" != "a")   // false
	fmt.Println(" 256 != 256\t:", 256 != 256) // false

	fmt.Println("\n<")
	fmt.Println("   a  < b\t:", "a" < "b") //true
	fmt.Println("   A  < a\t:", "A" < "a") //true
	fmt.Println("   5  < 25\t:", 5 < 25)   //true
	fmt.Println("   x  < X\t:", "x" < "X") //false
	fmt.Println("   5  < 5\t:", 5 < 5)     //false

	fmt.Println("\n<=")
	fmt.Println("   a  <= a\t:", "a" <= "a") //true
	fmt.Println("   5  <= 5\t:", 5 <= 5)     //true
	fmt.Println("   7  <= 5\t:", 7 <= 5)     //false

	fmt.Println("\n>")
	fmt.Println("   7  > 5\t:", 7 > 5)     //true
	fmt.Println("   A  > a\t:", "A" > "a") //false
	fmt.Println("   7  > 7\t:", 7 > 7)     //false

	fmt.Println("\n>=")
	fmt.Println("   7 >= 7\t:", 7 >= 7)           //true
	fmt.Println(" cde >= abc\t:", "cde" >= "abc") //true
	fmt.Println("  AB >= ab\t:", "AB" >= "ab")    //false
}
