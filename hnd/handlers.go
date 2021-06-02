package hnd

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oknors/comhttpus/utl"
	"net/http"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"truncate": utl.Truncate,
	}
)

func parseFiles(tpl string) (*template.Template, error) {
	return template.ParseFiles("tpl/"+tpl+".gohtml", "tpl/sys/base.gohtml", "tpl/sys/head.gohtml", "tpl/sys/style.gohtml", "tpl/sys/error.gohtml", "tpl/sys/footer.gohtml")
}
func Handlers() http.Handler {
	r := mux.NewRouter()
	tld := r.Host("com-http.{tld}").Subrouter()
	tld.HandleFunc("/", indexHandler())

	tld.HandleFunc("/{section}", sectionHandler())
	tld.HandleFunc("/{section}/{page}", pageHandler())

	sub := r.Host("{sub}.com-http.{tld}").Subrouter()
	sub.HandleFunc("/", subHandler())
	sub.HandleFunc("/{section}", sectionHandler())
	sub.HandleFunc("/{section}/{page}", pageHandler())

	r.Headers("Access-Control-Allow-Origin", "*")
	return handlers.CORS()(handlers.CompressHandler(utl.InterceptHandler(r, utl.DefaultErrorHandler)))

}

func indexHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		data := map[string]interface{}{
			"TLD":   tld,
			"Title": "Beyond blockchain - " + tld,
		}
		template.Must(parseFiles("tld/"+tld)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}

func subHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		sub := mux.Vars(r)["sub"]
		data := map[string]interface{}{
			"TLD":   tld,
			"Title": "Beyond blockchain - " + tld + sub,
		}
		funcMap := template.FuncMap{
			"truncate": utl.Truncate,
		}
		template.Must(parseFiles("tld/"+tld)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}
func sectionHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		section := mux.Vars(r)["section"]
		title := section + " - Beyond blockchain"
		data := map[string]interface{}{
			"Title":   title,
			"Section": section,
		}
		template.Must(parseFiles("tld/"+tld)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}

func pageHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tld := mux.Vars(r)["tld"]
		section := mux.Vars(r)["section"]
		page := mux.Vars(r)["page"]
		title := section + " - Beyond blockchain - " + page
		data := map[string]interface{}{
			"Title":   title,
			"Section": section,
			"Page":    page,
		}
		template.Must(parseFiles("tld/"+tld)).Funcs(funcMap).ExecuteTemplate(w, "base", data)
	}
}
