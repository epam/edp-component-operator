package v1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EDPComponentSpec defines the desired state of EDPComponent
type EDPComponentSpec struct {
	// specifies a type of component, e.g. 'nexus', 'gerrit', etc.
	Type string `json:"type"`

	// specifies a link to component
	Url string `json:"url"`

	// base64 encoded SVG icon of a component
	Icon string `json:"icon"`

	// specifies whether a component is visible
	Visible bool `json:"visible"`
}

// EDPComponentStatus defines the observed state of EDPComponent
type EDPComponentStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:path=edpcomponents,scope=Namespaced

// EDPComponent is the Schema for the edpcomponents API
type EDPComponent struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EDPComponentSpec   `json:"spec,omitempty"`
	Status EDPComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EDPComponentList contains a list of EDPComponent
type EDPComponentList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`
	Items           []EDPComponent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EDPComponent{}, &EDPComponentList{})
}
