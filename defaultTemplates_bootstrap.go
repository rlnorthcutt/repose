package main

const DefaultTemplate_bootstrap = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const PageTemplate_bootstrap = `<!-- page.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
    <title>{{ .Title }}</title>
</head>
<body>
    {{ template "header.tmpl" . }}
    {{ template "navigation.tmpl" . }}
    <div class="main container">
        {{ template "default.tmpl" . }}
    </div>
    {{ template "footer.tmpl" . }}
</body>
</html>
`

const HeaderTemplate_bootstrap = `<!-- header.tmpl -->
<header>
    <h1>Site Logo Here</h1>
    <h2>{{ .SiteName }}</h2>
</header>
`

const NavigationTemplate_bootstrap = `<!-- navigation.tmpl -->
<nav>
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/about">About Us</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
</nav>
`

const FooterTemplate_bootstrap = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
