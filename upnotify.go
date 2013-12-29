//Upnotify.go
//
//A simple app for sending a push notification via pushover.net

package main

import "net/url"
import "net/http"
import "io/ioutil"
import "encoding/json"

type PushoverMessage struct {
  Token   string `json:"token"`
  User    string `json:"user"`
  Message string `json:"message"`
}

func LoadConfig() (PushoverMessage, error){
  var msg PushoverMessage
  
  config, err := ioutil.ReadFile("upnotify.conf")
  if (err != nil){
    return msg, err
  }

  err = json.Unmarshal(config, &msg)
  if (err != nil){
    return msg, err
  }

  return msg, err
}

func main() {
  msg, err := LoadConfig();
  if(err != nil){
    return
  }

  http.PostForm("http://api.pushover.net/1/messages.json",
                url.Values{"token": {msg.Token},
                           "user":    {msg.User},
                           "message": {msg.Message}})
}

