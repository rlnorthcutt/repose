package main

const DefaultTemplate_tailwind = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const PageTemplate_tailwind = `<!-- page.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.tailwindcss.com"></script>
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

const HeaderTemplate_tailwind = `<!-- header.tmpl -->
<header>
    <h1>Site Logo Here</h1>
    <h2>{{ .SiteName }}</h2>
</header>
`

const NavigationTemplate_tailwind = `<!-- navigation.tmpl -->
<nav>
    <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/about">About Us</a></li>
        <li><a href="/contact">Contact</a></li>
    </ul>
</nav>
`

const FooterTemplate_tailwind = `<!-- footer.tmpl -->
<footer>
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
