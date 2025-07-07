package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct{
	infolog *log.Logger
	errorlog *log.Logger
	snippets *mysql.SnippetModel
}


//returns a sql db connection pool
	func openDB(dsn string) (*sql.DB, error) {
       db, err := sql.Open("mysql", dsn)
	   if err != nil{
		return nil, err
	   }
	   if err = db.Ping(); err != nil { 
			return nil, err 
		} 
	
	    return db, nil
	}


func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:hardikagarwal27@/snippetbox?parseTime=true", "MySQL Database")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)


	db, err := openDB(*dsn)
	if err != nil {
		errorlog.Fatal(err)
	} 
	defer db.Close()
    
    // Initialize a new instance of application containing the dependencies.
	app := &application{
		infolog: infolog,
		errorlog: errorlog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorlog,
		Handler: app.routes(),
	}

	infolog.Printf("Starting server on : %s", *addr)
//	err := http.ListenAndServe(*addr, mux)
    err = srv.ListenAndServe()


    errorlog.Fatal(err)
}
