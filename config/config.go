package config

type dirConfig struct {
	Content string
	Public  string
	Static  string
}

var Dirs = &dirConfig{
	Content: "content",
	Public:  "public",
	Static:  "static",
}
