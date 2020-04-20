package initfunc

import "fmt"

var dict map[string]string

func init() {
	dict = make(map[string]string)

	dict["China"] = "CN"
	dict["PRC"] = "CN"
	dict["The States"] = "US"
	dict["The United States"] = "US"
}

func GetShortName(c string) string {
	r, found := dict[c]

	if !found {
		fmt.Println("not found: " + c)
	}
	return r
}
