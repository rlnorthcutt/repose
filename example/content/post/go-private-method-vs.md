is there any value in adding a private method to the type?
func (b *Builder) walkContentDir()   vs   func walkContentDir()

Adding a private method to a type in Go, as opposed to a standalone private function, can provide several benefits depending on the context of your application and how you're structuring your code. Each approach has its use cases, and understanding when to use one over the other can help in designing more maintainable and intuitive software architectures.

### Private Method on a Type

```go
func (b *Builder) walkContentDir() {
    // Implementation details...
}
```

**Advantages:**

1. **Context and State Access**: A method associated with a type (like `Builder`) has access to the instance (`b`) on which it's called. This allows the method to read or modify the state of the instance, making it useful for operations that are inherently tied to the state of the object.

2. **Object-Oriented Design**: Attaching methods to types is a cornerstone of object-oriented design in Go. It allows you to encapsulate behavior with the data it operates on, leading to code that is logically grouped and easier to understand.

3. **Readability and Organization**: By associating functionality closely with the types it operates on, you can make your codebase more navigable and self-documenting. It's clear which type a method belongs to and operates on.

**Use Cases**: When the functionality you're implementing is closely tied to the state or behavior of an instance of your type. For example, if `walkContentDir` needs to access or modify fields in `Builder` or calls other methods on `Builder`, it makes sense as a method.

### Standalone Private Function

```go
func walkContentDir() {
    // Implementation details...
}
```

**Advantages:**

1. **Simplicity**: A standalone function is simple and straightforward. It takes input, produces output, and doesn't depend on or alter the state of an object. This can make it easier to reason about, especially for pure functions.

2. **Reusability**: If a function doesn't need to access the state of a specific type, making it a standalone function can make it more reusable across different parts of your application.

3. **Testing**: Standalone functions, especially those that do not rely on global state or complex inputs, can be easier to test as their behavior is fully determined by their inputs.

**Use Cases**: When the functionality is generic and not inherently tied to a specific type's state. For example, if `walkContentDir` is a utility function that could be used by multiple types or in different contexts, it might be better as a standalone function.

### Conclusion

The choice between adding a private method to a type versus using a standalone private function depends on the specific needs of your application, how you're managing state, and your overall design philosophy. In general, if the function operates on or requires access to the state of an instance of a type, make it a method. If the function is more of a utility that can stand alone, a standalone function might be more appropriate.