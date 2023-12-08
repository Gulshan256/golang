package main
import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println("type of x is: ", fmt.Sprintf("%T", x));

	//  boolean
	var y bool = true
	fmt.Println(y)

	//  integer
	var z int = 10
	fmt.Println(z)

	//  float
	var a float32 = 10.5
	fmt.Println(a)

	//  default value
	var b string
	fmt.Println(b)

	
	// implicit type
	var c = "Hello World"
	fmt.Println(c)
	fmt.Println("type of c is: ", fmt.Sprintf("%T", c));

	//  implicit type
	var d = true
	fmt.Println(d)
	fmt.Println("type of d is: ", fmt.Sprintf("%T", d));


	// no var declarations
	e := "Hello World"
	fmt.Println(e)
	fmt.Println("type of e is: ", fmt.Sprintf("%T", e));

	// no var declarations
	f := true
	fmt.Println(f)
	fmt.Println("type of f is: ", fmt.Sprintf("%T", f));

	// const declarations
	const g string = "Hello World"
	fmt.Println(g)
	fmt.Println("type of g is: ", fmt.Sprintf("%T", g));
	

}

