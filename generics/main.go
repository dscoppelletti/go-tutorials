package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
//
// A generic function needs a way to declare what types it supports. Calling
// code, on the other hand, needs a way to specify whether it is calling with an
// integer or float map.
//
// To support this, the function that declares type parameters in addition to
// its ordinary function parameters. These type parameters make the function
// generic, enabling it to work with arguments of different types. You’ll call
// the function with type arguments and ordinary function arguments.
//
// Each type parameter has a type constraint that acts as a kind of meta-type
// for the type parameter. Each type constraint specifies the permissible type
// arguments that calling code can use for the respective type parameter.
//
// While a type parameter's constraint typically represents a set of types, at
// compile time the type parameter stands for a single type - the type provided
// as a type argument by the calling code. If the type argument's type isn't
// allowed by the type parameter's constraint, the code won't compile.
//
// A type parameter must support all the operations the generic code is
// performing on it. For example, if your function’s code were to try to perform
// string operations (such as indexing) on a type parameter whose constraint
// included numeric types, the code wouldn't compile.
//
// To run your code, in each call the compiler replaced the type parameters with
// the concrete types specified in that call.
//
// . Declare a SumIntsOrFloats function with two type parameters (inside the
// square brackets), K and V, and one argument that uses the type parameters, m
// of type map[K]V. The function returns a value of type V.
// . Specify for the K type parameter the type constraint comparable. Intended
// specifically for cases like these, the comparable constraint is predeclared
// in Go. It allows any type whose values may be used as an operand of the
// comparison operators == and !=. Go requires that map keys be comparable. So
// declaring K as comparable is necessary so you can use K as the key in the map
// variable. It also ensures that calling code uses an allowable type for map
// keys.
// . Specify for the V type parameter a constraint that is a union of two types:
// int64 and float64. Using | specifies a union of the two types, meaning that
// this constraint allows either type. Either type will be permitted by the
// compiler as an argument in the calling code.
// . Specify that the m argument is of type map[K]V, where K and V are the types
// already specified for the type parameters. Note that we know map[K]V is a
// valid map type because K is a comparable type. If we hadn’t declared K
// comparable, the compiler would reject the reference to map[K]V.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// You can declare a type constraint as an interface. The constraint allows any
// type implementing the interface. For example, if you declare a type
// constraint interface with three methods, then use it with a type parameter in
// a generic function, type arguments used to call the function must have all of
// those methods.
// Declaring constraints in this way helps streamline code, such as when a
// constraint is more complex.
type Number interface {
    int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
// This generic function has the same logic as the generic function
// SumIntsOrFloats, but with the new interface type Number as the type
// constraint.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

	// . Call the generic function you just declared, passing each of the maps
	// you created.
	// . Specify type arguments – the type names in square brackets – to be
	// clear about the types that should replace type parameters in the function
	// you're calling.
	// . Print the sums returned by the function.
	// fmt.Printf("Generic Sums: %v and %v\n",
	// 	SumIntsOrFloats[string, int64](ints),
	// 	SumIntsOrFloats[string, float64](floats))

	// You can omit type arguments in calling code when the Go compiler can
	// infer the types you want to use. The compiler infers type arguments from
	// the types of function arguments.
	// Note that this isn’t always possible. For example, if you needed to call
	// a generic function that had no arguments, you would need to include the
	// type arguments in the function call.
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	// Call SumNumbers with each map, printing the sum from the values of each.
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
