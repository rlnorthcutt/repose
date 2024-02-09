# The Benefit of loose coupling with good coding practices

keeping functions from getting too big naturally creates systems, utilites and sub systems.
That makes it easier to refactor code in pieces - even replacing whole segments
By having smaller functions that take advantage of good standardization, you have
multiple entry points into the branching code flow.

The art is NOT in creating functions, building utilities, or following the "rules" strictly.
The art and science of good coding lies in the interface.

^ ### Begin with the end in mind

Well designed interfaces leave plenty of room for reinterpretation and refactoring beyond what you see today

For example, if you have a `logger` object, with a public method `logger.Warn()`, you can completely alter or replace the method without changing anything
as long as it accepts the same inputs and returns the expected output. With the spread operator, we can even leave the door open wider:

```go
// WARN method for the logger type
// This method formats and prints a warning message with yellow color
func (l *Logger) Warn(message string, value ...any) {
	tag := "\u001B[0;33m[WARNING]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}


// Example usage
Logger.Warn("Danger ahead - proceed with caution")  // 0 replacement values

Logger.Warn("Your %s has %i problems. Check the manual", $car, $number)  // 2 replacement values
```

In this case, we can put as many values for string replacement as we want  (...any) without breaking the flow. So, we can use it 
as it exists a million ways. At the same time, we can add extra logging to a file or remote service without changing anything in the
entire application. 

So what does "good interface" mean? Well, we have the spread operator that allows us to 
provide any number of values for string replacement in the message. But can we abstracti it further? Can we make it __more__ configurable? 
Yes - we can:

```go
func (l *Logger) Output(tag string, message string, value ...any) {
	fmt.Printf(tag+message+"\n", value...)
}

```
Simple, works for ANYTHING, I can literally output anything I want... but at this point, I may as well use `fmt.Printf()`
I've abstracted this into nothing. It is not easier or more convenient to use this method, so it has little value.
This is another way to thin kabout how the interface works - it defines the way the developer uses it (not just the code). Good developer experience i
s critical. Plus, I will end up with inconsistent messaging across the entire application, which makes the user experience poor.

Lets try again.


```go
// Remove all 5 logger methods and replace with a single one
unc (l *Logger) Output(type string, message string, value ...any) {
    switch type {
        case "warn"
	    tag := "\u001B[0;33m[WARNING]\u001B[0;39m "
        ...

    }
	fmt.Printf(tag+message+"\n", value...)
}

```

This is MUCH better. It provides a consistent output experience, it is a simple and convenient 
method to use (only one more parameter), and it reduces the number of lines of code. In fact
I'd say it was a good solution.

But it is not good enough.

The first problem is that if and when I need to add extra functionality, the method will grow and get bigger.
Which meand I may want to break it into smaller pieces... but that will require making lots of
other code changes, so I won't __want__ to do that. It will limit not only my options but also
my ability to make architectural changes.

Also, I already have to put the type into the method, so it is slightly more verbose in usage than the original. Ideally, 
functions should be simply named and easily understood without needing close examination. `Logger.Output` is less clear than `Logger.Warn`.

Also, we are flirting with breaking the single responsibility principle. We can't let dogmatic adherance to standards make our coding choices,
but we've now got a single method that is doing the work of five. Thats a bit of an antipattern. If I 


So, the best option is actually the one we started with:
```go
// WARN method for the logger type
// This method formats and prints a warning message with yellow color
func (l *Logger) Warn(message string, value ...any) {
	tag := "\u001B[0;33m[WARNING]\u001B[0;39m "
	fmt.Printf(tag+message+"\n", value...)
}
```

It is better because the __interface__ is better:
 - it is flexible enough to handle almost antyhing
 - it is easy to extend and replace in the future
 - all 5 methods use the same parameter structure

So, all the developer needs to remember is the specific versions (info, warn, fatal, success) ... 
and, those methods are the same name as what the message is. Plus, that also follows a common pattern for the types
of events and messages that are typically used. So - integration with future logging systems is baked in.

Oh - and it is incredibly easy to extend. I originally had the 5 methods, but I realied later that I wanted a 6th one:

```go
// PLAIN method for the logger type
// This method formats and prints a pain message without color
func (l *Logger) Plain(message string, value ...any) {
	tag := "------- "
	fmt.Printf(tag+message+"\n", value...)
}
```

I was able to add this in seconds, and then start implementing it immediately. I was able to refactor a few
`fmt.Printf` calls quickly, and others when I ran across them. Adding it, using it, and managing it are all simple...
and the interface structure told __me__ when to add it... it was when I needed it.

LEarning to think about interfaces is a subtle thing, and it goes way beyond OOP. It is the ahbit of thinking not jsut about what you 
are building, but also __how it will be used__. Taking the time to focus on how things will be used,
and how to make them better... WHILE you are building it... is what elevates your ability to build robust and useful things