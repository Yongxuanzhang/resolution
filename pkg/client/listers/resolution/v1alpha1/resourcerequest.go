/*
Copyright 2022 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/resolution/pkg/apis/resolution/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceRequestLister helps list ResourceRequests.
// All objects returned here must be treated as read-only.
type ResourceRequestLister interface {
	// List lists all ResourceRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceRequest, err error)
	// ResourceRequests returns an object that can list and get ResourceRequests.
	ResourceRequests(namespace string) ResourceRequestNamespaceLister
	ResourceRequestListerExpansion
}

// resourceRequestLister implements the ResourceRequestLister interface.
type resourceRequestLister struct {
	indexer cache.Indexer
}

// NewResourceRequestLister returns a new ResourceRequestLister.
func NewResourceRequestLister(indexer cache.Indexer) ResourceRequestLister {
	return &resourceRequestLister{indexer: indexer}
}

// List lists all ResourceRequests in the indexer.
func (s *resourceRequestLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceRequest))
	})
	return ret, err
}

// ResourceRequests returns an object that can list and get ResourceRequests.
func (s *resourceRequestLister) ResourceRequests(namespace string) ResourceRequestNamespaceLister {
	return resourceRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceRequestNamespaceLister helps list and get ResourceRequests.
// All objects returned here must be treated as read-only.
type ResourceRequestNamespaceLister interface {
	// List lists all ResourceRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceRequest, err error)
	// Get retrieves the ResourceRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceRequest, error)
	ResourceRequestNamespaceListerExpansion
}

// resourceRequestNamespaceLister implements the ResourceRequestNamespaceLister
// interface.
type resourceRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceRequests in the indexer for a given namespace.
func (s resourceRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceRequest))
	})
	return ret, err
}

// Get retrieves the ResourceRequest from the indexer for a given namespace and name.
func (s resourceRequestNamespaceLister) Get(name string) (*v1alpha1.ResourceRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcerequest"), name)
	}
	return obj.(*v1alpha1.ResourceRequest), nil
}