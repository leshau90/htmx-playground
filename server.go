
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

type Contact struct {
  Name string
  Email string
}

func newContact(name string, email string) Contact {
  return Contact{
    Name: name,
    Email: email,
  }
}

type Contacts = []Contact

type Data struct {
  Contacts Contacts
}

type newData() Data {
  return Data{
    Contacts:[]Contact{
      newContact("ilman","ilman@sam.com"),
      newContact("taqi","taqi@taqi.com"),
    },
  }
}


func (t *Templates) Render (w io.Writer, name string, data interface{}, c echo.Context) error{
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates{
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}



func main(){
  e := echo.New()
  e.Use(middleware.Logger())

  data:= newData()

  e.Renderer = newTemplate()
  
  e.GET("/",func (c echo.Context) error {
    return c.Render(200, "index",data)
  })
  
  e.POST("/more-contact",func(c echo.Context) error{
    name := c.FormValue("name")
    contact := c.FormValue("email")
    data.Contacts = append(data.Contacts, newContact(name,email))
    return c.Render(200,"index",data)
  })

  e.Logger.Fatal(e.Start(":8888"))
}
