package main

import "backend/db"

func main() {
	var sql = new(db.SQL)
	sql.Connect()
	defer sql.Close()
}
