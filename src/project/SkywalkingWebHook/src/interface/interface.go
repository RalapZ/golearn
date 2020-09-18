package main

type Config interface {
	ReadConfig(string) *Config
}
