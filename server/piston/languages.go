package piston

type Language struct {
	Version     string `json:"version"`
	PistonAlias string `json:"pistonAlias,omitempty"`
}

var SUPPORTED_LANGUAGES = map[string]Language{
	"javascript": {Version: "20.11.1", PistonAlias: "node"},
	"python":     {Version: "3.11.0", PistonAlias: "python"},
}
