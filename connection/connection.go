package connection

import (
  "context"
  "fmt"
  "os"

  "github.com/jackc/pgx/v4"
)

// globar var pointing to Conn struct of pgx library to execute sql queries 
var Conn *pgx.Conn

func DatabaseConnect() {
  var err error
  databaseUrl := "postgres://postgres:098765@localhost:5432/personal-web"

  // storing to global var Conn
  // context background allows it to keep running
  Conn, err = pgx.Connect(context.Background(), databaseUrl)
  if err != nil {

  //os.Stderr >> write standar error message
	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
  // exit status 1
	os.Exit(1)
  
  }
  fmt.Println("Success connect to database")

}