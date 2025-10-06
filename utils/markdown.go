package utils

import (
	"log"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/blitz-cloud/ettiWatcher/types"
)

func ParseMdString(data string) (types.FrontmatterMetaDataType, string) {
	var frontmatterData types.FrontmatterMetaDataType
	mdContent, err := frontmatter.Parse(strings.NewReader(data), &frontmatterData)
	if err != nil {
		log.Println("Parser")
		log.Println(data)
		log.Fatal(err)

	}
	return frontmatterData, string(mdContent)
}
