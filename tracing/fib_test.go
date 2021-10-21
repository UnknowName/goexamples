package main

import (
	"log"
	"os"
	"os/signal"
)

func main2() {
	l := log.New(os.Stdout, "", 0)
	app := NewApp(os.Stdin, l)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)
	go func() {
		errCh <- app.Run()
	}()
	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
