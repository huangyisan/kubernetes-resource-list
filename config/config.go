package config

import (
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
)

func InitKubeConfig() *string {
	var kubeConfig *string
	kubeConfig = pflag.StringP("kubeconfig", "c", "", "(optional) absolute path to the kubeconfig file")
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
