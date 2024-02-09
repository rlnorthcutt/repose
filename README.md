# Repose
An elegant and simple GOlang static site generator.

## Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
- init    - Initialize a new Repose project
- new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
- build   - Build the site.
- help    - Show this help message 
- preview - Setup a local server to preview the site
	
Options:
-r, --root <ROOT> Directory to use as root of project (default: ./)

### To build the command
```
go build
```
However, we can create a smaller binary with this command:
```
go build -ldflags="-s -w"
```

### Checklist for beta
- template file pattern for teasers
- update listing page html to use templates
- update code to use optional content type template overrides
- 

### Checklist for RC
- refactor config and commands
- refine templates
- default md metadata overrides per content type (templates/metadata.post.yml)
- generate the md override when creating content type template (with default)
- autowire metadata to metatags in template (name them for the metatags)
- 


### Future commands
- demo*   - Generate demo content (*not implemented)
- update* - Update the repose binary (*not implemented)