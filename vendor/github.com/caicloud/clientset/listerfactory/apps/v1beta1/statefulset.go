/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/apps/v1beta1"
)

var _ v1beta1.StatefulSetLister = &statefulSetLister{}

var _ v1beta1.StatefulSetNamespaceLister = &statefulSetNamespaceLister{}

// statefulSetLister implements the StatefulSetLister interface.
type statefulSetLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStatefulSetLister returns a new StatefulSetLister.
func NewStatefulSetLister(client kubernetes.Interface) v1beta1.StatefulSetLister {
	return NewFilteredStatefulSetLister(client, nil)
}

func NewFilteredStatefulSetLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.StatefulSetLister {
	return &statefulSetLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all StatefulSets in the indexer.
func (s *statefulSetLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AppsV1beta1().StatefulSets(v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

func (s *statefulSetLister) GetPodStatefulSets(*corev1.Pod) ([]*appsv1beta1.StatefulSet, error) {
	return nil, nil
}

// StatefulSets returns an object that can list and get StatefulSets.
func (s *statefulSetLister) StatefulSets(namespace string) v1beta1.StatefulSetNamespaceLister {
	return statefulSetNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// statefulSetNamespaceLister implements the StatefulSetNamespaceLister
// interface.
type statefulSetNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all StatefulSets in the indexer for a given namespace.
func (s statefulSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1beta1.StatefulSet, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AppsV1beta1().StatefulSets(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the StatefulSet from the indexer for a given namespace and name.
func (s statefulSetNamespaceLister) Get(name string) (*appsv1beta1.StatefulSet, error) {
	return s.client.AppsV1beta1().StatefulSets(s.namespace).Get(name, v1.GetOptions{})
}
