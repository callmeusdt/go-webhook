package bootstrap

import "flag"

type SubCommand string

const (
	CommandServe   SubCommand = "serve"
	CommandDump    SubCommand = "dump"
	CommandSystemd SubCommand = "systemd"
)

type CliOption struct {
	Config     string     `json:"config" comment:"默认配置文件名"`
	SubCommand SubCommand `json:"sub_command" comment:"子命令"`
}

func ReadCliOption() *CliOption {
	var configName string
	var subCommand SubCommand
	flag.StringVar(&configName, "c", "main", "config file name")
	flag.StringVar(&configName, "config", "main", "config file name")
	flag.StringVar((*string)(&subCommand), "s", string(CommandServe), "subcommand")
	flag.StringVar((*string)(&subCommand), "subcommand", string(CommandServe), "subcommand")
	flag.Parse()

	return CliOption{
		Config:     configName,
		SubCommand: subCommand,
	}.AttachDefault()
}

func (receiver CliOption) AttachDefault() *CliOption {
	//if receiver.Config == "" {
	//	receiver.Config = "main"
	//}

	return &receiver
}
