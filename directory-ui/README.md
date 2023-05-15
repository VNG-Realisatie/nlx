Directory UI
---

> Using Server Side Rendering

```shell
(cd ./ports/ui/assets; npm start)
```

# Optimizations

1. Move filter logic into the Directory API endpoints
2. Move environment selection / redirection logic into its own handler
3. Consider compiling Scss files to CSS using Go, see https://github.com/bep/godartsass
4. Add headers for caching & CSP to static files handler
5. Introduce SVGO for SVGs

# Reference

1. We use BEM for structuring our Sass.
1. Sass files should be prefixed with underscores, since they are partials (https://sass-lang.com/documentation/at-rules/import#partials).
1. Use `dict` syntax eg. `{{template "link" (dict "url" "https://golang.org" "text" "the Go language")}}`
