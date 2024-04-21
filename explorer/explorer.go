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

type PageData struct {
	PageTitle   string
	CurrentTime string
	Station     []subway.Subway
}

var templates *template.Template

func getNowTime() string {
	now := time.Now()
	kst, _ := time.LoadLocation("Asia/Seoul")
	kstTime := now.In(kst)
	return kstTime.Format("2006-01-02 15:04:05 KST")
}

func home(rw http.ResponseWriter, r *http.Request) {
	stations := []string{"남위례", "복정"}
	if r.URL.Path != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	kstTime := getNowTime()
	var subwayStationData []subway.Subway
	for _, station := range stations {
		subwayStationArrivals := subway.GetRealtimeStationArrival(station)
		subwayStationData = append(subwayStationData, subwayStationArrivals)
	}

	data := PageData{
		PageTitle:   "Home",
		CurrentTime: kstTime,
		Station:     subwayStationData,
	}

	templates.ExecuteTemplate(rw, "home", data)
}

func Start() {
	handler := http.NewServeMux()
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
