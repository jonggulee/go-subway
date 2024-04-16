package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/jonggulee/go-subway/subway"
)

const (
	port        int    = 8080
	templateDir string = "explorer/templates/"
)

var templates *template.Template

func getNowTime() string {
	now := time.Now()
	kst, _ := time.LoadLocation("Asia/Seoul")
	kstTime := now.In(kst)
	return kstTime.Format("2006-01-02 15:04:05 KST")
}

func home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	kstTime := getNowTime()
	data := subway.GetRealtimePosition()
	templates.ExecuteTemplate(rw, "home", data)

	fmt.Println(kstTime)
	fmt.Println(data)

}

func Start() {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	// handler.HandleFunc("/jake", homeJake)
	// handler.HandleFunc("/health", health)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
