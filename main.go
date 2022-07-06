package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	tea "github.com/charmbracelet/bubbletea"
)

var hasReceivedMessage = false

func main() {
  err := readConfig()
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  discord, err := discordgo.New("Bot " + config.Token)

  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  discord.AddHandler(messageCreate)

  discord.Identify.Intents = discordgo.IntentsAll

  tui := tea.NewProgram(GuildsNavigation{
    Discord: discord,
  })

  if err := tui.Start(); err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  discord.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  for _, v := range config.AllowedList {
    if m.GuildID == v {
      hasReceivedMessage = true
      return 
    }
  }
}
