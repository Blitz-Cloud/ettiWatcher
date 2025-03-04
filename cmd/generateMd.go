package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Blitz-Cloud/ettiHelper/utils"
	"github.com/blitz-cloud/semHelper/config"
	"github.com/davecgh/go-spew/spew"
)

func BlogPost2Md(blogPost utils.BlogPost) error {
	mdTempalte, err := os.ReadFile("/home/ionut/projects/semHelper/tamplate.md")
	if err != nil {
		log.Fatal(err)
	}
	md := fmt.Sprintf(string(mdTempalte), blogPost.Title, blogPost.Date, blogPost.Description, blogPost.UnivYearAndSemester, blogPost.Content, "")
	fileName := fmt.Sprintf("/home/ionut/ettiHelperData/newWan/%s-%d-%s%s", blogPost.Title, blogPost.UnivYearAndSemester, blogPost.Date, ".md")
	spew.Dump(md)
	err = os.WriteFile(fileName, []byte(md), 0766)
	if err != nil {
		return err
	}
	return nil

}
func NewLabsContentParser(file *utils.File, contentArray *[]utils.BlogPost) {
	if strings.Contains(file.Name, "main") {
		metaData := strings.Split(file.Parent.Name, "-")
		date, err := time.Parse("2_Jan_2006", metaData[2])
		if err != nil {
			log.Fatal(err)
		}
		uniYearAndSemester, _ := strconv.Atoi(metaData[1])
		newLab := utils.BlogPost{
			Title:               metaData[0],
			Date:                fmt.Sprintf("%d-%s-%d", date.Day(), date.Month().String()[:3], date.Year()),
			Content:             string(file.Content),
			UnivYearAndSemester: uniYearAndSemester,
		}
		*contentArray = append(*contentArray, newLab)
	}
}

func generateMd(args []string) error {
	var labsFolder utils.FsNode
	var labs []utils.BlogPost
	conf := config.ConfigFile{}
	conf.ReadConfigFile()
	utils.Explorer(conf.LabsLocation, &labsFolder, ".cpp", &labs, NewLabsContentParser)
	spew.Dump(labs)
	for i := 0; i < len(labs); i++ {
		err := BlogPost2Md(labs[i])
		if err != nil {
			fmt.Printf("Am parcurs pana la pozitia %d\n", i)
			log.Fatal(err)
		}
	}
	return nil
}
