package main

import (
	"github.com/bwmarrin/discordgo"
	tea "github.com/charmbracelet/bubbletea"
  "fmt"
)

type GuildsNavigation struct {
  Cursor [2]int
  IsOnServerTab bool
  Discord *discordgo.Session
}

type Guild struct {
  Channels []Channel
  Name string

}

type Channel struct {
  Id string
  Name string
  Messages []string
}

var Guilds []Guild

func (m GuildsNavigation) Init() tea.Cmd {
  m.Discord.Open()
  fmt.Print("\033[H\033[2J")

  for _, guildId := range config.AllowedList {
    guild, _ := m.Discord.State.Guild(guildId)
    var channels []Channel

    for _, channel := range guild.Channels {
      channels = append(channels, Channel{
        Id: channel.ID,
        Name: channel.Name,
      })
    }
    Guilds = append(Guilds, Guild{
      Channels: channels,
      Name: guild.Name,
    })
  }
  return nil
}

func (m GuildsNavigation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c":
      return m, tea.Quit
    }
  }
  return m, nil
}

func (m GuildsNavigation) View() string {
  guilds := ""

  for _, guild := range Guilds {
    guilds += guild.Name + "\n"
  }

  return guilds
}
