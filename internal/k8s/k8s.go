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
}

func AddBot(key string, login string) {
	clientset.CoreV1().Pods("default").Create(
		context.Background(),
		&v1.Pod{
			TypeMeta: metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{
				Name:      login,
				Namespace: "default",
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{{
					Name:  login,
					Image: "vladimish/client:1.0",
					Ports: []v1.ContainerPort{
						{
							Name:          "",
							ContainerPort: 1721,
							Protocol:      "TCP",
						},
					},
					Env: []v1.EnvVar{
						{
							Name:  "DB_USER",
							Value: "root",
						},
						{
							Name:  "DB_PASS",
							Value: "root",
						},
						{
							Name:  "DB_ADDR",
							Value: "localhost:3306",
						},
						{
							Name:  "TG_KEY",
							Value: key,
						},
					},
				}},
				RestartPolicy: "Always",
			},
		},
		metav1.CreateOptions{},
	)
}
