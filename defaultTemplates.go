package main

const DefaultTemplate_none = `<!-- default.tmpl -->
<article>
    <div>{{ .MdContent }}</div>
</article>
`

const ListTemplate_none = `<!-- list.tmpl -->
<article>
    <ul>
    {{ range .Links }}
    <li><a href="{{ . }}">{{ . }}</a></li>
    {{ end }}
    </ul>
</article>
`

const PageTemplate_none = `<!-- fullpage.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="/asset/css/styles.css">
</head>
<body>
    {{ template "header.tmpl" . }}
    {{ template "navigation.tmpl" . }}
    <div class="main container">
        {{ .TemplateContent }}
    </div>
    {{ template "footer.tmpl" . }}
</body>
</html>
`

const HeaderTemplate_none = `<!-- header.tmpl -->
<header>
    <h1>Site Logo Here</h1>
    <h2>{{ .SiteName }}</h2>
</header>
`

const NavigationTemplate_none = `<!-- navigation.tmpl -->
<nav>
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/about">About Us</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
</nav>
`

const FooterTemplate_none = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`

const css_none = `/* styles.css */
    
`
