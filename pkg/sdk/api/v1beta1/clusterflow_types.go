// Copyright © 2019 Banzai Cloud
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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

// ClusterFlow is the Schema for the clusterflows API
type ClusterFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Name of the logging cluster to be attached
	Spec   FlowSpec   `json:"spec,omitempty"`
	Status FlowStatus `json:"status,omitempty"`
}

type ClusterMatch struct {
	*ClusterSelect  `json:"select,omitempty"`
	*ClusterExclude `json:"exclude,omitempty"`
}

type ClusterSelect struct {
	Namespaces []string          `json:"namespaces,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

type ClusterExclude struct {
	Namespaces []string          `json:"namespaces,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

// FlowSpec is the Kubernetes spec for Flows
type ClusterFlowSpec struct {
	Match      []ClusterMatch `json:"match,omitempty"`
	Filters    []Filter       `json:"filters,omitempty"`
	LoggingRef string         `json:"loggingRef,omitempty"`
	OutputRefs []string       `json:"outputRefs"`
}

// +kubebuilder:object:root=true

// ClusterFlowList contains a list of ClusterFlow
type ClusterFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterFlow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterFlow{}, &ClusterFlowList{})
}
