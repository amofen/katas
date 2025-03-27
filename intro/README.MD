## Starter Commands 
```go
go init example.com/mymodule
go run .
go mod tidy
go mod edit --replace example.com/mymodule=../mymodule
go test 
go test -v
go build
go list -f '{{.target}}'
go env -w GOBIN=/custom/path/bin
go install
```

## Key Notes 
### General stuff
* Names are exported and accessible from the package if they start with Capital letter. Otherwise they are undefined. 
* Naked return allow function to return named values ( names declared in the function). Not recommended for long functions -> Readability 
* We can declare variables using ***var*** statement. Can be in function or package level.
* If a variable is initialized, the type is optional.
* **:=** is a short assignement statement, can be used only inside functions. 
* Outside of functions, all statements begin with a keyword ( ***var***, ***func***, ...) 
* We call multi-line import statement "factored" import. 
* Wen use factored ***var*** declarations . 
* Variables without initial values, take their zero value ( 0, false, "" -> empty string)
* Basic types in go are : bool, string, int, uint,uintptr, uintX, intX, float32,float64,byte ( int8 alias ),rune (int32 alias),complex64, complex128. X can be : 8,16,32,64. int,uint and uintptr can be 32bits wide on 32bits systems and same for 64bits. By default use int,unint,uintptr , unless specific need. 
* T(v) will convert the type of the value : v to T. In Go we do explicit type conversion. Example (var i int = int(X) where X is of type float64, var i int = X will not work)   
* When declaring a variable using ***var*** or := the type is inferred from the value on the right side. If the value is untyped (example : numeric value 46 or 23.233 or 3283 +3i) the type can be either int, float64, complex128 depending on the precision of the value. 
* Constant are declared like variables but with keyword ***const*** . They cannot be declared using := . 

### Flow control
* There is only one loop available in Go : for loop . for [init_statement];[condition];[post_statement]
* No parentheses required, but braces are. 
* Init and post statements are optional. i:=1 for ;i<=34; { i++ } will work. In this case we can drop the semicolons, and we will have the equivalent of while loop. 
* We can ommit the condition aswell, for {} will loop forever 
* Same as for loop, the if can have an init statement. Variables in the init statement are available only inside the ***if*** or the ***else***. 
* Switch statement in Go has implicit ***break***, unlike other languages where all cases are executed, In Go only one condition is executed. ***switch [init_statement];[switch_variable]{...}***
* Switch with no condition is equivalent to ***switch true {}***. Can be useful to replace long if then else chains. 

### Flow control : new things 
* ***defer*** statement defers the execution of a function until the surrounding function returns. The function arguments are evaluated, but the function call  not executed until surrounding function returns. 
* Deferred function calls are pushed to a stack . And called in ***LIFO***

### Other types
* Like C language, Go have pointer type . *T ( Example *int , *string ) . Its zero value is ***nil***. But unlike C we can't do arithmetic with them. 
* we can generate a pointer using the ***&*** operand ( example ***intPointer = &intVar***)
* the operator * denotes the pointer value (example ***fmt.Println(*somePtr)***). This is called *dereferencing* or *indirecting*
* ***struct*** is a collection of fields . ***type Coordinates struct {x int y int z int}*** . We can access field using "." example : c := Coordinates(1,2,3)  c.X = 2 fmt.Println(c.X)
* To access a field in a struct pointer we normaly need to do ***(*structPtr).Field*** but we the language permit doing ***structPtr.Field*** ( for simplicity )  
* We can use ***Field_Name:Value*** syntax to specify a subset of values for a struct ( order is irrelevant ) example c:=Coordinates{x:12} 
* [n]T is an array of n values of type T. Length is a part of the array type -> so they can't be resized. 
* []T is a slice. It is a dynamically sized, flexible view into the elements of an array. It is formed by specifying low and high bound ( forming a half open interval on the right ) separated by a colomn example ***mySlice:=myArray[2:4]*** -> *mySlice* will give access to myArray[2] and myArray[3] values . ( important -> it is a view and note a copy )
* We can declare a slice using a slice literal which creates an array and builds a slice that references it. 
* When slicing an array we can ommit the high or the low bound to use the default value ( or both -> mySlice[:]).
* A slice has a length and a capacity. The length is the number of elements in the slice. Capacity is the number of elements in the underlying array counting from the first element. We can obtain the two values using ***len(mySlice)*** and ****cap(mySlice)***. We can increase the length of a slice by reslicing it, if it has the necessary capacity. Going beyond capacity will give a runtime error. 
* Zero value slice is nil. nil slice has a length and a capacity of 0. It has no underlaying array. 
* slices can contain any values including other slices.
* when appending values to a slice , if the underlying array have space,  the values will be replaced .
* There is a ***range*** form of the for loop :  for index,value := range mySlice {...} . 
* We can skip a function returning value with _ . In previous example we can skip index:  for _,value := range mySlice {}. If we want the first value only we can do :  for index := range mySlice {}


