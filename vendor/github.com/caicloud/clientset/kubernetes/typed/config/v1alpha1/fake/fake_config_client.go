/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/caicloud/clientset/kubernetes/typed/config/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigV1alpha1 struct {
	*testing.Fake
}

func (c *FakeConfigV1alpha1) ConfigClaims(namespace string) v1alpha1.ConfigClaimInterface {
	return &FakeConfigClaims{c, namespace}
}

func (c *FakeConfigV1alpha1) ConfigReferences(namespace string) v1alpha1.ConfigReferenceInterface {
	return &FakeConfigReferences{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}