package k8s

import (
	"context"
	"github.com/vladimish/hack-gateway/pkg/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreatePod(name string, env []corev1.EnvVar) {
	newPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name + "-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: name, Image: name + ":latest", Env: env},
			},
		},
	}

	if clientset == nil {
		log.Get().Error("Client set isn't initialized")
	}
	pod, err := clientset.CoreV1().Pods("default").Create(context.Background(), newPod, metav1.CreateOptions{})
	if err != nil {
		log.Get().Error(err)
	}
	log.Get().Println(pod)
}
