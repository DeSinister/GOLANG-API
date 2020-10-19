package main

type Participant struct{ // MODEL FOR PARTICIPANT
  Name string `json:"Name" bson:"Name"`
  Email string `json:"email, unique" bson:"email, unique"`
  RSVP string `json:"RSVP" bson:"RSVP"`
}


type Participants []Participant
