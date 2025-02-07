package main

import (
	"net/http"
	"context"
	"fmt"
)


type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		} else {
			fmt.Fprint(w, data)
		}
	}
}