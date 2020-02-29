package main

var ext = map[string]string{
	"ada":          "ada",
	"bash":         "sh",
	"c":            "c",
	"clojure":      "clj",
	"coffeescript": "coffees",
	"cpp":          "cpp",
	"cpp14":        "cpp",
	"csharp":       "cs",
	"d":            "d",
	"elixir":       "ex",
	"erlang":       "erl",
	"fortran":      "for",
	"fsharp":       "fs",
	"go":           "go",
	"groovy":       "groovy",
	"haskell":      "hs",
	"java":         "java",
	"java8":        "java",
	"javascript":   "js",
	"julia":        "jl",
	"kotlin":       "kt",
	"lolcode":      "lol",
	"lua":          "lua",
	"pascal":       "pas",
	"perl":         "perl",
	"php":          "php",
	"r":            "r",
	"objectivec":   "m",
	"pypy":         "py",
	"pypy3":        "py",
	"sbcl":         "lsp",
	"racket":       "rkt",
	"smalltalk":    "st",
	"tcl":          "tcl",
	"visualbasic":  "vb",
	"ruby":         "ruby",
	"rust":         "rs",
	"scala":        "scala",
	"swift":        "swift",
	"ocaml":        "c",
	"python":       "py",
	"python3":      "py",
}

func GetExtension(lang string) string {
	if v, ok := ext[lang]; ok {
		return v
	}
	return "txt"
}
