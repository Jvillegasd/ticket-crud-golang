package main

import "os"

func main() {
	app := App{}
	app.InitDatabase(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))

	app.InitRoutes()
	app.Run()
}
