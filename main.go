package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubernetes-resource-list/config"
	"os"
	"strings"
)

var (
	kubeConfig = pflag.StringP("kubeconfig", "c", "", "(optional) absolute path to the kubeconfig file")
	prefer     = pflag.BoolP("prefer", "p", false, "(optional) only display the supported resources with the version preferred by the server.")
	search     = pflag.StringP("search", "s", "", "(optional) only display the supported resources for a group and version.")
	version    = pflag.BoolP("version", "v", false, "show app version")
	Version    string
	Build      string
)

func main() {
	pflag.Parse()

	// init kubeconfig file
	c := config.InitKubeConfig(kubeConfig)

	// show app info
	showInfo()

	discoveryClient := newDiscoveryClient(*c)

	if *search == "" {
		getResources(discoveryClient)
	} else {
		getServerResourcesByGroupVersion(discoveryClient, *search)
	}

}

func newDiscoveryClient(kubeconfig string) *discovery.DiscoveryClient {
	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientSet, err := kubernetes.NewForConfig(cfg)
	return clientSet.DiscoveryClient
}

func getResources(client *discovery.DiscoveryClient) {
	var APIResourceList []*v1.APIResourceList
	var err error
	if *prefer {
		APIResourceList, err = client.ServerPreferredResources()
	} else {
		_, APIResourceList, err = client.ServerGroupsAndResources()
	}

	if err != nil {
		panic(err.Error())
	}

	tableHeader()

	for _, v := range APIResourceList {
		for _, y := range v.APIResources {
			fmt.Printf("%-25s\t", getApiGroup(v.GroupVersion))
			fmt.Printf("%-20s\t", getApiGroupVersion(v.GroupVersion))
			fmt.Printf("%-35s\t", y.Name)
			fmt.Printf("%-s\n", y.Verbs)
		}
	}
}

func getServerResourcesByGroupVersion(client *discovery.DiscoveryClient, groupVersion string) {
	APIResourceList, err := client.ServerResourcesForGroupVersion(groupVersion)
	if err != nil {
		panic(err.Error())
	}

	searchTableHeader()

	for _, v := range APIResourceList.APIResources {
		fmt.Printf("%-35s\t", v.Name)
		fmt.Printf("%-s\n", v.Verbs)
	}

}

func getApiGroup(groupVersion string) string {
	if strings.Contains(groupVersion, "/") {
		return strings.Split(groupVersion, "/")[0]
	}
	return "core"

}

func getApiGroupVersion(groupVersion string) string {
	if strings.Contains(groupVersion, "/") {
		return strings.Split(groupVersion, "/")[1]
	}
	return groupVersion
}

func tableHeader() {
	fmt.Printf("%-25s\t%-20s\t%-35s\t%-s\n", "APIGROUPS", "GROUPVERSION", "NAME", "VERBS")
}

func searchTableHeader() {
	fmt.Printf("%-35s\t%-s\n", "NAME", "VERBS")
}

func showInfo() {
	if *version {
		showAppInfo()
		os.Exit(0)
	}
}

func showAppInfo() {
	fmt.Printf("Version:\t %s\n", Version)
	fmt.Printf("Build:\t %s\n", Build)

}
