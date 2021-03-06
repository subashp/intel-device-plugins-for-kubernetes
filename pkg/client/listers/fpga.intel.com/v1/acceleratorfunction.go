// Copyright 2018 Intel Corporation. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/intel/intel-device-plugins-for-kubernetes/pkg/apis/fpga.intel.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AcceleratorFunctionLister helps list AcceleratorFunctions.
type AcceleratorFunctionLister interface {
	// List lists all AcceleratorFunctions in the indexer.
	List(selector labels.Selector) (ret []*v1.AcceleratorFunction, err error)
	// AcceleratorFunctions returns an object that can list and get AcceleratorFunctions.
	AcceleratorFunctions(namespace string) AcceleratorFunctionNamespaceLister
	AcceleratorFunctionListerExpansion
}

// acceleratorFunctionLister implements the AcceleratorFunctionLister interface.
type acceleratorFunctionLister struct {
	indexer cache.Indexer
}

// NewAcceleratorFunctionLister returns a new AcceleratorFunctionLister.
func NewAcceleratorFunctionLister(indexer cache.Indexer) AcceleratorFunctionLister {
	return &acceleratorFunctionLister{indexer: indexer}
}

// List lists all AcceleratorFunctions in the indexer.
func (s *acceleratorFunctionLister) List(selector labels.Selector) (ret []*v1.AcceleratorFunction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AcceleratorFunction))
	})
	return ret, err
}

// AcceleratorFunctions returns an object that can list and get AcceleratorFunctions.
func (s *acceleratorFunctionLister) AcceleratorFunctions(namespace string) AcceleratorFunctionNamespaceLister {
	return acceleratorFunctionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AcceleratorFunctionNamespaceLister helps list and get AcceleratorFunctions.
type AcceleratorFunctionNamespaceLister interface {
	// List lists all AcceleratorFunctions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AcceleratorFunction, err error)
	// Get retrieves the AcceleratorFunction from the indexer for a given namespace and name.
	Get(name string) (*v1.AcceleratorFunction, error)
	AcceleratorFunctionNamespaceListerExpansion
}

// acceleratorFunctionNamespaceLister implements the AcceleratorFunctionNamespaceLister
// interface.
type acceleratorFunctionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AcceleratorFunctions in the indexer for a given namespace.
func (s acceleratorFunctionNamespaceLister) List(selector labels.Selector) (ret []*v1.AcceleratorFunction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AcceleratorFunction))
	})
	return ret, err
}

// Get retrieves the AcceleratorFunction from the indexer for a given namespace and name.
func (s acceleratorFunctionNamespaceLister) Get(name string) (*v1.AcceleratorFunction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("acceleratorfunction"), name)
	}
	return obj.(*v1.AcceleratorFunction), nil
}
