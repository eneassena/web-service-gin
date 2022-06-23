package insert

import (
	"database/sql"
	"fmt"
	"log"
  
	_ "github.com/go-sql-driver/mysql"
  )
  
  func conn() (*sql.DB, error) {
	dataSource := "root:root@tcp(localhost:3306)/movies_db"
  
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
	  return nil, err
	}
  
	if err := db.Ping(); err != nil {
	  log.Fatal("could not ping the database: ", err)
	  return nil, err
	}
  
	log.Println("connected")
	return db, nil
  }
  
  type movie struct {
	title       string
	rating      float32
	awards      uint32
	releaseDate string
  }
  
  func main() {
	db, err := conn()
	if err != nil {
	  log.Fatal("could not open the conection: ", err)
	}
  
	defer db.Close()
  
	var myMovies []movie
  
	stmt, err := db.Query(`
	 SELECT title, rating, awards, release_date 
	 FROM movies
	`)
	if err != nil {
	  log.Println("failed to query")
	}
	defer stmt.Close()
  
	for stmt.Next() {
	  var oneMovie movie
  
	  err := stmt.Scan(
		&oneMovie.title,
		&oneMovie.rating,
		&oneMovie.awards,
		&oneMovie.releaseDate,
	  )
	  if err != nil {
		log.Println("failed to scan")
	  }
  
	  myMovies = append(myMovies, oneMovie)
	}
  
	for _, item := range myMovies {
	  fmt.Println("Title: ", item.title)
	}
  }