/*
Copyright 2023 stackzoo.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodBusterSpec defines the desired state of PodBuster
type PodBusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PodBuster. Edit podbuster_types.go to remove/update
	Namespace string `json:"namespace,omitempty"`
}

// PodBusterStatus defines the observed state of PodBuster
type PodBusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Ok bool `json:"ok,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodBuster is the Schema for the podbusters API
type PodBuster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodBusterSpec   `json:"spec,omitempty"`
	Status PodBusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodBusterList contains a list of PodBuster
type PodBusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodBuster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodBuster{}, &PodBusterList{})
}
