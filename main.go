package main

import (
	"context"
	"fmt"
	"os"
	monitoringclient "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	mfc "github.com/manifestival/client-go-client"
)




func main() {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	mclient, err := monitoringclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	promRules, err := mclient.MonitoringV1().PrometheusRules("tks-system").List(context.TODO(), v1.ListOptions{})
	for _ , rule := range promRules.Items{
		fmt.Println(rule.Name)
	}
}

func ApplyManifest(manisfestPath string) {
	client, err := GetKubeClientFromKubeConfig()
	if err != nil {
		//panic("Failed to load manifest")
		panic("Unable get kube credintials")
	}
	manifest, err := mfc.NewManifest(manisfestPath, client)
	if err != nil {
		panic("Failed to load manifest")
	}
	err = manifest.Apply()
	if err != nil {
		panic("Failed apply manifest")
	}
}

func GetKubeClientFromKubeConfig() (*rest.Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath != "" {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, fmt.Errorf("unable to load kubeconfig from %s: %v", kubeconfigPath, err)
		}
		//kubeconfig = config
		return config, nil
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to load in-cluster config: %v", err)
	}

	return config, nil
}
