# Chapter 01

## Introduction
Go is written in text files that are then compiled to machine code and packaged into a single, standalone executable file. It is self-contained, with nothing to be installed first to allow it to run. With Go, you write your code once and run it anywhere.

Go has a statically typed and type-safe memory model with a garbage collector that automates memory management. As a result, Go is very performant and efficient compared to dynamically typed languages.

Go is designed to take advantage of multiple CPU cores to do as much work parallel or concurrently as possible.

Go supports also (multibyte) UTF-8 encoding. 

## Code

### Package
Each Go text file starts with a package declaration. To run it directly using `go run <main.go>` it must declare the `package main`. If not, it can be used as a library and imported into other Go code. The entry point to run a Go application is always the `main()` function in the `main` package.

Go's standard library names are `fmt, log, math/rand` etc. A custom library mostly contains a URL, like `github.com/<name>/<package>`.

Go has a module system that makes using external packages easy. To use a new module, add it to your import path. Go will automatically download it when you build your code.

### Variable
A variable holds data temporarily. When you declare a variable, it needs for things:

1. a declaring statement: e.g. _var_
2. a variable name: starting with a UTF-8 letter or `_` (not limited to Latin character)
3. a data type:
4. an initial value:

An example of declaring multiple variables at once is shown in `Chapter01/Exercise01.03`. In `Chapter01/Exercise01.04`, there is an example of a short variable declaration, where the type or an initial value is omitted. Only in a function, there is also another shorthand variant available by omitting also the `var` keyword and inferring its data type. The `:=` example is shown in `Chapter01/Exercise01.05`. Another shortcut is declaring multiple variables on the same line as shown in  `Chapter01/Exercise01.06`.

#### Zero Values
Go defines the following zero values for all core types:

* bool: `false`
* numbers (integer and float): `0`
* string: `` (empty)
* pointer, function, interface, slice, etc.: `nil`

#### Formatting
Go uses a template language to transform passed values to show a variable's value and/or type by using `fmt.Printf`.

* `%v`: any value
* `%+v`: value with extra information, such as struct field names
* `%#v`: Go syntax
* `%T`: print the variable's type
* `%d`: decimal (base 10)
* `%s`: string

Examples are shown in `Chapter01/Exercise01.12`.

### Operators
There are various groups of operators:

* Arithmetic operators: addition, subtraction, and multiplication
* Comparison operators: equal, not equal, less than, or greater than
* Logical operators: used with boolean values
* Address operators: pointers
* Receive operators: Go channels

Their usage is shown in `Chapter01/Exercise01.09`. For joining multiple `string` you can use the `+` symbol. Also, a number of shorthand operators exists: `--`, `++`, `+=`, `-=`. 


