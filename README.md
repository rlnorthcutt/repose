# Repose
An elegant and simple GOlang static site generator.

## Commands:
Usage: repose [OPTIONS] <COMMAND>

Commands:
- init    - Initialize a new Repose project
- new     - Create new content. Usage: repose new [CONTENTTYPE] [FILENAME]
- build   - Build the site.
- help    - Show this help message 

- preview*- Setup a local server to preview the site (*not implemented)
- demo*   - Generate demo content (*not implemented)
- update* - Update the repose binary (*not implemented)
	
Options:
-r, --root <ROOT> Directory to use as root of project (default: ./)
-c, --config <CONFIG> Path to configuration file (default: config.yml)

### To build the command
```
go build
```
However, we can create a smaller binary with this command:
```
go build -ldflags="-s -w"
```
