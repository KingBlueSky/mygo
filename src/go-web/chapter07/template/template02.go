package main

import (
	"html/template"
	"os"
)

func main() {

	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("kong pipeline if demo: {{if ``}} 不会输出。{{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	template.Must(tWithValue.Parse("bu wei kong pipeline if demo: {{if `anything`}} wo you nei rong, wo hui shu chu. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if bu fen {{else}} else bu fen.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

}
