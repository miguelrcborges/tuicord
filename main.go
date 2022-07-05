package main

import (
  "github.com/bwmarrin/discordgo"
  "fmt"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  err := readConfig()
  if err != nil {
    fmt.Println("Error: ", err)
  }

  discord, err := discordgo.New("Bot " + config.Token)

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  discord.AddHandler(messageCreate)

  fmt.Println(config)

  discord.Identify.Intents = discordgo.IntentsAll

  err = discord.Open()
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  for _, v := range config.Whitelist {
    if m.GuildID == v {
      guild, _ := s.State.Guild(m.GuildID)
      fmt.Println(guild.Name, m.Author, m.Content)
      return 
    }
  }
}
