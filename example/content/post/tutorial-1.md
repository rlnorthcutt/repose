---
title: "Go Tutorial - Building a Static Site Generator - Part 1"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, ZenForge]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  1
## Planning
Zenforge is focused on elegant simplicity. Day to day usage is easy - just add MD files in the directory structure for the site. When you add `content/post/amazing-tool.md` you will generate
`mysite.com/post/amazing-tool.html`

Templates are managed by Go templates one a component and layout based structure.

### Features
- Create and manage content as MD files
- Content is stored in the directory structure that matches your route
- Content directories without an index.md will auto-generate a listing page of all content
- Content md has header meta data for title, description, tags, image, noindex, author, publish date, template
- Content md meta defines metadata and templates
- Templates are Go templates as .tmpl files
- Template sub directories don’t matter - they are just for organization
- Templates are single file components - they store css, js, and html templates
- Templates ideally resolve to HTML5 with inline css/js
- Templates can resolve to custom web elements
MD files use components to extend default formatting
- Templates can define MD tags for rendering - like video, slider, or button
- Web folder is the web/public root for the website to be served from
- Web folder holds the static files and assets like images, common css, common js, specific html routes
- Generated static files are stored in the web directory
- Config file holds any system options or configuration

### Project folder structure
ZenForge takes the simplest approach possible - only 3 directories will be needed in your project. 

- config.yml - single config file in the root
- content - route and content combined
- template - holds .tmpl files 
- web - static assets

### CLI commands
Zenforge is a command line tool for managing your static site. 

- zenforge new      - Creates a new site structure in the given directory
- zenforge help      - List of commands
- zenforge build     - PArses the content directory and builds the site in the web directory
- zenforge serve    - Builds the site and starts an HTTP server on http://localhost:8080


Build options:
    -r, --root <ROOT> Directory to use as root of project (default: .)
    -c, --config <CONFIG> Path to configuration file (default: config.yml)








Development Plan for ZenForge:
Project Setup:
Initialize a new Golang project.
Set up a version control system (e.g., Git).
Create the basic folder structure (config, content, template, web).
Content Management:
Develop functionality to create and manage content as Markdown (MD) files.
Implement the directory structure for content that mirrors the site's routing.
Content Processing:
Code the logic to auto-generate listing pages for directories without an index.md.
Parse header metadata in MD files for various attributes (title, description, tags, etc.).
Template System:
Implement the handling of Go templates (.tmpl files).
Develop the system to process single file components (combining CSS, JS, and HTML).
Template and MD Extensions:
Enable templates to resolve to HTML5 with inline CSS/JS and to custom web elements.
Extend Markdown formatting with custom components.
Static File Generation:
Code the functionality to generate static files and store them in the web directory.
Configuration Management:
Create a system to read and apply settings from a config.yml file.
CLI Tool Development:
Develop CLI commands (new, help, build, serve) with options for root directory and configuration file.
Testing and Debugging:
Implement testing frameworks and methodologies.
Regularly test each component and the entire system for bugs and issues.
Documentation and User Guides:
Write comprehensive documentation covering setup, usage, and customization.
Create user guides for different levels of users.
Deployment and Distribution:
Plan the distribution method (e.g., repository, packaging).
Set up a system for updates and maintenance.

