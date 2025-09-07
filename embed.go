package main

import (
	_ "embed"
)

//go:embed static/docs.html
var DocsHTML []byte
