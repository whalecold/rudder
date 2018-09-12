/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/batch/v1beta1"
)

var _ v1beta1.CronJobLister = &cronJobLister{}

var _ v1beta1.CronJobNamespaceLister = &cronJobNamespaceLister{}

// cronJobLister implements the CronJobLister interface.
type cronJobLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCronJobLister returns a new CronJobLister.
func NewCronJobLister(client kubernetes.Interface) v1beta1.CronJobLister {
	return NewFilteredCronJobLister(client, nil)
}

func NewFilteredCronJobLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.CronJobLister {
	return &cronJobLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all CronJobs in the indexer.
func (s *cronJobLister) List(selector labels.Selector) (ret []*batchv1beta1.CronJob, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.BatchV1beta1().CronJobs(v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// CronJobs returns an object that can list and get CronJobs.
func (s *cronJobLister) CronJobs(namespace string) v1beta1.CronJobNamespaceLister {
	return cronJobNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// cronJobNamespaceLister implements the CronJobNamespaceLister
// interface.
type cronJobNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all CronJobs in the indexer for a given namespace.
func (s cronJobNamespaceLister) List(selector labels.Selector) (ret []*batchv1beta1.CronJob, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.BatchV1beta1().CronJobs(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the CronJob from the indexer for a given namespace and name.
func (s cronJobNamespaceLister) Get(name string) (*batchv1beta1.CronJob, error) {
	return s.client.BatchV1beta1().CronJobs(s.namespace).Get(name, v1.GetOptions{})
}
