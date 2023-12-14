package libs

import (
	"regexp"
	"strings"
)

var lists []ICommand

func NewCommands(cmd *ICommand) {
	lists = append(lists, *cmd)
}

func GetList() []ICommand {
	return lists
}

func Get(c *NewClientImpl, m *IMessage) {
	prefix := "#"
	for _, cmd := range lists {
		re := regexp.MustCompile(`^` + cmd.Name + `$`)
		if reg := len(re.FindAllString(strings.ReplaceAll(m.Command, prefix, ""), -1)) > 0; reg {
			var cmdWithPref bool
			var cmdWithoutPref bool

			if cmd.IsPrefix && strings.HasPrefix(m.Command, prefix) {
				cmdWithPref = true
			} else {
				cmdWithPref = false
			}

			if !cmd.IsPrefix {
				cmdWithoutPref = true
			} else {
				cmdWithoutPref = false
			}

			if !cmdWithPref && !cmdWithoutPref {
				continue
			}

			//Checking
			if cmd.IsOwner && !m.IsOwner {
				continue
			}

			if cmd.IsMedia && m.Media == nil {
				m.Reply("Media Di Butuhkan")
				continue
			}

			if cmd.IsQuerry && m.Querry == "" {
				m.Reply("Querry Di Butuhkan")
				continue
			}

			if cmd.IsGroup && !m.IsGroup {
				m.Reply("Hanya Khusus Group")
				continue
			}

			// if cmd.IsAdmin && (m.IsGroup && !m.IsAdmin) {
			// 	m.Reply("Akses Admin Di Butuhkan.")
			// 	continue
			// }

			cmd.Exec(c, m)
		}
	}
}
