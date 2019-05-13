package main

import (
  "fmt"
  "bufio"
  "github.com/stretchr/goweb"
  "github.com/jmhodges/levigo"
)

func main() {

  opts := levigo.NewOptions()
  opts.SetCreateIfMissing(true)
  ro := levigo.NewReadOptions()
  wo := levigo.NewWriteOptions()

  db, err := levigo.Open("slugs.leveldb", opts)
  if err == nil { }
  goweb.ConfigureDefaultFormatters()
  goweb.MapFunc("/{slug}", func(c *goweb.Context) {
    
    key := []byte(c.PathParams["slug"])

    if c.IsGet() {
      fmt.Printf("GET %s\n", c.PathParams["slug"])
      data, err := db.Get(ro, key)
      if err == nil { 
        if data == nil {
          c.WriteResponse("Not found", 404)
        } else {
         c.ResponseWriter.Header().Set("Location", string(data))
         c.WriteResponse(data, 307)
       }
      }
    }

    if c.IsPost() {
      bodyio := bufio.NewReader(c.Request.Body)
      body, err := bodyio.ReadString('\n')
      if err == nil { }
      
      fmt.Printf("POST %s %s\n", key, body)
      err = db.Put(wo, key, []byte(body))
    }
    
  })

  fmt.Println("ready")
  goweb.ListenAndServe(":8080")
  fmt.Println("bye bye")
}
