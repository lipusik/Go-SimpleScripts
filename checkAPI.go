 package main

 import (
     "os"
     "fmt"
     "net/http"
     "bytes"
     "io/ioutil"
 )

 func CheckLen() error {

 // just call panic
   if len(os.Args[1]) < 1 {
         os.Exit(1)
   }
   return nil
 }

 func main() {
    defer func() {
         if rec := recover(); rec != nil {
            fmt.Println(rec,"\nUse external arguments in input")
         }
    }()
    CheckLen()
    CreatePets()
 }

 func CreatePets() {

    url_petstore := "https://petstore.swagger.io/v2/pet"
    fmt.Println("URL:>", url_petstore)
    petname := os.Args[1]
    fmt.Println("name:", petname)

    var jsonStr = []byte("{\"name\":\""+petname+"\"}")
    req, err := http.NewRequest("POST", url_petstore, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("RespStatus:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("RespBody:", string(body))
 }
