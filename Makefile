templ:
	templ generate --watch --proxy="http://localhost:8080/" --open-browser="false"

dev:
	air -c .air.toml

gen:
	go run ./cmd/gen

start:
	go run ./cmd/start
