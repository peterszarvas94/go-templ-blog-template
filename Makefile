# dev:

dev:
	air -c .air.toml

templ:
	templ generate --watch --proxy="http://localhost:8080/" --open-browser="false"

# ssr:

build:
	go build -o ./bin/ssr ./cmd/ssr

ssr:
	./bin/ssr
	
# ssg:

gen:
	go run ./cmd/gen

ssg:
	go run ./cmd/ssg
