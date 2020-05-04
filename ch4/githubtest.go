package main

import (
	"fmt"
	"legopl/ch4/github"
	"legopl/ch4/templates"
	"log"
	"os"
	"text/template"
	"time"
)

func ageDesc(createdAt time.Time) string {
	//	fmt.Println(createdAt)
	t := createdAt.AddDate(0, 1, 0)
	if t.After(time.Now()) {
		return "< 1month "
	} else {
		t1 := createdAt.AddDate(1, 0, 0)
		if t1.After(time.Now()) {
			return "[1month,  1year)"
		} else {
			return ">  1year "
		}
	}
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------
Number: {{.Number}}
  User: {{.User.Login}}
 Title: {{.Title | printf "%.64s"}}
   Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func testTemplate2() {
	tc1 := templates.TestCase{
		RawInput: "raw",
		Titles:   []string{"title1"},
		RemInput: "rem",
		Suffixes: []string{"suf1"},
		IsName:   true}
	templates.Output(&tc1)
}

func testGithub() {

	// report, err := template.New("report").
	// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	// 	Parse(templ)
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	report := template.Must(template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	args := os.Args[1:] //[]string{"golang", "elasticsearch", "deep", "learning"} //os.Args[1:]
	result, err := github.SearchIssues(args)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("%20.20s #%-5d %9.9s %.55s\n",
			ageDesc(item.CreatedAt), item.Number, item.User.Login, item.Title)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	testTemplate2()
	//testGithub()
}
