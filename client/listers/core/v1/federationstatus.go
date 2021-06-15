// Copyright Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "maistra.io/api/core/v1"
)

// FederationStatusLister helps list FederationStatuses.
// All objects returned here must be treated as read-only.
type FederationStatusLister interface {
	// List lists all FederationStatuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.FederationStatus, err error)
	// FederationStatuses returns an object that can list and get FederationStatuses.
	FederationStatuses(namespace string) FederationStatusNamespaceLister
	FederationStatusListerExpansion
}

// federationStatusLister implements the FederationStatusLister interface.
type federationStatusLister struct {
	indexer cache.Indexer
}

// NewFederationStatusLister returns a new FederationStatusLister.
func NewFederationStatusLister(indexer cache.Indexer) FederationStatusLister {
	return &federationStatusLister{indexer: indexer}
}

// List lists all FederationStatuses in the indexer.
func (s *federationStatusLister) List(selector labels.Selector) (ret []*v1.FederationStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.FederationStatus))
	})
	return ret, err
}

// FederationStatuses returns an object that can list and get FederationStatuses.
func (s *federationStatusLister) FederationStatuses(namespace string) FederationStatusNamespaceLister {
	return federationStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederationStatusNamespaceLister helps list and get FederationStatuses.
// All objects returned here must be treated as read-only.
type FederationStatusNamespaceLister interface {
	// List lists all FederationStatuses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.FederationStatus, err error)
	// Get retrieves the FederationStatus from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.FederationStatus, error)
	FederationStatusNamespaceListerExpansion
}

// federationStatusNamespaceLister implements the FederationStatusNamespaceLister
// interface.
type federationStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederationStatuses in the indexer for a given namespace.
func (s federationStatusNamespaceLister) List(selector labels.Selector) (ret []*v1.FederationStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.FederationStatus))
	})
	return ret, err
}

// Get retrieves the FederationStatus from the indexer for a given namespace and name.
func (s federationStatusNamespaceLister) Get(name string) (*v1.FederationStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("federationstatus"), name)
	}
	return obj.(*v1.FederationStatus), nil
}