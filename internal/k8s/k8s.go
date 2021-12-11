package k8s

import (
	"context"
	"github.com/vladimish/hack-gateway/pkg/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

func InitKube() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubecfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	cfg, err := kubecfg.ClientConfig()
	if err != nil {
		log.Get().Error(err)
	}

	clientset = kubernetes.NewForConfigOrDie(cfg)
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Get().Error(err)
	}

	for _, n := range nodeList.Items {
		log.Get().Println(n.Name)
	}

	CreatePod("mysql", []v1.EnvVar{{
		Name:      "MYSQL_ROOT_PASSWORD",
		Value:     "root",
		ValueFrom: nil,
	}})
}
