package command

import (
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/command"
	"github.com/spf13/cobra"
)

const name = "stage"

type Command struct {
	cmd *cobra.Command
}

func (p *Command) Type() plugin.Type {
	return command.Type
}

func (p *Command) Init() error {
	p.cmd = NewStageCommand()
	return nil
}

func (p *Command) Name() string {
	return name
}

func (p *Command) GetCobraCommand() *cobra.Command {
	return p.cmd
}
