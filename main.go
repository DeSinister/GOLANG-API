package main

import (
    "fmt"
    "log"
    "net/http"
)
func homePage(w http.ResponseWriter, r *http.Request){ // HOMEPAGE
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}



func main() {

  router := http.NewServeMux()
  router.HandleFunc("/", homePage)
  router.HandleFunc("/meetings", func(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/meetings" {
            http.NotFound(w, r)
            return
        }
    if r.Method == "GET" {
           MeetingTime(w, r)
    }else if r.Method == "POST" {
            MeetingCreate(w, r)
        } else {
            http.Error(w, "Invalid request method.", 405)
        }

  })
  router.HandleFunc("/meeting/", func(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/meeting/" {
            http.NotFound(w, r)
            return
        }
    if r.Method == "GET" {
           MeetingShow(w, r)
    } else {
            http.Error(w, "Invalid request method.", 405)
        }
  })

  router.HandleFunc("/meeting", func(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/meeting" {
            http.NotFound(w, r)
            return
        }
    if r.Method == "GET" {
           MeetingShow(w, r)
    } else {
            http.Error(w, "Invalid request method.", 405)
        }
  })



  router.HandleFunc("/all_meetings", func(w http.ResponseWriter, r *http.Request){
    if r.URL.Path != "/all_meetings" {
            http.NotFound(w, r)
            return
        }
    if r.Method == "GET" {
           MeetingIndex(w, r)
    } else {
            http.Error(w, "Invalid request method.", 405)
        }

  })

    log.Fatal(http.ListenAndServe(":10000", router)) //HOSTING AT PORT 10000
}
