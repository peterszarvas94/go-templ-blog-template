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

var Title = "lytepage"
var IndexTitle = "Lite ssg with the power of go + templ"

var DateLayout = "2006.01.02"
var TimeLayout = "15:04"
