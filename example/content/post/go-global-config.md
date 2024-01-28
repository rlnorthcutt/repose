---
title: "Go app config options"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, ZenForge]
image: /images/zen-static.jpg
noindex: false
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Managing global config in Go
🧙🏾‍♂️: Using a struct to represent your configuration is a common and effective approach in Go, especially for structured data like YAML or JSON. However, if you're looking for alternatives or simpler ways to make your configuration accessible throughout your application, here are a few options:

### 1. Global Variable

You can load the configuration into a global variable at the start of your program. This makes the configuration accessible from anywhere in your application.

```go
var AppConfig Config

func main() {
    var err error
    AppConfig, err = readConfig("config.yml")
    if err != nil {
        log.Fatalf("Error reading config: %v", err)
    }

    // Now AppConfig is accessible throughout your application
}
```

**Pros:**
- Simple and straightforward.
- Easy to access from anywhere in your application.

**Cons:**
- Global state can lead to harder-to-maintain code and potential issues with concurrent access.
- Goes against the principles of encapsulation and modular design.

### 2. Dependency Injection

Pass the configuration object to functions or methods that need it. This is more in line with the principles of clean architecture.

```go
func someFunction(config Config) {
    // Use config
}

func main() {
    config, err := readConfig("config.yml")
    if err != nil {
        log.Fatalf("Error reading config: %v", err)
    }

    someFunction(config)
}
```

**Pros:**
- More maintainable and testable.
- Avoids global state.

**Cons:**
- Requires more boilerplate as you need to pass the config object around.

### 3. Using a Singleton Pattern

Create a singleton that holds your configuration. This is a more controlled way of having a global state.

```go
var once sync.Once
var instance *Config

func GetConfig() *Config {
    once.Do(func() {
        // Load config
        instance, _ = readConfig("config.yml") // Handle error appropriately
    })
    return instance
}
```

**Pros:**
- Ensures only one instance of the config is ever created.
- Lazily loaded.

**Cons:**
- Still a global state, with similar downsides.

### Conclusion

Each approach has its trade-offs. Using a struct and loading it into a global variable is a common pattern in Go and works well for many applications, especially smaller ones or those with less complex configuration needs. For larger applications or those requiring more maintainability and testability, dependency injection or a singleton pattern might be more appropriate.

Choose the approach that best fits your application's complexity, size, and design philosophy. 🚀📚