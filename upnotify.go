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

//Loads a json encoded config from upnotify.conf
//For an example see file upnotify.conf.example
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

  //Pushover do not accept plain json encoding. It has to be either
  //http Form of percent encoded format. This is not indicated in the
  //api docs, but in a blog post. Looks to be a result of a rails vuln.
  http.PostForm("http://api.pushover.net/1/messages.json",
                url.Values{"token": {msg.Token},
                           "user":    {msg.User},
                           "message": {msg.Message}})
}

