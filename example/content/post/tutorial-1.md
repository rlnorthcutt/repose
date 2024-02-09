---
title: "Go Tutorial - Building a Static Site Generator - Part 1"
description: "Exploring the beauty and simplicity of static websites in the modern web era."
tags: [static sites, web development, Repose]
image: /images/zen-static.jpg
index: true
publish: true
author: "Ron Northcutt"
publish_date: "2024-01-30"
template: "blog-post.tmpl"
---

# Go Tutorial - Building a Static Site Generator - Part  1
## Planning
Repose is focused on elegant simplicity. Day to day usage is easy - just add MD files in the directory structure for the site. When you add `content/post/amazing-tool.md` you will generate
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
Repose takes the simplest approach possible - only 3 directories will be needed in your project. 

- config.yml - single config file in the root
- content - route and content combined
- template - holds .tmpl files 
- web - static assets

### CLI commands
Repose is a command line tool for managing your static site. 

- reposenew      - Creates a new site structure in the given directory
- reposehelp      - List of commands
- reposebuild     - PArses the content directory and builds the site in the web directory
- reposeserve    - Builds the site and starts an HTTP server on http://localhost:8080


Build options:
    -r, --root <ROOT> Directory to use as root of project (default: .)

## Why another static site gnerator?
- simpler format
- smaller - hugo is 80Mb  (vs <5Mb>)
- less opinionated







