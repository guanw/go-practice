package main

import (
	"flag"
	"go.uber.org/fx"
	"log"
	"os"
)
type logIn struct {
	username string
	password string
}

func newFlags() logIn {
	username := flag.String("username", "judew", "username")
	password := flag.String("password", "w201008474", "password")
	return logIn{
		username: *username,
		password: *password,
	}
}

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags */)
	logger.Print("Executing NewLogger.")

	return logger
}

func printFlags( lI logIn, logg *log.Logger) string {
	logg.Print(lI.username + "   " + lI.password)
	return "empty"
}

func main() {
	app := fx.New(
		fx.Provide(
			newFlags,
			NewLogger,
		),
		fx.Invoke(printFlags),
	)

	app.Run()
}