package cmder

import "github.com/spf13/cobra"

type BaseCmd struct {
	cmd *cobra.Command
}

func (bc *BaseCmd) GetCommand() *cobra.Command {
	return bc.cmd
}

func NewBaseCmd(cmd *cobra.Command) *BaseCmd {
	return &BaseCmd{cmd: cmd}
}
