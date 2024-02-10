

1. Parse templates???
2. Walk the content directory
    - build a map of directories & files (which needs to be created)
    - identify which directories need an index
3. Parse files
    1) parse metadata & add to map
    2) generate html - below header
        - md file, generate HTML
        - html file, clean up?
    3) check the metadata for the template file name to use
    4) generate the full HTML from the templates
    5) write the file
4. Create the index files as needed
    - use the map to find the directories that need an index
    - use the template for that type (top level content dir)
    - parse metadata to build index page
    - build index pages with pagination
    - write the index pages
5. Create the sitemap, rss, and JSON search