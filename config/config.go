package config

import (
	"os"
	"path/filepath"
)

func InitKubeConfig(kubeConfig *string) *string {
	if *kubeConfig == "" {
		if home := homeDir(); home != "" {
			c := filepath.Join(home, ".kube", "config")
			*kubeConfig = c
		}
	}
	return kubeConfig
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
