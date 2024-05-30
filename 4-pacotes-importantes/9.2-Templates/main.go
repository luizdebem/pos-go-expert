package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"C#", 30},
		{"Python", 20},
		{"Java", 10},
		{"JavaScript", 5},
		{"C", 10},
		{"C++", 10},
	})
	if err != nil {
		panic(err)
	}
}
