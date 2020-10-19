package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"
    "io"
    "strconv"
)

func Index(w http.ResponseWriter, r *http.Request) { // BASICS
    fmt.Fprintln(w, "Welcome!")
}

func MeetingCreate(w http.ResponseWriter, r *http.Request) { // CREATE A MEETING
    var meeting Meeting
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &meeting); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    t, flag := RepoCreateMeeting(meeting)
    if flag == "false"{
      fmt.Println(w, "Not Enough Parameters")
    }else if flag == "true"{
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusCreated)
      if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
      }
    }else{
      fmt.Println(w, "There is a Conflict for Particpant: ", flag) // CONFLICT FOUND
      fmt.Println(w, "Meeting ID: ", t.Id)
      fmt.Println(w, "Starting from: ", t.Start)
      fmt.Println(w, "Ending at: ", t.End)
    }
}

func MeetingIndex(w http.ResponseWriter, r *http.Request) { // SHOWS ALL MEETINGS
    fmt.Fprintln(w, "all THe Meetings")
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(meetings); err != nil {
        panic(err)
    }
}
func MeetingTime(w http.ResponseWriter, r *http.Request) { .. CUSTOM MEETING SEARCH
    fmt.Println("GET params were:", r.URL.Query())
    query := r.URL.Query()
    a := query.Get("start")
    b := query.Get("end")
    c := query.Get("participant")
    if a!="" && b!="" {
      fmt.Fprint(w, "Meetings from ", a)
      fmt.Fprintln(w, " To ", b)
      if c == ""{
          fmt.Fprintln(w, RepoMeetingTime(a, b))
      } else {
          fmt.Fprintln(w,"Meetings for Participant", c)
          fmt.Fprintln(w, RepoMeetingTime2(a, b, c))
      }
    } else if c != "" {
      fmt.Fprintln(w,"Meetings for Participant", c)
      fmt.Fprintln(w, RepoMeetingTime3(c))
    } else {
      fmt.Fprintln(w, "Unknown Parameters Found")
    }

}

func MeetingShow(w http.ResponseWriter, r *http.Request) { // FINDING A MEETING USING ID
    query := r.URL.Query()
    v := query.Get("Id")
    MeetingId, err := strconv.Atoi(v)
    _ = err
    fmt.Fprintln(w, "MEETING ID: ", MeetingId)
    fmt.Fprintln(w, "Meeting show:", RepoFindMeeting(int(MeetingId)))
}
