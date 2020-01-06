package main

import (
	"log"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestPositive(t *testing.T) {
	app := NewApp("first", "second", "third")

	gr := errgroup.Group{}

	gr.Go(func() error {
		_, err := app.GetFirst()
		return err
	})
	gr.Go(func() error {
		_, err := app.GetThird()
		return err
	})

	if err := gr.Wait(); err != nil {
		log.Fatal(err)
	}
}

func TestError(t *testing.T) {
	app := NewApp("first", "second", "third")

	gr := errgroup.Group{}

	gr.Go(func() error {
		_, err := app.GetFirst()
		return err
	})
	gr.Go(func() error {
		_, err := app.GetSecond()
		return err
	})
	gr.Go(func() error {
		_, err := app.GetThird()
		return err
	})

	if err := gr.Wait(); err != errorSecond {
		log.Fatal(err)
	}
}
