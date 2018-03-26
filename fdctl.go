package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "flag"
    "fmt"
    "log"
    "net/http"
    "strings"
    "os"
)

var avatars = map[string]struct {
    Username string
    Image    string
}{
    "SLM": {
        Username: "FIREDRILL (SLM)",
        Image:    "https://cdn1.iconfinder.com/data/icons/user-pictures/100/female1-512.png",
    },
    "TECHOPS": {
        Username: "FIREDRILL (TECHOPS)",
        Image:    "https://cdn1.iconfinder.com/data/icons/user-pictures/100/supportmale-512.png",
    },
    "DIRECTOR": {
        Username: "FIREDRILL (DIRECTOR)",
        Image:    "https://cdn1.iconfinder.com/data/icons/user-pictures/101/malecostume-512.png",
    },
    "CONTROL": {
        Username: "FIREDRILL",
        Image: "https://cdn4.iconfinder.com/data/icons/ballicons-2-free/100/match-512.png",
    },
}

type slackMessage struct {
    Text     string `json:"text"`
    UserName string `json:"username"`
    IconURL  string `json:"icon_url"`
}

type slackService struct {
    client   *http.Client
    endpoint string
}

func (s *slackService) Send(msg slackMessage) error {
    b := new(bytes.Buffer)

    if err := json.NewEncoder(b).Encode(msg); err != nil {
        return err
    }

    resp, err := s.client.Post(s.endpoint, "application/json; charset=utf-8", b)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return errors.New("failed to post")
    }

    return nil
}

func main() {
    var operation = flag.String("operation", "say", "start, stop, or say")
    var role = flag.String("role", "control", "role to use")
    var msg = flag.String("message", "", "Message to send")

    flag.Parse()

    *operation = strings.ToUpper(*operation)
    *role = strings.ToUpper(*role)

    switch *operation {
    case "START", "STOP", "SAY":
        // ok
    default:
        log.Fatal("invalid operation")
    }

//    switch *role {
//    case "SLM", "TECHOPS", "DIRECTOR":
        // ok
//    default:
//        log.Fatal("invalid role")
//    }

    slack := &slackService{
        client:   http.DefaultClient,
        endpoint: os.Getenv("FIREDRILL_WEBHOOK_URL"),
    }

    var title = *msg
    if *operation == "START" {
        title = "FIREDRILL START"
    }
    if *operation == "STOP" {
        title = "FIREDRILL STOP"
    }



//    Image := "https://cdn1.iconfinder.com/data/icons/user-pictures/100/male3-512.png"

    settings, ok := avatars[*role]
    if !ok {
        settings.Username = fmt.Sprintf("FIREDRILL (%s)", *role)
        settings.Image = "https://cdn1.iconfinder.com/data/icons/user-pictures/100/male3-512.png"
    }

    err := slack.Send(slackMessage{
        Text:     title,
        UserName: settings.Username,
        IconURL:  settings.Image,
    })

    if err != nil {
        log.Fatal(err)
    }
}
