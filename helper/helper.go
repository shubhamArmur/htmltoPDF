package helper

import (
	"bytes"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func ConvertMarkDownToHTML(markdown string) string {
	// convert string to byte
	markDownbyteContent := []byte(markdown)

	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	if err := md.Convert(markDownbyteContent, &buf); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
