### Literal, Variable and Operators

#### Literal

Basically, literals are untyped value. This makes it easier to be assigned to typed variables. If the variable is untyped then the default literal value will be used as type.

```go
var anInteger = 1000
// anInteger type is int (inferred from the literal)
```

There are many literals
```go
var anInt = 1000 // int literal
var aFloat = 20000 // float64 literal
var aBoolean = false // bool literal
var aString = "test" // string literal
var aRune = '\u1AFF' // rune literal (special values)
```

#### Variable

There are many ways to declare and initialize variables in Go

##### 1. Using the `var` keyword

```go
var myInt int = 10
```

##### 2. Type inferred `var`

```go
var myInferredInt = 200
```

##### 3. Multiple type inferred vars

```go
var var1, var2 = 1, "Bustanil"
```

##### 4. Multiple mixed vars

```go
var (
		anotherInt int  = 19
		anotherBool bool = false
		anotherString, andAnotherString = "test", "test2"
		anotherFloat = 99.9
	)
```

##### 5. The `:=` operator

```go
n := 10 // type inferred (without var keyword)
n, m := 15, 20 // can reassign existing variable as long as there's a new var declared (weird!!)
```

    Note: This form is usually used to declare variables inside functions


##### Idioms

- use the complete form `var x int` to init with zero value
- user the complete form `var x boolean = true` when declaring and initializing var (from literal)
- use `:=` only to assign variables to function return values

#### Operators

- The usual C-like operators, nothing different from C or Java.
- Explicit type conversion

    ```go
    var myFloat = float64(9000)
    ```

#### Const

It's not like Javascript's `const`. It's more limited. It's just a way in Go to name literals. That's it. We cannot assign a variable to a `const`. The value should be resolved at compile time.

```go
const distance = 10
```

There's also **typed** constant

```go
const female bool = true
```

#### Naming Convention

- Use camelCase to name variables and constants.