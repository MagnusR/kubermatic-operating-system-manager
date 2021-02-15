/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

package fake

import (
	v1alpha1 "k8c.io/operating-system-manager/pkg/crd/client/clientset/versioned/typed/osm/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOperatingsystemmanagerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOperatingsystemmanagerV1alpha1) OperatingSystemConfigs(namespace string) v1alpha1.OperatingSystemConfigInterface {
	return &FakeOperatingSystemConfigs{c, namespace}
}

func (c *FakeOperatingsystemmanagerV1alpha1) OperatingSystemProfiles(namespace string) v1alpha1.OperatingSystemProfileInterface {
	return &FakeOperatingSystemProfiles{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOperatingsystemmanagerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
