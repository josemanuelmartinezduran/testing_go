module github.com/gobuffalo/github_flavored_markdown

go 1.17

exclude (
	// https://github.com/advisories/GHSA-wxc4-f4m6-wwqv
	gopkg.in/yaml.v2 v2.2.2
	gopkg.in/yaml.v2 v2.2.4
)

require (
	github.com/microcosm-cc/bluemonday v1.0.22
	github.com/sergi/go-diff v1.3.1
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e
	golang.org/x/net v0.7.0
)

require (
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
)
