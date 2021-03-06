package main

import (
	"os"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"html/template"
	"fmt"
	"log"
)

type ThemeList struct {
	XMLName xml.Name `xml:"Themes"`
	Themes  []Theme `xml:"Theme"`
}

type Theme struct {
	XMLName xml.Name `xml:"Theme"`
	Name    string `xml:"Name,attr"`
	Page    string `xml:"Page,attr"`
	Tag     string `xml:"Tag,attr"`
	Day     string `xml:"Day,attr"`
}

// Fetch the current XML document and return the Theme[]
func openXML(filename string) ThemeList {

	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()
	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var t ThemeList
	xml.Unmarshal(XMLdata, &t)

	return t
}

func handler(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/listing.html")
	themedata := openXML("db/themedata.xml")
	template.Execute(w, &themedata)
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}


