package main

 

import (
  "database/sql"
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

  var myMovie movie

  stmt := db.QueryRow(`
         SELECT title, rating, awards, release_date 
         FROM movies 
         WHERE id = ?`,
    2)

  err = stmt.Scan(
    &myMovie.title,
    &myMovie.rating,
    &myMovie.awards,
    &myMovie.releaseDate,
  )
  if err != nil {
    log.Println(err)
  }

  log.Println(myMovie)
}