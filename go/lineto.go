package main

import (
    "log"
    "net/http"
    "net/url"
    "strings"
)

func main() {
    accessToken := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    msg := "test"

    URL := "https://notify-api.line.me/api/notify"

    u, err := url.ParseRequestURI(URL)
    if err != nil {
        log.Fatal(err)
    }

    c := &http.Client{}

    form := url.Values{}
    form.Add("message", msg)

    body := strings.NewReader(form.Encode())

    req, err := http.NewRequest("POST", u.String(), body)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", "Bearer "+accessToken)

    _, err2 := c.Do(req)

    if err2 != nil {
        log.Fatal(err2)
    }
}
