package proverbs

import (
	"discord-bot/configs"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func Proverbs(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Bot cannot use your own command
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmds := strings.Split(m.Content, " ")
	
	if cmds[0] != configs.GetPrefix() {
		return
	}

	// There is no other command
	if len(cmds) == 1 {
		s.ChannelMessageSend(m.ChannelID, "Please enter a valid command. You can use 'gobot help' if you have any questions")
		return
	}

	if cmds[1] == "proverbs" {
		rand.Seed(time.Now().UnixNano())
		selection := rand.Intn(len(proverbs))

		_, err := s.ChannelMessageSend(m.ChannelID, proverbs[selection])
		if err != nil {
			fmt.Println(err)
		}
	}
}