package main

import(
  "fmt"
  "time"
  )

var currentId int

var meetings Meetings


func init() { //TESTING SEED
    RepoCreateMeeting(Meeting{Title: "Write presentation", })
    RepoCreateMeeting(Meeting{Title: "Host meetup"})
}

func RepoFindMeeting(id int) Meeting { // FINDS MEETING WITH MEETING ID
    for _, t := range meetings {
        if t.Id == id {
            return t
        }
    }
    return Meeting{}
}

func RepoMeetingTime(start_time string, end_time string) Meetings{ // FINDING MEETINGS BETWEEN A TIME PERIOD
  start, err := time.Parse(time.RFC3339, start_time)
  end, err := time.Parse(time.RFC3339, end_time)
  a := Meetings{}
  for _, t := range meetings {
      temp := t.Start
      if (temp.Before(end) && temp.After(start)) || temp.Equal(start) || temp.Equal(end) {
          a = append(a, t)
      }
  }
    _ = err
  return a
}

func RepoMeetingTime2(start_time string, end_time string, email string) Meetings{ // FINDING MEETINGS BETWEEN A TIME PERIOD FOR A PARTICIPANT
  start, err := time.Parse(time.RFC3339, start_time)
  end, err := time.Parse(time.RFC3339, end_time)
  a := Meetings{}
  for _, t := range meetings {
      temp := t.Start
      if (temp.Before(end) && temp.After(start)) || temp.Equal(start) || temp.Equal(end) {
          for _, i :=range t.Participants{
            if i.Email == email{
              a = append(a, t)
            }
         }
      }
  }
    _ = err
  return a
}

func RepoMeetingTime3(email string) Meetings{ // FINDING MEETINGS FOR A PARTICIPANT
  a:= Meetings{}
  for _, t := range meetings{
    for _, i := range t.Participants{
      if i.Email == email{
        a = append(a, t)
      }
   }
}
return a
}
func RepoCreateMeeting(t Meeting) (Meeting, string) { // CREATES A MEETING WITHOUT ANY CONFLICTS
    if (t.Title == "" || t.Start.Equal(t.End) || t.Participants == nil){//CONSTRAINT CHECKING
      fmt.Println("Not Enough Parameters")
      return t, "false"
    }
    for _, p := range t.Participants{
      if p.Email == "" || p.Name =="" || (p.RSVP != "Yes" && p.RSVP != "NO" && p.RSVP != "MayBe" && p.RSVP != "Not Answered"){// CONSTRAINT CHECKING
        return t, "false"
      }
    }
    for _, i := range meetings{
        if (i.Start.Before(t.Start) && t.Start.Before(i.End)) || (i.End.After(t.End) && i.Start.Before(t.End)){ // Check if there is a Nested Time Interval
        for _, j := range i.Participants{
          if j.RSVP == "Yes"{ // To check COnflicts if the Participant do not RSVP yes for 2 meetings with nested Time
            for _, p := range t.Participants{
              if (j.Email == p.Email) && (p.RSVP == "Yes"){
                fmt.Println("There is a Conflict for Particpant: ", p.Email)
                fmt.Println("Meeting ID: ", i.Id)
                fmt.Println("Starting from: ", i.Start)
                fmt.Println("Ending at: ", i.End)
                return  i, p.Email//Return False if COnflict Occurs
              }
            }
          }
        }
      }
    }
    currentId += 1
    t.Id = currentId
    meetings = append(meetings, t)
    return t, "true"
}

func RepoDestroyMeeting(id int) error { // DELETE A MEETING
    for i, t := range meetings {
        if t.Id == id {
            meetings = append(meetings[:i], meetings[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Meeting with id of %d to delete", id)
}
