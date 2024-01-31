// defaultcontent.go
package main

const HelpText = `Repose Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
	init    - Initialize a new Repose project
	new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
	build   - Build the site.
	preview - Setup a local server to preview the site
	demo    - Generate demo content
	update  - Update the repose binary
	help    - Show this help message 
	
Options:
	-r, --root <ROOT> Directory to use as root of project (default: .)
	-c, --config <CONFIG> Path to configuration file (default: config.toml)
`

const DefaultConfig = `sitename: {sitename}
author: {author}
editor: nano
contentDirectory: content
outputDirectory: web
url = "http://localhost:8080"

`

const DefaultHTML = `<!DOCTYPE html>
<head>
    <title>{{ .Title }}</title>
</head>
<body>
{{ .Content }}
</body>
</html>
`

const NewMD = `---
title: "{title}"
description: "{contentType} about {title}"
tags: []
image: 
index: true
author: "{author}"
publish_date: 
template: "{contentType}.tmpl"
---
	
# {title}

`
