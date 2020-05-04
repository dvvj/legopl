package templates

import (
	"html/template"
	"log"
	"os"
)

type TestCase struct {
	RawInput string
	Titles   []string
	RemInput string
	Suffixes []string
	IsName   bool
}

const templ1 = `new Object[]{
 "{{.RawInput}}",
 {{.Titles}},
 "{{.RemInput}}",
 {{.Suffixes}},
 {{.IsName}},
},`

func Output(testcase *TestCase) {
	tcOut := template.Must(
		template.New("templ1").
			Parse(templ1))

	if err := tcOut.Execute(os.Stdout, testcase); err != nil {
		log.Fatal(err)
	}
}
