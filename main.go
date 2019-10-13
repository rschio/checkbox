package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"
)

type Page struct {
	Title string
	Boxes []string
}

var tmpl = `
<!DOCTYPE html><html><body><h1>{{.Title}}</h1><form>{{range .Boxes}}<p><label>
<input type="checkbox"/><span>{{.}}</span></label></p></form>{{end}}</body></html>`

func main() {
	titleFlag := flag.String("t", "", "the title")
	boxesFlag := flag.String("b", "", "the boxes sep by ','")
	outputFlag := flag.String("o", "checkbox.html", "the output file")
	flag.Parse()

	if *titleFlag == "" || *boxesFlag == "" {
		flag.Usage()
		return
	}

	boxes := strings.Split(*boxesFlag, ",")
	t := template.Must(template.New("checkbox").Parse(tmpl))
	p := Page{Title: *titleFlag, Boxes: boxes}

	f, err := os.Create(*outputFlag)
	if err != nil {
		log.Println(err)
		return
	}

	t.Execute(f, p)
	f.Close()
}
