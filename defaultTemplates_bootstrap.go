package main

const DefaultTemplate_bootstrap = `<!-- default.tmpl -->
<article>
    <div>{{ .Content }}</div>
</article>
`

const ListTemplate_bootstrap = `<!-- list.tmpl -->
<article>
    <ul>
    {{ range .Files }}
    <li><a href="{{ .OutputPath }}">{{ .MetaData.title }}</a></li>
    {{ end }}
    </ul>
</article>
`

const PageTemplate_bootstrap = `<!-- fullpage.tmpl -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL" crossorigin="anonymous"></script>
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="/asset/css/styles.css">
</head>
<body>
    {{ template "header.tmpl" . }}
    <main class="main container">
        {{ .Content }}
    </main>
    {{ template "footer.tmpl" . }}
</body>
</html>
`

const HeaderTemplate_bootstrap = `<!-- header.tmpl -->
<div class="px-4">
    <header class="d-flex flex-wrap justify-content-center py-3 mb-4 border-bottom">
      <a href="/" class="d-flex align-items-center mb-3 mb-md-0 me-md-auto text-dark text-decoration-none">
        {{ .Logo }}  
        <span class="fs-4 mx-3">{{ .Title }}</span>
      </a>

      {{ template "navigation.tmpl" . }}
    </header>
</div>
`

const NavigationTemplate_bootstrap = `<!-- navigation.tmpl -->
<ul class="nav nav-pills">
    <li class="nav-item"><a href="/" class="nav-link">Home</a></li>
    <li class="nav-item"><a href="/test.html" class="nav-link">Test page</a></li>
    <li class="nav-item"><a href="#" class="nav-link">About</a></li>
</ul>
`

const FooterTemplate_bootstrap = `<!-- footer.tmpl -->
<footer class="pt-5 my-5 text-muted border-top container">
    <p>&copy; 2024 Site Name. All rights reserved.</p>
</footer>
`
const css_bootstrap = `/* styles.css */

`
