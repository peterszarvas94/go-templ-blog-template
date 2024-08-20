# dev mode:

templ:
	templ generate --watch --proxy="http://localhost:8080/" --open-browser="false"

dev:
	air -c .air.toml

# ssr mode:

build:
	templ generate
	go build -o ./bin/ssr ./cmd/ssr

ssr:
	./bin/ssr
	
# ssg mode:

gen:
	templ generate
	go run ./cmd/gen

ssg:
	go run ./cmd/ssg

# theme switch:
# e.g. make theme name=default

theme:
	@if [ -d "themes/$(name)/static" ]; then \
		cp -r themes/$(name)/static/* static; \
		cp themes/$(name)/templates/*.templ templates 2>/dev/null || true; \
		echo "Theme '$(name)' applied."; \
	else \
		echo "Theme '$(name)' not found."; \
	fi
