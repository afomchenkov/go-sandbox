## CLI
go build - compile
go run - compile and run // go run main.go deck.go structs.go
go fmt - format
go install - compiles and installs the package
go get - downloads raw sources of someone's package
go test - runs assosiated tests


## Packages
- executable
- reusable

package == project == workspace

Example:
package main
 - main.go
 - support.go
 - helper.go


if -> Error message "go: go.mod file not found in current directory or any parent directory; see 'go help modules'"
Change this: go env -w GO111MODULE=auto
to this: go env -w GO111MODULE=off



Pointers in Go:
- Go is pass-by-value language

```
                RAM
-------------------------------------
Address | Value
-------------------------------------
0000    |  
0001    | human{firstName: "Jim"...}
0002    | 
0003    | 
0004    |
.......
```

&variable - give me the memory address of the value this variable is pointing at
*pointer  - give me the value this memory address is pointing at

Turn address into value with *address
Turn value into address with &value

```

Arrays         |    Slices
----------------------------------------
- primitive DS | - can grow/shrink
- cannot be    | - used 99% of the time
resized        |
- rarely used  |
directly       |
```

Value types: int, float, string, bool, structs - use pointers to change these things in a function
Reference types: slices, maps, channels, pointers, functions - do not worry about pointers with these


```
               Map             |         Struct
-----------------------------------------------------------
- All keys must be same type   | - Values can be different types
- All values must be same type | - Keys do not support indexing
- Keys are indexed, we can     |
iterate over them              | - You need to know all different
- Use to represent a collection| fields at compile time
of related properties          | - Use to represent a "thing"
- Do not need to know all the  | with a lot of different properties
the keys at compile time       | 
- Reference type!!!            | - Value type!!!
```

## Interfaces
func (d deck) shuffle() - can only shuffle a value of type 'deck'
func (s []float64) shuffle() - can only shuffle a value of type '[]float64'
func (s []string) shuffle() - can only shuffle a value of type '[]string'
func (s []int) shuffle() - can only shuffle a value of type '[]int'
