package cmder

import "github.com/spf13/cobra"

type ICommand interface {
	GetCommand() *cobra.Command
}

type CommandBuilder struct {
	rootCmd  *cobra.Command
	commands []ICommand
}

func (cb *CommandBuilder) AddCommands(commands ...ICommand) *CommandBuilder {
	cb.commands = append(cb.commands, commands...)
	return cb
}

func NewCommandBuilder(rootCmd *cobra.Command) *CommandBuilder {
	return &CommandBuilder{rootCmd: rootCmd}
}

func (cb *CommandBuilder) Build() error {
	for _, cmd := range cb.commands {
		cb.rootCmd.AddCommand(cmd.GetCommand())
	}
	return cb.rootCmd.Execute()
}
