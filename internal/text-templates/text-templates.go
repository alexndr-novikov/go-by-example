package text_templates

import (
	"fmt"
	"os"
	"text/template"
)

func RunSample() {
	fmt.Println("text-templates package output:")
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	t1.Execute(os.Stdout, true)
	t1.Execute(os.Stdout, []int{1, 2, 3})
	t1.Execute(os.Stdout, "not a value")

	create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	write := func(t *template.Template, data any) {
		t.Execute(os.Stdout, data)
	}

	t2 := create("t2", "Name: {{.Name}}??\n")
	write(t2, struct {
		Name string
	}{"John"})

	write(t2, map[string]int{"Name": 1})

	t3 := create("t2", "Condition: {{if . -}} yes {{else -}} no {{end}}\n")
	write(t3, "")
	write(t3, "not empty")

	t4 := create("t4", "Range: \n {{range .}}{{.}}\n {{end}}")
	write(t4, []string{"Go", "Python", "JavaScript", "PHP"})

	fmt.Println("")
}
