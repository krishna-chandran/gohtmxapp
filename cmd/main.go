package main

import ( 
	"html/template"  
	"io" 
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Contact struct {
	Name  string
	Email string
}

func newContact(name, email string) *Contact {
	return &Contact{
		Name:  name,
		Email: email,
	}
}

type contacts = []Contact

type Data struct {
	Contacts contacts
}

func newData() *Data {
	return &Data{
		Contacts: []Contact{
			*newContact("Krishnakumar", "krishna@gmail.com"),
			*newContact("Kiran", "kiran@gmail.com"),
		},
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()

	data := newData()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		newContact := *newContact(name, email)
		data.Contacts = append(data.Contacts, newContact)
		return c.Render(200, "newentry", newContact)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
