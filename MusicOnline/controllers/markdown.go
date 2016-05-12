package controllers


import (
	"github.com/russross/blackfriday"
)

func ParseDoc(input []byte) []byte {
	htmlBytes := blackfriday.MarkdownCommon(input)

	return htmlBytes
}
