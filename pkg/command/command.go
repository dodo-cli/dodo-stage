package command

import (
	api "github.com/dodo-cli/dodo-core/api/v1alpha1"
	"github.com/dodo-cli/dodo-core/pkg/plugin"
	"github.com/dodo-cli/dodo-core/pkg/plugin/command"
	"github.com/spf13/cobra"
)

const name = "stage"

var _ command.Command = &Command{}

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

func (p *Command) PluginInfo() (*api.PluginInfo, error) {
	return &api.PluginInfo{Name: name}, nil
}

func (p *Command) GetCobraCommand() *cobra.Command {
	return p.cmd
}
