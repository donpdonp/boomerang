package main

import (
  "fmt"
  "bufio"
  "net/http"
  "github.com/stretchr/goweb"
  "github.com/stretchr/goweb/context"
  "github.com/jmhodges/levigo"
)

func main() {

  opts := levigo.NewOptions()
  opts.SetCreateIfMissing(true)
  ro := levigo.NewReadOptions()
  wo := levigo.NewWriteOptions()

  db, err := levigo.Open("slugs.leveldb", opts)
  if err != nil { fmt.Printf("db %v\n", err) }

  goweb.Map("GET", "/{slug}", func(c context.Context) error {
    key := c.PathValue("slug")
    fmt.Printf("GET %s\n", key)
    data, err := db.Get(ro, []byte(key))
    if err == nil {
      responseCode := 0
      responseBytes := []byte("")
      if data == nil {
        responseBytes = []byte("Not found")
        responseCode = 404
      } else {
        c.HttpResponseWriter().Header().Set("Location", string(data))
        responseBytes = data
        responseCode = 307
      }
      return goweb.Respond.With(c, responseCode, responseBytes)
    } else {
      return goweb.Respond.With(c, http.StatusInternalServerError, nil)
    }
  })

  goweb.Map("POST", "/{slug}", func(c context.Context) error {
    key := c.PathValue("slug")

    bodyio := bufio.NewReader(c.HttpRequest().Body)
    body, err := bodyio.ReadString('\n')
    if err == nil {
      fmt.Printf("POST %s %s\n", key, body)
      err = db.Put(wo, []byte(key), []byte(body))
      fmt.Printf("POST2 %s %s\n", key, body)
      if err == nil {
        fmt.Printf("POST db good\n")
        return goweb.Respond.With(c, 200, nil)
      } else {
        fmt.Printf("POST db %s %v\n", key, err)
        return goweb.Respond.With(c, http.StatusInternalServerError, nil)
      }
    } else {
      fmt.Printf("POST body err %v\n", key, err)
      return goweb.Respond.With(c, http.StatusInternalServerError, nil)
    }
  })

  addr := ":8080"
  fmt.Printf("ready %s\n", addr)
  s := &http.Server{Addr: addr, Handler: goweb.DefaultHttpHandler()}
  s.ListenAndServe()
  fmt.Println("bye bye")
}
