package main

import "os"

func main() {
	app := App{}
	app.initDatabase(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")
	)
	app.initRoutes()
	app.run()
}