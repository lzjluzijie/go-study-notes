package notes

import (
	"github.com/urfave/cli"
)

var commends []cli.Command

func addCommand(c cli.Command) {
	commends = append(commends, c)
}

func GetCommands() []cli.Command {
	return commends
}
