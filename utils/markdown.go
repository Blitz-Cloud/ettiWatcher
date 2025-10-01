package utils

import (
	"log"
	"strings"

	"github.com/adrg/frontmatter"
)

func ParseMdString(data string) (FrontmatterMetaDataType, string) {
	var frontmatterData FrontmatterMetaDataType
	mdContent, err := frontmatter.Parse(strings.NewReader(data), &frontmatterData)
	if err != nil {
		log.Println("Parser")
		log.Println(data)
		log.Fatal(err)

	}
	return frontmatterData, string(mdContent)
}
