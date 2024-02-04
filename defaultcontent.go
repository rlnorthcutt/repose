// defaultcontent.go
package main

const HelpText = `Repose Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
	init    - Initialize a new Repose project
	new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
	build   - Build the site.
	preview*- Setup a local server to preview the site (*not implemented)
	demo*   - Generate demo content (*not implemented)
	update* - Update the repose binary (*not implemented)
	help    - Show this help message 
	
Options:
	-r, --root <ROOT> Directory to use as root of project (default: ./)
	-c, --config <CONFIG> Path to configuration file (default: config.yml)
`

const DefaultConfig = `sitename: Repose site
author: Creator
editor: nano
contentDirectory: content
outputDirectory: web
url: mysite.com
previewUrl: http://localhost:8080`

const NewMD = `---
title: {title}
description: {contentType} about {title}
tags: []
image: 
index: true
author: {author}
publish_date: 
template: {contentType}.tmpl
---
	
# {title}

`

const DefaultTemplate = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const PageTemplate = `<!-- page.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Title }}</title>
</head>
<body>
    {{ template "header.tmpl" . }}
    {{ template "navigation.tmpl" . }}
    {{ template "default.tmpl" . }}
    {{ template "footer.tmpl" . }}
</body>
</html>
`

const HeaderTemplate = `<!-- header.tmpl -->
<header>
    <h1>Site Logo Here</h1>
    <h2>{{ .SiteName }}</h2>
</header>
`

const NavigationTemplate = `<!-- navigation.tmpl -->
<nav>
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/about">About Us</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
</nav>
`

const FooterTemplate = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
const MarkdownTest = `
---
title: Markdown Test Page
description: A test page to check markdown rendering
tags: []
image: 
index: true
author: Creator
publish_date: 
template: default.tmpl
---

# h1 Heading
## h2 Heading
### h3 Heading
#### h4 Heading
##### h5 Heading
###### h6 Heading
___

## Typographic replacements
Enable typographer option to see result.
(c) (C) (r) (R) (tm) (TM) (p) (P) +-
test.. test... test..... test?..... test!....
!!!!!! ???? ,,  -- ---
"Smartypants, double quotes" and 'single quotes'

## Emphasis
**This is bold text**
__This is bold text__
*This is italic text*
_This is italic text_
~~Strikethrough~~

## Blockquotes
> Blockquotes can also be nested...
>> ...by using additional greater-than signs right next to each other...
> > > ...or with spaces between arrows.

## Lists
Unordered
+ Create a list by starting a line with '+', '-', or '*'
+ Sub-lists are made by indenting 2 spaces:
  - Marker character change forces new list start:
    * Ac tristique libero volutpat at
    + Facilisis in pretium nisl aliquet
    - Nulla volutpat aliquam velit
+ Very easy!

Ordered
1. Lorem ipsum dolor sit amet
2. Consectetur adipiscing elit
3. Integer molestie lorem at massa


1. You can use sequential numbers...
1. ...or keep all the numbers as '1.'

## Tables

| Option | Description |
| ------ | ----------- |
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |

Right aligned columns

| Option | Description |
| ------:| -----------:|
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default. |
| ext    | extension to be used for dest files. |


## Links

[link text](http://dev.nodeca.com)
[link with title](http://nodeca.github.io/pica/demo/ "title text!")
Autoconverted link https://github.com/nodeca/pica (enable linkify to see)

## Images

![Minion](https://octodex.github.com/images/minion.png)
![Stormtroopocat](https://octodex.github.com/images/stormtroopocat.jpg "The Stormtroopocat")

Like links, Images also have a footnote style syntax

![Alt text][id]

With a reference later in the document defining the URL location:

[id]: https://octodex.github.com/images/dojocat.jpg  "The Dojocat"

### Definition lists

Term 1

:   Definition 1
with lazy continuation.

Term 2 with *inline markup*

:   Definition 2

        { some code, part of Definition 2 }

    Third paragraph of definition 2.

_Compact style:_

Term 1
  ~ Definition 1

Term 2
  ~ Definition 2a
  ~ Definition 2b
`
