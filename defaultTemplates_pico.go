package main

const DefaultTemplate_pico = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
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
</head>
<body>
    {{ template "header.tmpl" . }}
    <div class="main container">
        {{ template "default.tmpl" . }}
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
        <li><a href="/about">About Us</a></li>
        <li><a href="/test">Test page</a></li>
    </ul>
`

const FooterTemplate_pico = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
