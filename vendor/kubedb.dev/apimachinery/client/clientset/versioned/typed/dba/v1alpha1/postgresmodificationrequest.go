/*
Copyright The KubeDB Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/dba/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PostgresModificationRequestsGetter has a method to return a PostgresModificationRequestInterface.
// A group's client should implement this interface.
type PostgresModificationRequestsGetter interface {
	PostgresModificationRequests() PostgresModificationRequestInterface
}

// PostgresModificationRequestInterface has methods to work with PostgresModificationRequest resources.
type PostgresModificationRequestInterface interface {
	Create(*v1alpha1.PostgresModificationRequest) (*v1alpha1.PostgresModificationRequest, error)
	Update(*v1alpha1.PostgresModificationRequest) (*v1alpha1.PostgresModificationRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PostgresModificationRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.PostgresModificationRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresModificationRequest, err error)
	PostgresModificationRequestExpansion
}

// postgresModificationRequests implements PostgresModificationRequestInterface
type postgresModificationRequests struct {
	client rest.Interface
}

// newPostgresModificationRequests returns a PostgresModificationRequests
func newPostgresModificationRequests(c *DbaV1alpha1Client) *postgresModificationRequests {
	return &postgresModificationRequests{
		client: c.RESTClient(),
	}
}

// Get takes name of the postgresModificationRequest, and returns the corresponding postgresModificationRequest object, and an error if there is any.
func (c *postgresModificationRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.PostgresModificationRequest, err error) {
	result = &v1alpha1.PostgresModificationRequest{}
	err = c.client.Get().
		Resource("postgresmodificationrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PostgresModificationRequests that match those selectors.
func (c *postgresModificationRequests) List(opts v1.ListOptions) (result *v1alpha1.PostgresModificationRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PostgresModificationRequestList{}
	err = c.client.Get().
		Resource("postgresmodificationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested postgresModificationRequests.
func (c *postgresModificationRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("postgresmodificationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a postgresModificationRequest and creates it.  Returns the server's representation of the postgresModificationRequest, and an error, if there is any.
func (c *postgresModificationRequests) Create(postgresModificationRequest *v1alpha1.PostgresModificationRequest) (result *v1alpha1.PostgresModificationRequest, err error) {
	result = &v1alpha1.PostgresModificationRequest{}
	err = c.client.Post().
		Resource("postgresmodificationrequests").
		Body(postgresModificationRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a postgresModificationRequest and updates it. Returns the server's representation of the postgresModificationRequest, and an error, if there is any.
func (c *postgresModificationRequests) Update(postgresModificationRequest *v1alpha1.PostgresModificationRequest) (result *v1alpha1.PostgresModificationRequest, err error) {
	result = &v1alpha1.PostgresModificationRequest{}
	err = c.client.Put().
		Resource("postgresmodificationrequests").
		Name(postgresModificationRequest.Name).
		Body(postgresModificationRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the postgresModificationRequest and deletes it. Returns an error if one occurs.
func (c *postgresModificationRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("postgresmodificationrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *postgresModificationRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("postgresmodificationrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched postgresModificationRequest.
func (c *postgresModificationRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresModificationRequest, err error) {
	result = &v1alpha1.PostgresModificationRequest{}
	err = c.client.Patch(pt).
		Resource("postgresmodificationrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
