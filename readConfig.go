package main

import (
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  Token string 
  Whitelist []string
}

var config Config

func readConfig() error {

  file, err := ioutil.ReadFile("config.json")
  if err != nil {
    return err
  }

  err = json.Unmarshal(file, &config)
  if err != nil {
    return err
  }
  return nil
}
