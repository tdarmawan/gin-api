package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	// "github.com/carlmjohnson/feed2json"
	"html/template"

	"github.com/carlmjohnson/gateway"
)

type peserta struct {
	nama      string
	email			string
	alamat    string
	alasan    string
}


func main() {
	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := ""
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		// http.Handle("/", http.FileServer(http.Dir("./public")))
		http.HandleFunc("/api/", RenderTemplate)
		http.HandleFunc("/api/login", Login)
		http.HandleFunc("/api/logout", Logout)
	}
	log.Fatal(listener(portStr, nil))
}

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
}

http.Error(w, "", http.StatusBadRequest)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		var email = r.FormValue("email")

		emails := []string{"tdarmawan@hacktiv8.com","rendih@gmail.com","tiara@gmail.com","shelly@gmail.com","tika@gmail.com"}
		output := GenerateBiodata(emails)

		var argint int

		for i, x := range emails {
			if email == x {
				argint = i
			}
		}

		for i, x := range output {
			if argint == i {
				var data = map[string]string{"email": email, "message":"Welcome "+ x.nama, "alamat":x.alamat, "alasan":x.alasan}
				if err := tmpl.Execute(w, data); err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		}

		return
}

http.Error(w, "", http.StatusBadRequest)
}

func Logout(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
			var tmpl = template.Must(template.New("form").ParseFiles("login.html"))
			var err = tmpl.Execute(w, nil)

			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
	}

http.Error(w, "", http.StatusBadRequest)
}

func GenerateBiodata(bio []string) []peserta {
	nama := []string{"tdarmawan","rendih","Tiaramut","ShellyS","AtikaR"}
	addr := []string{"Kota A", "Kota B", "Kota C", "Kota D", "Kota E"}
	alasan := []string{"Alasan Thomas", "Alasan Rendi", "Alasan Tiara", "Alasan Shelly", "Alasan Tika"}
	generate := make([]peserta, 0)
	var p peserta

	for key, value := range bio {
		p.nama = nama[key]
		p.email = value
		p.alamat = addr[key]
		p.alasan = alasan[key]
		generate = append(generate, p)
	}

	return generate
}

// func cacheControlMiddleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Cache-Control", "public, max-age=300")
// 		h.ServeHTTP(w, r)
// 	})
// }
