package utils

import (
	"log"
	"strings"

	"github.com/adrg/frontmatter"
)

type FrontmatterMetaData struct {
	Title              string   `yaml:"title"`
	Date               string   `yaml:"date"`
	Subject            string   `yaml:"subject"`
	Description        string   `yaml:"description"`
	Tags               []string `yaml:"tags"`
	UniYearAndSemester int      `yaml:"uniYearAndSemester"`
}

func ParseMdString(data string) (FrontmatterMetaData, string) {
	var frontmatterData FrontmatterMetaData
	mdContent, err := frontmatter.Parse(strings.NewReader(data), &frontmatterData)
	if err != nil {
		log.Println(err)
	}
	return frontmatterData, string(mdContent)
}
