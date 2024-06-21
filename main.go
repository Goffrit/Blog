package main

import (
    "database/sql"
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Open database connection
    db, err := sql.Open("sqlite3", "sqlite.db")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    // Query to select all rows from the demo table
    rows, err := db.Query("SELECT * FROM demo")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var id int
        var name string
        var hint string
        err := rows.Scan(&id, &name, &hint)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("ID: %d, Name: %s, Hint: %s\n", id, name, hint)
    }

    // Check for errors during iteration
    if err = rows.Err(); err != nil {
        fmt.Println(err)
        return
    }
}
