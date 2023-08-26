package cmder

import "github.com/spf13/cobra"

type BaseCmd struct {
	Cmd *cobra.Command
}

func (bc *BaseCmd) GetCommand() *cobra.Command {
	return bc.Cmd
}

func NewBaseCmd(cmd *cobra.Command) *BaseCmd {
	return &BaseCmd{Cmd: cmd}
}
