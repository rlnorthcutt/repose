---
title: "Go Tutorial - Building a Static Site Generator - Part 2"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, ZenForge]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  2
One of the best things about Golang is that here is no right or wrong way to build an application. It may be more or less efficient, but if it compiles then it is going to be pretty stable.

coming from OOP background, I wanted to organize it in a similar way according to my own preferences and experience. At the same, time, I want to avoid overly complex

## Strategy
- berak it into sections so we can get to a stopping point each time
- "wire up" most of the files to "mockup" the functionality
- build out each section as we go along
- we will have a stable (but limited functionality) whenever we step away

## MVC model
- Model - data stored in MD files
- View - Go templates
- Routes - `content` structure
- Controller - go files


## Object structure
Golang has interfaces and a type of inheritance, but it isn't actually OOP - it is procedural. But, you can mimic an OOP-like structure on your project, which we will. So, we will talk about objects and methods, but they are'nt _actually_ objects or methods. The goal is to use those words to make it more clear.
- `main.go` - main controller
- `logger.go` - output logging methods
- `command.go` - CLI commands
- `filesystem.go` - read, write, delete files
- `template.go` - manages template rendering

## Rough build plan
@TODO - fix this
### 1. Project Setup:
- Initialize a new Golang project.
- Set up a version control system (e.g., Git).
- Create the basic folder structure (config, content, template, web).
### 1. Content Management:
- Develop functionality to create and manage content as Markdown (MD) files.
- Implement the directory structure for content that mirrors the site's routing.
### 1. Content Processing:
- Code the logic to auto-generate listing pages for directories without an index.md.
- Parse header metadata in MD files for various attributes (title, description, tags, etc.).
### 1. Template System:
Implement the handling of Go templates (.tmpl files).
Develop the system to process single file components (combining CSS, JS, and HTML).
### 1. Template and MD Extensions:
- Enable templates to resolve to HTML5 with inline CSS/JS and to custom web elements.
- Extend Markdown formatting with custom components.
### 1. Static File Generation:
- Code the functionality to generate static files and store them in the web directory.
### 1. Configuration Management:
- Create a system to read and apply settings from a config.yml file.
### 1. Testing and Debugging:
- Implement testing frameworks and methodologies.
Regularly test each component and the entire system for bugs and issues.
### 1. CLI Tool Development:
- Develop CLI commands (new, help, build, serve) with options for root directory and configuration file.
### 1. Documentation and User Guides:
- Write comprehensive documentation covering setup, usage, and customization.
- Create user guides for different levels of users.
### 1. Deployment and Distribution:
- Plan the distribution method (e.g., repository, packaging).
- Set up a system for updates and maintenance.

