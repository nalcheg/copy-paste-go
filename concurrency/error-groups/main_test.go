package main

import (
	"context"
	"log"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestPositive(t *testing.T) {
	app := NewApp("first", "second", "third")

	errGr := errgroup.Group{}

	errGr.Go(func() error {
		_, err := app.GetFirst()
		return err
	})
	errGr.Go(func() error {
		_, err := app.GetThird()
		return err
	})

	if err := errGr.Wait(); err != nil {
		log.Fatal(err)
	}
}

func TestError(t *testing.T) {
	app := NewApp("first", "second", "third")

	errGr := errgroup.Group{}

	errGr.Go(func() error {
		_, err := app.GetFirst()
		return err
	})
	errGr.Go(func() error {
		_, err := app.GetSecond()
		return err
	})
	errGr.Go(func() error {
		_, err := app.GetThird()
		return err
	})

	if err := errGr.Wait(); err != errorSecond {
		log.Fatal(err)
	}
}

func TestErrgroupWithContext(t *testing.T) {
	tests := []struct {
		name        string
		timeout     time.Duration
		wantedError error
	}{
		{
			name:        "big timeout",
			timeout:     1000 * time.Millisecond,
			wantedError: context.Canceled,
		}, {
			name:        "too small timeout",
			timeout:     10 * time.Millisecond,
			wantedError: context.DeadlineExceeded,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app := NewApp("first", "second", "third")
			ctx, cancel := context.WithTimeout(context.Background(), test.timeout)
			defer cancel()

			errGr, ctx := errgroup.WithContext(ctx)

			go func() {
				select {
				case <-ctx.Done():
					// run something on context timeout
					log.Print("timeout error ", time.Now())
				}
			}()

			errGr.Go(func() error {
				_, err := app.GetFirst()
				return err
			})
			errGr.Go(func() error {
				_, err := app.GetThird()
				return err
			})

			if err := errGr.Wait(); err != nil {
				t.Error(err)
			}

			if test.wantedError != ctx.Err() {
				t.Error("unexpected context error")
			}
		})
	}
}
