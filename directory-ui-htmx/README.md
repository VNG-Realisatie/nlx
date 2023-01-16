Directory UI
---

> Using Server Side Rendering

# Optimizations

1. Consider adding package.json with Sass, Htmx & Source Sans font
2. Move filter logic into the Directory API endpoints
3. Move environment selection / redirection logic into its own handler
4. Consider compiling Scss files to CSS using Go, see https://github.com/bep/godartsass
5. Add headers for caching & CSP to static files handler
6. Introduce SVGO for SVGs

# Reference

1. We use BEM for structuring our Sass.
1. Sass files should be prefixed with underscores, since they are partials (https://sass-lang.com/documentation/at-rules/import#partials).
1. Use `dict` syntax eg. `{{template "link" (dict "url" "https://golang.org" "text" "the Go language")}}`
