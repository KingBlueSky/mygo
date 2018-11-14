package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

func main() {

	t := template.New("fieldName example")
	t, _ = t.Parse("hello {{.}} {{.UserName}}! {{.email}}")
	p := Person{UserName: "jbwang", email: "jbwang0106@163.com"}
	t.Execute(os.Stdout, p)
}
