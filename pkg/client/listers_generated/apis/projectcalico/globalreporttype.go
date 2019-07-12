// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package projectcalico

import (
	projectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlobalReportTypeLister helps list GlobalReportTypes.
type GlobalReportTypeLister interface {
	// List lists all GlobalReportTypes in the indexer.
	List(selector labels.Selector) (ret []*projectcalico.GlobalReportType, err error)
	// GlobalReportTypes returns an object that can list and get GlobalReportTypes.
	GlobalReportTypes(namespace string) GlobalReportTypeNamespaceLister
	GlobalReportTypeListerExpansion
}

// globalReportTypeLister implements the GlobalReportTypeLister interface.
type globalReportTypeLister struct {
	indexer cache.Indexer
}

// NewGlobalReportTypeLister returns a new GlobalReportTypeLister.
func NewGlobalReportTypeLister(indexer cache.Indexer) GlobalReportTypeLister {
	return &globalReportTypeLister{indexer: indexer}
}

// List lists all GlobalReportTypes in the indexer.
func (s *globalReportTypeLister) List(selector labels.Selector) (ret []*projectcalico.GlobalReportType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*projectcalico.GlobalReportType))
	})
	return ret, err
}

// GlobalReportTypes returns an object that can list and get GlobalReportTypes.
func (s *globalReportTypeLister) GlobalReportTypes(namespace string) GlobalReportTypeNamespaceLister {
	return globalReportTypeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlobalReportTypeNamespaceLister helps list and get GlobalReportTypes.
type GlobalReportTypeNamespaceLister interface {
	// List lists all GlobalReportTypes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*projectcalico.GlobalReportType, err error)
	// Get retrieves the GlobalReportType from the indexer for a given namespace and name.
	Get(name string) (*projectcalico.GlobalReportType, error)
	GlobalReportTypeNamespaceListerExpansion
}

// globalReportTypeNamespaceLister implements the GlobalReportTypeNamespaceLister
// interface.
type globalReportTypeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlobalReportTypes in the indexer for a given namespace.
func (s globalReportTypeNamespaceLister) List(selector labels.Selector) (ret []*projectcalico.GlobalReportType, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*projectcalico.GlobalReportType))
	})
	return ret, err
}

// Get retrieves the GlobalReportType from the indexer for a given namespace and name.
func (s globalReportTypeNamespaceLister) Get(name string) (*projectcalico.GlobalReportType, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(projectcalico.Resource("globalreporttype"), name)
	}
	return obj.(*projectcalico.GlobalReportType), nil
}