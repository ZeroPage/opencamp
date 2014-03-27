package viewengine

import (
	"net/http"
	"html/template"
	"fmt"
)

func Render(writer http.ResponseWriter, name string, data interface{}){
	writer.Header().Set("Content-Type", "text/html")
	if err := fillTemplate(writer, name, data); err != nil {
		writer.WriteHeader(500);
		fmt.Fprintf(writer, "Error : %v", err)
	}
}
func fillTemplate(writer http.ResponseWriter, name string, data interface{}) error {
	t, err := template.ParseFiles(
		"views/"+name+".template.html",
		"views/layout.template.html")
	t.ExecuteTemplate(writer ,"layout", data)
	return err
}
