# GO blog template

## Goals

1. Write your templates in go templ
2. Write your content in markdown
3. Generate static HTML
4. Serve HTML with go

## Dependencies

- go 1.22.5
- air
- templ
- make

## Commands

### SSR

- `make dev`: start dev server with hot reload [http://localhost:7331](http://localhost:7331)
- `make templ`: start templ file generation in watch mode

### SSG

- `make gen`: generate static HTML files
- `make serve`: start prod server [http://localhost:8080](http://localhost:8080) (SSG mode)
