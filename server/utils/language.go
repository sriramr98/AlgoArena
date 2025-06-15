package utils

var extensions = map[string]string{
	"javascript": ".js",
	"python":     ".py",
}

func GetLanguageExtension(language string) string {
	return extensions[language]
}
