package main

import (
	"os"
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

type ThemeList struct {
  XMLName xml.Name `xml:"Themes"`
  Themes []Theme `xml:"Theme"`
}

type Theme struct {
	XMLName xml.Name `xml:"Theme"`
	Name string `xml:"Name,attr"`
	Page int64 `xml:"Page,attr"`
	Tag string `xml:"Tag,attr"`
	Day string `xml:"Day,attr"`
}

// This function produces each theme as a simple <li /> wrapper
func (t Theme) String() string {
	return fmt.Sprintf("<li>%s, Tag:#%s - %s</li>", t.Name, t.Tag, t.Day)
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