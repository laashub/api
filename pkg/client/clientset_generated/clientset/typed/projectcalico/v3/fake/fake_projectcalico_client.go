// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/tigera/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeProjectcalicoV3 struct {
	*testing.Fake
}

func (c *FakeProjectcalicoV3) GlobalReportTypes(namespace string) v3.GlobalReportTypeInterface {
	return &FakeGlobalReportTypes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectcalicoV3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}