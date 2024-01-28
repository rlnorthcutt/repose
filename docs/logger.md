# logger.go
There are 4 different types of logger with different tags & additional actions:
* Info
* Warning
* Error
* Fatal

All logger methods follow the same format:
```go
// Defining a new type 'logger'
type logger int
// Declaring a global variable 'log' of type 'logger'
var Logger logger

logger.Error(message string, value ...any) 
```
### Implementation
- `message` is a formatted string, so you can pass parameters into it
- Inside the method, `value` is treated as a slice of `interface{}`. You can 
  iterate over it or pass it to other functions expecting a variadic input.
- `...any` is a variadic parameter. In Go, `any` is an alias for `interface{}`,
   which means it can accept values of any type.
- The `...` before `any` indicates that you can pass zero, one, or multiple 
  values of any type to the method.

### Using the Logger Methods
Here's how you can use these method in different scenarios. We will see all 4
types used interchangeably - the same implementation will work for each.

#### 1. Passing a Single Argument:
If you want to log a simple message without any additional data:
```go
logger.Fatal("An error occurred")
```

#### 2. Passing Multiple Arguments:
To include additional context or data with your log message:
```go
err := someFunction()
if err != nil {
    logger.Fatal("An error occurred:", err)
}
```
In this case, `err` is passed as an additional argument.

#### 3. Formatting Messages:
You can use `Fatal` like `Printf`, combining a format string with variadic arguments:
```go
name := "ZenForge"
err := someFunction()
if err != nil {
    logger.Fatal("Error in %s: %v", name, err)
}
```
Here, `%s` and `%v` are format specifiers, and `name` and `err` are the corresponding arguments.

#### 4. Passing a Slice of Values:
If you have a slice of values and want to pass it as arguments:

```go
values := []any{"value1", 42, true}
logger.Fatal("Multiple values:", values...)
```
The `...` after `values` unpacks the slice into individual arguments.
