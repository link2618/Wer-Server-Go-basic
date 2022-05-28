package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// if r.Header.Get("Authorization") != "Bearer 12345" {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }
			// f(w, r)
			flag := true
			fmt.Println("check auth")
			if flag {
				f(w, r)
			} else {
				return
				//w.WriteHeader(http.StatusUnauthorized)
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
