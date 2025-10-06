package types

import "time"

type FrontmatterMetaDataType struct {
	Title              string     `yaml:"title"`
	Description        string     `yaml:"description"`
	Date               *time.Time `yaml:"date"`
	Tags               []string   `yaml:"tags"`
	Subject            string     `yaml:"subject"`
	UniYearAndSemester int        `yaml:"uniYearAndSemester"`
	Remote             string
}
type ProjectMetadataType struct {
	FrontmatterMetaDataType
	GitEnable  bool
	DirOnly    bool
	Lang       string
	OpenEditor bool
}
