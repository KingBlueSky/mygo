package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "jbwang", Expires: expiration}
	http.SetCookie(w, &cookie)
	cookies, _ := r.Cookie("username")
	fmt.Fprintf(w, cookies.Name)

}

func main() {

}
