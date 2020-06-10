package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Service struct {
	clientSet *kubernetes.Clientset
}

func NewService(clientSet *kubernetes.Clientset) *Service {
	return &Service{clientSet: clientSet}
}

func (c *Service) Create(namespace string, service *v1.Service) (err error) {
	service, err = c.clientSet.CoreV1().Services(namespace).Create(context.TODO(), service, meta_v1.CreateOptions{})
	return
}

func (c *Service) Delete(namespace, name string) (err error) {
	ops := meta_v1.DeleteOptions{}
	err = c.clientSet.CoreV1().Services(namespace).Delete(context.TODO(), name, ops)
	return
}

func (c *Service) Update(namespace string, service *v1.Service) (err error) {
	service, err = c.clientSet.CoreV1().Services(namespace).Update(context.TODO(), service, meta_v1.UpdateOptions{})
	return
}

func (c *Service) Get(name, namespace string) (service *v1.Service, err error) {
	ops := meta_v1.GetOptions{}
	service, err = c.clientSet.CoreV1().Services(namespace).Get(context.TODO(), name, ops)
	return
}

func (c *Service) List(namespace string, ops meta_v1.ListOptions) (services *v1.ServiceList, err error) {
	services, err = c.clientSet.CoreV1().Services(namespace).List(context.TODO(), ops)
	return
}
