package k8s

import (
	"context"
	"github.com/vladimish/hack-gateway/pkg/log"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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

func AddBot(key string, login string, nodePort int32) error {
	_, err := clientset.CoreV1().Services("default").Create(
		context.Background(),
		&v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name: login + "-service",
			},
			Spec: v1.ServiceSpec{
				Type:     v1.ServiceTypeNodePort,
				Selector: map[string]string{"app": login},
				Ports: []v1.ServicePort{
					{
						Protocol:   "TCP",
						Port:       1721,
						TargetPort: intstr.IntOrString{},
						NodePort:   nodePort,
					},
				},
			},

			Status: v1.ServiceStatus{},
		},
		metav1.CreateOptions{},
	)
	if err != nil {
		return err
	}

	var one = int32(1)
	_, err = clientset.AppsV1().Deployments("default").Create(
		context.Background(),
		&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:   login + "-deployment",
				Labels: map[string]string{"app": login},
			},
			Spec: appsv1.DeploymentSpec{
				Replicas: &one,
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{"app": login},
				},
				Template: v1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{"app": login},
					},
					Spec: v1.PodSpec{
						Containers: []v1.Container{
							{Name: login + "-db",
								Image: "vladimish/client-db:1.0",
								Ports: []v1.ContainerPort{
									{
										HostPort:      30001,
										ContainerPort: 3306,
										Protocol:      "TCP",
									},
								},
								Env: []v1.EnvVar{
									{
										Name:  "MARIADB_ROOT_PASSWORD",
										Value: "root",
									},
								},
							},
							{
								Name:  login,
								Image: "vladimish/client:1.0",
								Ports: []v1.ContainerPort{
									{
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
			},
			Status: appsv1.DeploymentStatus{},
		},
		metav1.CreateOptions{},
	)

	if err != nil {
		return err
	}

	return nil
}
