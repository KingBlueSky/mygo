package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friend   []Friend
}

func main() {
	f1 := Friend{Fname: "jbwang"}
	f2 := Friend{Fname: "fy"}

	t := template.New("fieldName example")
	t, _ = t.Parse(`hello {{.UserName}}!
					{{range .Emails}}
						an email {{.}}
					{{end}}
					{{with .Friends}}
					{{range .}}
						a friend {{.Fname}}
					{{end}}
					{{end}}
				`)

	p := Person{UserName: "HELLO", Emails: []string{"jbwang0106@163.com", "fy@163.com"}, Friend: []Friend{f1, f2}}
	t.Execute(os.Stdout, p)
}
