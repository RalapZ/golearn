package model

type IniConfig struct {
	Etcdc
}
type Etcdc struct {
	Address []string `ini:"Address"`
	Timeout int      `ini:"Timeout"`
}
