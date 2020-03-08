package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SweepSpec defines the desired state of Sweep, which is a single scan of all Projects/Namespaces for Old and Unused
// Resources.
type SweepSpec struct {
	// +kubebuilder:validation:MinItems=1
	IgnoreProjects []string `json:"ignore,omitempty"`

	// +kubebuilder:validation:MinLength=1
	IgnoreAnnotation string `json:"ignoreAnnotation,omitempty"`

	// +kubebuilder:validation:Minimum=1
	MaximumAgeDays int `json:"maximumAgeDays"`
}

// SweepStatus defines the observed state of Sweep
type SweepStatus struct {
	Active bool `json:"active"`

	Started  *metav1.Time `json:"started,omitempty"`
	Finished *metav1.Time `json:"finished,omitempty"`

	ProjectsDeleted []string `json:"projectsDeleted,omitempty"`

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sweep is the Schema for the sweeps API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sweeps,scope=Namespaced
type Sweep struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SweepSpec   `json:"spec,omitempty"`
	Status SweepStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SweepList contains a list of Sweep
type SweepList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sweep `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sweep{}, &SweepList{})
}
