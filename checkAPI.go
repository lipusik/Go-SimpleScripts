 package main

 import (
     "os"
     "fmt"
     "net/http"
     "bytes"
     "io/ioutil"
     "regexp"
     "strings"
     "time"
     "github.com/joho/godotenv"
     "log"
 )

// var URL_API = "https://petstore.swagger.io/v2/pet"
func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

 func CheckLen() error {

 // just call panic
   if len(os.Args[1]) < 1 {
         os.Exit(1)
   }
   return nil
 }

 func main() {
url_petstore, exists := os.LookupEnv("URL_API")
    defer func() {
         if rec := recover(); rec != nil {
            fmt.Println(rec,"\nUse external arguments in input")
         }
    }()
    if exists {
    CheckLen()
}
    resp_body := CreatePets(url_petstore)
    time.Sleep(10 * time.Second)

    GetPets(resp_body,url_petstore)
 }

 func CreatePets(url_petstore string) (resp_body string) {

    fmt.Println("URL:>", url_petstore)
    petname := os.Args[1]
    fmt.Println("name:", petname)

    var jsonStr = []byte("{\"name\":\""+petname+"\"}")

    body := Request(url_petstore,string(jsonStr),"POST")
    return string(body)
 }

func GetPets (getpet_body string, url_petstore string) {

    jsonResp := getpet_body
    id_value := GetValueFromJson(jsonResp, "id")
    fmt.Println("The id is:", id_value)
    get_url_petstore := url_petstore + "/" + id_value
    fmt.Println("URL:>", get_url_petstore)
    body := Request(get_url_petstore,"nil","GET")
    state_value := GetValueFromJson(body,"status")
    fmt.Println("Status:", state_value)
}


func GetValueFromJson(body string, value string) string {
    key_str := "\"" + value + "\":[^,;\\]}]*"
    reg, _ := regexp.Compile(key_str)
    match := reg.FindString(body)
    keyValMatch := strings.Split(match, ":")
    return string(keyValMatch[1])
}

func Request(url string, body string, method string) string {

    var jsonStr = []byte(body)
    jsonStr_b := bytes.NewBuffer(jsonStr)

    req, err := http.NewRequest(method, url, jsonStr_b)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout:   1 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    fmt.Println("RespStatus:", resp.Status)

    get_body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("RespBody:", string(get_body))
    resp.Body.Close()
    return string(get_body)
}

