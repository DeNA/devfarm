package subcmd

type CommandTable map[string]SubCommandDef

func GetSubCommandDef(name string, table CommandTable) (command *SubCommandDef, err error) {
	if cmd, ok := table[name]; ok {
		return &cmd, nil
	}

	return nil, CommandMissingError(name, table)
}
