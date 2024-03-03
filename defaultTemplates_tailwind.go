package main

const DefaultTemplate_tailwind = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const ListTemplate_tailwind = `<!-- list.tmpl -->
<article>
    <ul>
    {{ range .Files }}
    <li><a href="{{ .OutputPath }}">{{ .MetaData.title }}</a></li>
    {{ end }}
    </ul>
</article>
`

const PageTemplate_tailwind = `<!-- fullpage.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.tailwindcss.com"></script>
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="/asset/css/styles.css">
</head>
<body>
    {{ template "header.tmpl" . }}
    {{ template "navigation.tmpl" . }}
    <div class="main container">
        {{ .Content }}
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

const css_tailwind = `/* styles.css */
    
`
