package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct{
	infolog *log.Logger
	errorlog *log.Logger
}


func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

    // Initialize a new instance of application containing the dependencies.
	app := &application{
		infolog: infolog,
		errorlog: errorlog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: mux,
	}

	infolog.Printf("Starting server on : %s", *addr)
//	err := http.ListenAndServe(*addr, mux)
    err := srv.ListenAndServe()

	
    errorlog.Fatal(err)
}
