/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/keshiha96/wso2am-k8s-controller/pkg/apis/apim/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApimanagerLister helps list Apimanagers.
type ApimanagerLister interface {
	// List lists all Apimanagers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Apimanager, err error)
	// Apimanagers returns an object that can list and get Apimanagers.
	Apimanagers(namespace string) ApimanagerNamespaceLister
	ApimanagerListerExpansion
}

// apimanagerLister implements the ApimanagerLister interface.
type apimanagerLister struct {
	indexer cache.Indexer
}

// NewApimanagerLister returns a new ApimanagerLister.
func NewApimanagerLister(indexer cache.Indexer) ApimanagerLister {
	return &apimanagerLister{indexer: indexer}
}

// List lists all Apimanagers in the indexer.
func (s *apimanagerLister) List(selector labels.Selector) (ret []*v1alpha1.Apimanager, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Apimanager))
	})
	return ret, err
}

// Apimanagers returns an object that can list and get Apimanagers.
func (s *apimanagerLister) Apimanagers(namespace string) ApimanagerNamespaceLister {
	return apimanagerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApimanagerNamespaceLister helps list and get Apimanagers.
type ApimanagerNamespaceLister interface {
	// List lists all Apimanagers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Apimanager, err error)
	// Get retrieves the Apimanager from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Apimanager, error)
	ApimanagerNamespaceListerExpansion
}

// apimanagerNamespaceLister implements the ApimanagerNamespaceLister
// interface.
type apimanagerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Apimanagers in the indexer for a given namespace.
func (s apimanagerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Apimanager, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Apimanager))
	})
	return ret, err
}

// Get retrieves the Apimanager from the indexer for a given namespace and name.
func (s apimanagerNamespaceLister) Get(name string) (*v1alpha1.Apimanager, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanager"), name)
	}
	return obj.(*v1alpha1.Apimanager), nil
}