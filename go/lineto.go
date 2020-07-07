package main

import (
    "log"
    "net/http"
    "net/url"
    "strings"
    "flag"
    "fmt"
    "os"
    "io/ioutil"
)

var (
    h bool
    token string
    msg string
    fn string
)

func usage() {
    fmt.Println (`lineto v1.0.0
Usage: lineto [Options]
`)
    flag.PrintDefaults()
}

func init() {
    flag.BoolVar   (&h,     "h", false,   "show help")
    flag.StringVar (&token, "t", "",      "assign `token`")
    flag.StringVar (&msg,   "m", "hello", "`message` to send")
    flag.StringVar (&fn,    "f",  "",     "assign `file`")
    flag.Usage = usage
}

func main() {
    flag.Parse()

    if h {
        flag.Usage()
        os.Exit(0)
    }

    //fmt.Println ("token is", token)
    //fmt.Println ("message is", msg)

    if len(token) == 0 {
        fmt.Println ("token is empty")
        os.Exit(1)
    }

    if len(fn) >0 {
        //fmt.Println ("fn is ", fn)
        file, err := os.Open(fn)
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        defer file.Close()

        content, err := ioutil.ReadAll(file)
        //fmt.Print(string(content))
        if len(string(content)) >0 {
            msg = string(content)
        }
    }

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
    req.Header.Set("Authorization", "Bearer "+token)

    _, err2 := c.Do(req)

    if err2 != nil {
        log.Fatal(err2)
    }
}
