package main

import "fmt"

func showVar_ErrVersion(vars []string) {
	var funcs []func()
	for _, v := range vars {
		showFunc := func() {
			fmt.Printf("Showing var %s\n", v)
		}
		funcs = append(funcs, showFunc)
	}

	for _, f := range funcs {
		f()
	}
}

func showVar_CorrectVersion(vars []string) {
	var funcs []func()
	for _, v := range vars {
		localVar := v
		showFunc := func() {
			fmt.Printf("Showing var %s\n", localVar)
		}
		funcs = append(funcs, showFunc)
	}

	for _, f := range funcs {
		f()
	}
}

func main() {
	strs := []string{
		"abc",
		"def",
	}
	fmt.Println("------- showVar_ErrVersion")
	showVar_ErrVersion(strs)
	fmt.Println("------- showVar_CorrectVersion")
	showVar_CorrectVersion(strs)
}
