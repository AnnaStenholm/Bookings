package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AnnaStenholm/bookings/pkg/config"
	"github.com/AnnaStenholm/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates (a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
		return td
}

//renders templates using html templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache{
		//get the template cache from app config
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}




	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render template

	_, err := buf.WriteTo((w))
	if err != nil{
		fmt.Println("Error writing template to browser",err)
	}


}


func CreateTemplateCache()(map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//fet all of the files names *.page.tmpl from /templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts,err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}


