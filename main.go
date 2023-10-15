package main

import "fmt"


func main() {

// Variables and declaration

var first_variable string // First declaration 
first_variable = "some text" // Second declaration
second_variable := "some text 2" // Third declaration

fmt.Println(first_variable)
fmt.Println(second_variable)




/*

ZERO VALUES

    int = 0
    string = ""
    bool = false
    float = 0.0
    byte = 0
    rune = 0
    map = nil
    interface = nil
    slices = nil
    pointer = nil

*/



/*

OUTPUT

    GENERAL
%v default value, default format
%#v go-syntax format

%T print the type of the value

%% Prints the sign %


    INTEGER

%b base 2
%d base 10
%+d base 10 and always show sign
%o base 8
%O base 8 with leading 0o
%x base 16, lowercase
%X base 16, uppercase
%#x base 16, with leading 0x
%4d pad with spaces (width 4, right justified)
%-4d pad with spaces (width 4, left justified)
%04d pad with zeroes (width 4)


    STRING
%s prints the value as a plain string
%q prints the value as a double-quoted string
%8s Prints the value as plain string (width 8, right justified)
%-8s Prints the value as plain string (width 8, left justified)
%x Prints the value as hex dump of byte values
% x Prints the value as hex dump with spaces


    BOOLEAN
%t value of the boolean operator in true or false format ( same as using %v )



    FLOAT
%e Scientific notation with 'e' as exponent
%f Decimal point, no exponent
%.2f Default width, precision 2
%6.2f Width 6, precision 2
%g Exponent as needed, only necessary digits

*/




// Conditionals

    // IF

    if len(first_variable) > 8 {
        fmt.Println("The lenght of first_variable is greather than 8")
    }

    // ELSE IF - ELSE

    if len(first_variable) > len(second_variable) {
        fmt.Println("The lenght of first_variable is greather than second_variable")
    } else if  len(first_variable) == len(second_variable){
        fmt.Println("The first_variable and the second_variable have the same the lenght")
    } else {
        fmt.Println("The lenght of second_variable is greather than the first_variable")
    }



// LOGICAL OPERATORS

// && AND
// || OR 
// ! NOT 



// FUNCTIONS
firstfunction()
fmt.Println(plus(1,2,3))

// MULTIPLE RETURN FUNCTIONS

fmt.Println(vals())

_, c := vals() // BLANCK IDENTIFIER, to skip values
fmt.Println(c)

fmt.Println(retrn_implicit(1,2))

fmt.Println(other_function(2, "Hello"))

_, b := other_function(1, "Hi")
fmt.Println(b)


// ANONYMOUS FUNCTION 

func() {
    fmt.Println("This is an anonymous function")
}()


// ARRAYS

var arr [5]int
arr[0] = 1

var arr2 = [3]int{1,2,3}
fmt.Println(arr2)

arr3 := [2]int{2,3}
fmt.Println(arr3)

arr4 := [...]int{1,2,3} // inferred lengths
fmt.Println(arr4)


// Change elements of an array

arr[0] = 2
arr[1] = 5

// arr1 := [5]int{} // not initialized
// arr2 := [5]int{1,2} // partially initialized
// arr3 := [5]int{1,2,3,4,5} // fully initialized
// arr1 := [5]int{1:10,2:40} // specific element

// Lenght
fmt.Println(len(arr))


// SLICES

slc := []int{1,2,3,5}
fmt.Println(slc)

// cap
fmt.Println(cap(slc))


// Slice with make()
// slice_name := make([]type, length, capacity)

slc2 := make([]string, 5, 5)
fmt.Println(slc2)
fmt.Println(cap(slc2))
fmt.Println(len(slc2))



// MAPS

var map1 = map[string]int{"Julio": 21, "Horacio": 52, "Valentina": 18}

map2 := map[int]string{1: "Opcion 1", 2: "Opcion 2", 3: "Opcion 3"}


fmt.Println(map1)
fmt.Println(map2)


map_make := make(map[string]int)
map_make["Ferrari"] = 1998
map_make["Mercedez Benz"] = 1974

fmt.Println(map_make)


var empty_map map [string]string 
fmt.Println(empty_map)


delete(map1, "Valentina") // Remove element from map
fmt.Println(map1)



val, ok := map1["Valentina"] // Ckeck for a specific element in map
fmt.Println(val, ok) // 0 False

value, exist := map1["Horacio"]
fmt.Println(value, exist) // 52, true

// RANGE
for key, value := range map1 {
    fmt.Println(key, value)
}



// STRUCS

// type struc_name struct{}

type car struct {
    color string
    year int
    overprice bool
}


var ferrari car

ferrari.color = "Red"
ferrari.year = 2023
ferrari.overprice = false

fmt.Println("year: ", ferrari.year)


fmt.Println(ferrari)


// TYPE CASTING??
// INTERFACE? vs STRUCT?
// RUNE?
// BYTE?
// ERRORS, PANIC, RECOVER?


// POINTERS

    a := 5

    fmt.Println(a)

    var ccc  = *&a 
    fmt.Println(ccc)



}


// FUNCTIONS

func firstfunction(){
    fmt.Println("Hey, this is my first function")
}

func plus(a, b, c int) int {
    return a + b + c
}

func vals() (int, int) {
    return 3, 7
}

func retrn_implicit(a, b int) (result int) {
    result = a + b
    return
}

func other_function(x int, y string) (result int, txt1 string) {
    result = x + x
    txt1 = y + " World!"
    return
}


// git config: 
[user]
	name = Thiago
	email = thiagomowszet@gmail.com
[alias]
	co = checkout
	br = branch
	ci = commit
	st = status
	ps = push
    cl = clone
    pl = pull
    sw = switch
[credential]
	helper = store
[http]
	cookiefile = /home/thiago/.gitcookies


