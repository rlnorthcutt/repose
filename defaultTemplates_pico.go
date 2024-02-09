package main

const DefaultTemplate_pico = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const ListTemplate_pico = `<!-- list.tmpl -->
<article>
    <ul>
    {{ range .Links }}
    <li><a href="{{ . }}">{{ . }}</a></li>
    {{ end }}
    </ul>
</article>
`

const PageTemplate_pico = `<!-- page.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css">
    <link rel="stylesheet" href="/asset/css/styles.css">
</head>
<body>
    {{ template "header.tmpl" . }}
    <div class="main container">
        {{ template .ContentTemplateName . }}
    </div>
    {{ template "footer.tmpl" . }}
</body>
</html>
`

const HeaderTemplate_pico = `<!-- header.tmpl -->
<nav class="container-fluid">
    <ul>
        <li><a href="/" aria-label="Back home">
            {{ .Logo }}
            </a>
        </li>
        <li style="font-size: 2rem">{{ .SiteName }}</li>
    </ul>
    {{ template "navigation.tmpl" . }}
</nav>
`

const NavigationTemplate_pico = `<!-- navigation.tmpl -->
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/test.html">Test page</a></li>
        <li><a href="#">About Us</a></li>
    </ul>
`

const FooterTemplate_pico = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
const css_pico = `/* styles.css */
    
`
