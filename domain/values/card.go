package values

import (
	"encoding/base64"
	"path/filepath"
)

type (
	Card struct {
		Answer    int    `json:"answer"`
		Base64Img string `json:"image"`
		FileName  string `json:"-"`
		FileExt   string `json:"-"`
	}
)

func NewCard(answer int, name string) *Card {
	e := filepath.Ext(name) // ex) sample.jpeg -> .jpeg

	return &Card{
		Answer:   answer,
		FileName: name,
		FileExt:  string([]rune(e)[1:]), // remove [.] ex) .jpeg -> jpeg
	}
}

func (c *Card) SetBase64Img(image []byte) {
	c.Base64Img = "data:image/" + c.FileExt + ";base64," + base64.StdEncoding.EncodeToString(image)
}
