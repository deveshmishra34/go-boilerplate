package config

import "github.com/deveshmishra34/groot/pkg/config/features"

func InitEnv() {
	Env.Init()
}

func InitFeatures() {
	Env.InitFeatures()
}

func ResolveDevMode() {
	Env.CheckAndSetDevMode()
}

func ResolveFlags() {
	Env.OverrideUsingFlags()
}

func PrintEnvInEnvMode() {
	if !EnvModeFlag {
		return
	}
	Env.PrintEnvironment()
}

func SetServiceName(name string) {
	Env.ServiceName = name
}

func SetServiceVersion(version string) {
	Env.Version = version
}

func Feature(name string) *features.Feature {
	return Env.Features.GetFeatureByName(name)
}

func OverrideLoggerUsingFlags() {
	Env.OverrideLoggerUsingFlags()
}
