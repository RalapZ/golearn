package mock

type Config interface {
	ReadConfig(string) *Config
}
