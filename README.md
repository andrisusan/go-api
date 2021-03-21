# go-api

### Data type in go
#### Integer
- Integer 
    - int8 : Signed 8-bit integers (-128 to 127)
    - int16 : Signed 16-bit integers (-32768 to 32767)
    - int32 : Signed 32-bit integers (-2147483648 to 2147483647)
    - int64 : Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
    - uint8 : Unsigned 8-bit integers (0 to 255)
    - uint16 : Unsigned 16-bit integers (0 to 65535)
    - uint32 : Unsigned 32-bit integers (0 to 4294967295)
    - uint64 : Unsigned 64-bit integers (0 to 18446744073709551615)
    
- Floating types
    - float32 : IEEE-754 32-bit floating-point numbers (1.18x10^-38 - 3.4x10^38)
    - float64 : IEEE-754 64-bit floating-point numbers (2.23x10^-38 - 1.80x10^308)
    - complex64 : Complex numbers with float32 real and imaginary parts
    - complex128 : Complex numbers with float64 real and imaginary parts
    
- Other numeric type 
    - byte : same as unit8
    - rune : same as int32
    - uint : min uint 32 or uint 64 bits following os that used
    - int : same size as unit (min 32 or 64 bits following os that used)
    - uintptr : an unsigned integer to store the uninterpreted bits of a pointer value

#### Boolean 

- bool : true or false

#### String 

- string : collection of character, string begin and end with quotes 
    - len : count character in string
    - "string[number]" : get character with input index
    
    
### Type
- type : alias for data type, sample 
```
type ini string
var hello ini = "Hello" 
```
    
hello will be string data type because ini is alias from string

### Math Operation

- `+` add
- `-` minus
- `*` multiply
- `/` div
- `%` mod

#### augmented assignment
```$xslt
a +=10 like a = a + 10
a -=10 like a = a - 10
a *=10 like a = a * 10
a /=10 like a = a / 10
a %=10 like a = a % 10
```
#### unary operator
```$xslt
++ like a = a+1
-- like a = a-1
+ positive
- negative
! negation
```
### Array
```$xslt
var fruit[2] string

// initial array value
fruit[0] = "Mango"
fruit[1] = "Banana"

//get array data
a := fruit[0]
```

### Slice
Slice is slicing from the array, slice like array with dynamic size. Slice has 3 data : pointer, length and capacity. 
```$xslt
array
```

    
