# go

1. go is fast, statically typed (variable types are checked by compiler), compiled language (compiled to machine code), strongly typed language(can't change var type in future)
2. general purpose programming lang
3. built-in testing support
4. object oriented language

### ok ? 

1. if declared and not use -> error 
2. usually entry point is main.go
3. In Go, a directory is a package, and a package can only have one function with a given name (with the exception of init(), which is a special case)
4. no while loop


### Cheatsheet


Here's a compact Go cheatsheet covering essential syntax, concepts, and features:

---

### Basics
- **Package declaration**: All Go files begin with `package` keyword.
  ```go
  package main
  ```
- **Imports**: Use `import` to bring in packages.
  ```go
  import "fmt"
  ```

- **Entry point**: Main function to run a Go program.
  ```go
  func main() {
      fmt.Println("Hello, Go!")
  }
  ```

### Variables
- **Declaration**: Variables declared with `var` or `:=` syntax.
  ```go
  var x int = 10       // Explicit type
  var y = "Hello"      // Implicit type
  z := true            // Short declaration (inside functions only)
  ```

- **Constants**: Declared with `const`.
  ```go
  const Pi = 3.14
  ```

### Data Types
- **Primitive types**: `int`, `float64`, `string`, `bool`
- **Composite types**: `struct`, `array`, `slice`, `map`

### Control Structures
- **If-Else**
  ```go
  if x > 10 {
      fmt.Println("x is greater than 10")
  } else {
      fmt.Println("x is 10 or less")
  }
  ```

- **For Loop** (the only loop in Go)
  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }

  // Range loop
  for i, v := range arr {
      fmt.Println(i, v)
  }
  ```

- **Switch**
  ```go
  switch x {
  case 1:
      fmt.Println("One")
  case 2:
      fmt.Println("Two")
  default:
      fmt.Println("Other")
  }
  ```

### Functions
- **Basic function syntax**
  ```go
  func add(a int, b int) int {
      return a + b
  }
  ```

- **Multiple return values**
  ```go
  func swap(x, y string) (string, string) {
      return y, x
  }
  ```

- **Named return values**
  ```go
  func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
      return
  }
  ```

### Pointers
- **Declaration and usage**
  ```go
  var p *int
  i := 42
  p = &i          // p now points to i
  fmt.Println(*p) // Dereference p
  ```

### Structs
- **Define and initialize**
  ```go
  type Person struct {
      Name string
      Age  int
  }

  p := Person{Name: "Alice", Age: 30}
  ```

- **Access fields**
  ```go
  fmt.Println(p.Name)
  ```

### Methods
- **Method declaration**
  ```go
  func (p Person) greet() string {
      return "Hello, " + p.Name
  }
  ```

### Interfaces
- **Define interface**
  ```go
  type Speaker interface {
      Speak() string
  }

  func (p Person) Speak() string {
      return "Hello, I'm " + p.Name
  }
  ```

- **Use interface in functions**
  ```go
  func sayHello(s Speaker) {
      fmt.Println(s.Speak())
  }
  ```

### Error Handling
- **Error type**
  ```go
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, fmt.Errorf("cannot divide by zero")
      }
      return a / b, nil
  }
  ```

### Goroutines and Channels
- **Goroutine** (lightweight thread)
  ```go
  go func() {
      fmt.Println("Running in a goroutine")
  }()
  ```

- **Channels**: For safe communication between goroutines
  ```go
  ch := make(chan int)
  ch <- 5       // Send to channel
  x := <-ch     // Receive from channel
  ```

### Concurrency Patterns
- **Buffered channels**
  ```go
  ch := make(chan int, 2)
  ```

- **Select statement** (for multiple channel operations)
  ```go
  select {
  case msg := <-ch1:
      fmt.Println(msg)
  case msg := <-ch2:
      fmt.Println(msg)
  default:
      fmt.Println("No messages")
  }
  ```

### Packages and Modules
- **Initialize new module**
  ```bash
  go mod init moduleName
  ```

- **Get packages**
  ```bash
  go get packageName
  ```

### Useful Built-in Functions
- **String length**: `len(str)`
- **Append to slice**: `append(slice, elem)`
- **Delete from map**: `delete(map, key)`

---

This cheatsheet covers the essentials of Go to help you quickly recall key syntax and concepts. Let me know if you need any deeper explanations on any specific topic!
