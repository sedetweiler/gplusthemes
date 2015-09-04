package main

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

type ThemeList struct {
  XMLName xml.Name `xml:"Themes"`
  Themes []Theme `xml:"Theme"`
}

type Theme struct {
	XMLName xml.Name `xml:"Theme"`
	Name string `xml:"Name,attr"`
	Page string `xml:"Page,attr"`
	Tag string `xml:"Tag,attr"`
	Day string `xml:"Day,attr"`
}

// This function produces each theme as a simple <li /> wrapper
func (t Theme) String() string {

	page := strings.Trim(t.Page," ")
	s := ""

	if len(page) > 0 {
		s = fmt.Sprintf("<li><a href=\"https://plus.google.com/+%s\" target=\"_blank\">%s</a><p>#%s  : + curated by ... </p></li>", page, t.Name, t.Tag)
	} else {
		s = fmt.Sprintf("<li>%s<p>#%s  : + curated by ... </p></li>", t.Name, t.Tag)
	}

	return s
}

func main() {
	xmlFile, err := os.Open("db/themedata.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	XMLdata, _ := ioutil.ReadAll(xmlFile)
	
	var q ThemeList
	xml.Unmarshal(XMLdata, &q)
	
	for _, t := range q.Themes {
		fmt.Printf("%s\n ", t.String())
	}

}