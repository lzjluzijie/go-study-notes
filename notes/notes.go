package notes

import (
	"github.com/lzjluzijie/go-study-notes/notes/bitcoin"
	"github.com/urfave/cli"
)

var commends []cli.Command

func init() {
	NewCommand(cli.Command{
		Name:    "bitcoin",
		Aliases: []string{"btc"},
		Usage:   "Something about bitcoin.",
		Subcommands: []cli.Command{
			{
				Name:    "new",
				Aliases: []string{"new"},
				Usage:   "Generate a new bitcoin wallet.",
				Action:  bitcoin.NewWallet,
			},
		},
	})
}

func NewCommand(c cli.Command) {
	commends = append(commends, c)
}

func GetCommands() []cli.Command {
	return commends
}
