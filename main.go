package main

import (
	"context"
	"fmt"
	"os"

	monitoringclient "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	mclient, err := monitoringclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	promRules, err := mclient.MonitoringV1().PrometheusRules("tks-system").List(context.TODO(), v1.ListOptions{})
	for _, rule := range promRules.Items {
		fmt.Println(rule.Name)
	}
}
