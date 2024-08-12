package pkg

type AppConfig struct {
	ContentDir string
	PublicDir  string
	StaticDir  string
}

var Config = AppConfig{
	ContentDir: "content",
	PublicDir:  "public",
	StaticDir:  "static",
}
