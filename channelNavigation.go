package main

import (
	"github.com/bwmarrin/discordgo"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
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
      if channel.Type != 0 { continue }
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

    case "up", "k", "w":
      if m.Cursor[0] > 0 && m.IsOnServerTab {
        m.Cursor[0] -= 1
        m.Cursor[1] = 0
      } else if m.Cursor[1] > 0 && !m.IsOnServerTab {
        m.Cursor[1] -= 1
      }

    case "down", "j", "s":
      if m.Cursor[0] < len(Guilds) - 1 && m.IsOnServerTab {
        m.Cursor[0] += 1
        m.Cursor[1] = 0
      } else if m.Cursor[1] < len(Guilds[m.Cursor[0]].Channels) && !m.IsOnServerTab {
        m.Cursor[1] += 1
      }

    case "right", "l", "d":
      m.IsOnServerTab = false

    case "left", "h", "a":
      m.IsOnServerTab = true
      m.Cursor[1] = 0

    }
  }
  return m, nil
}

func (m GuildsNavigation) View() string {
  var guilds []string
  var channels []string

  for i, guild := range Guilds {
    if i == m.Cursor[0] {
      if m.IsOnServerTab {
        guilds = append(guilds, selected.Render(guild.Name))
      } else {
        guilds = append(guilds, highlighted.Render(guild.Name))
      }
    } else {
      guilds = append(guilds, guild.Name)
    }
  }

  for i, channel := range Guilds[m.Cursor[0]].Channels {
    if i == m.Cursor[1] {
      if !m.IsOnServerTab {
        channels = append(channels, selected.Render(channel.Name))
      } else {
        channels = append(channels, highlighted.Render(channel.Name))
      }
    } else {
      channels = append(channels, channel.Name)
    }
  }

  return lipgloss.JoinHorizontal(
    lipgloss.Top,
    lipgloss.JoinVertical(lipgloss.Left, guilds...),
    lipgloss.JoinVertical(lipgloss.Left, channels...),
  );
}
