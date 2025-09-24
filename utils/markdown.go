package utils

import (
	"log"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

type FrontmatterMetaDataType struct {
	Title              string     `yaml:"title"`
	Description        string     `yaml:"description"`
	Date               *time.Time `yaml:"date"`
	Tags               []string   `yaml:"tags"`
	Subject            string     `yaml:"subject"`
	UniYearAndSemester int        `yaml:"uniYearAndSemester"`
}

func ParseMdString(data string) (FrontmatterMetaDataType, string) {
	var frontmatterData FrontmatterMetaDataType
	mdContent, err := frontmatter.Parse(strings.NewReader(data), &frontmatterData)
	if err != nil {
		log.Println("Parser")
		log.Fatal(err)

	}
	return frontmatterData, string(mdContent)
}
