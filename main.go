package main

// LN-1 There are three way to import packages in a go file
// LN-1.1
import ("fmt" ;"math")

// LN-1.2
// import ("fmt" 
// 	"math")

// LN-1.3 
// import ("fmt")
// import ("math")




func main() {
	fmt.Printf("What is the square root of 7. \nIt is %g", math.Sqrt(7))

	// LN-2 In Go, if a function, variable, or type name starts with an uppercase letter, it's exported (i.e. public — you can access it from another package). In JS, you have to use export keyword to make something public.

	// LN-2.1 
	// fmt.Println(math.pi)  // ❌ not exported

	// LN-2.2
	fmt.Println("\nValue of pi ", math.Pi)  // ✅ exported constant

	fmt.Println("\n", add(2,5))

	fmt.Println(subtract(2,9))

	a,b,c,d := swap("Hello", "how", "are", "you")
	fmt.Println(a,b,c,d)

	fmt.Println(split(20))

	
	fmt.Println(consoleNumber(23))

	fmt.Println(variablesWithInitializers())

	shortVariableDeclarations()

	basicTypes()

	zeroValues()

	contants()

	numericConstants()
} 

// LN-3 A function can take zero or more arguments. In this example, add takes two parameters of type int.
// LN-4 Notice that the type comes after the variable name.
func add(x int, y int) int{
	return x + y 
}

// LN-5 When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
// Like in add function we decalred each type of varible but subtract only 1.
func subtract(x, y int) int {
	return x - y
}

// LN-6 A function can return any number of results.
// The swap function returns three strings.
func swap(x, y , z, a string) (string, string, string, string) {
	return y,a,z,x
}

// LN-7 Go's return values may be named. If so, they are treated as variables defined at the top of the function.A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


// LN-8 The 'var' statement could declare a list of variables at a time, as in function argument lists, the type is last.
// these are declared at the package level or we could say global level in terms of JS
var i, mython, name string 

// LN-9 A var statement can be at package or function level. We see both in this example. just like local level in JS
func consoleNumber(x int) int {
	var i int = x
	return i
}

func variablesWithInitializers() (int, int, bool, bool, string){
	// LN-10 A 'var' declaration can include initializers, one per variable. Like below
	var q,w int = 13,45
	// LN-11 But if an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var c, python, java = true, false, "ishq Bulava!"

	return q,w,c,python, java
}

func shortVariableDeclarations() {
	// LN-12 Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	var x, y = 23,45
	// Shrt variable declaration
	z:= 343
	fmt.Println(x,y,z)
	//LN-13 Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
}


func basicTypes() {
	/* LN-14 Types in golang
		The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

		The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	*/

	// 1. bool

	// 2. string

	// 3. int  int8  int16  int32  int64
	// 4. uint uint8 uint16 uint32 uint64 uintptr

	// 5. byte // alias for uint8
	var b byte = 'B'
	fmt.Println(b)         // prints 66 (ASCII of 'B')

	// 6. rune // alias for int32 (To store things like emoji's and Unicode characters like ₹)
	var x rune = '✅'
	fmt.Println(x) // prints (ASCII code) of ✅ which is 9989


	// 7. float32 float64
	var y float32 = 3.121312
	fmt.Println(y)

	// 8.complex64 complex128 (To store real + imaginary numbers)
	var z complex128 = 1 + 2i
	fmt.Println(z)       // prints (1+2i)
	
}

func zeroValues() {
	// LN-15 Variables declared without an explicit initial value are given their zero value.
	var a int
	var b float32
	var c complex64
	var d bool
	var e uint 
	fmt.Println(a,b,c,d,e)
}


func contants() {
	/* LN-16
		Constants are declared like variables, but with the const keyword.
		Constants can be character, string, boolean, or numeric values.
		Constants cannot be declared using the := syntax.
	 */

	 const name  = "Gurkirat"
	fmt.Println(name)

}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// a << b = a * 2^b
	Big = 1 << 100

	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	// a >> b = a / (2^b)
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	/* LN-17
		Numeric constants are high-precision values.

		An untyped constant takes the type needed by its context.

		Try printing needInt(Big) too.

		(An int can store at maximum a 64-bit integer, and sometimes less.)
	 */

	fmt.Println(needInt(Small)) //21
	fmt.Println(needFloat(Small)) //0.2
	fmt.Println(needFloat(Big)) //1.2676506002282295e+29
}