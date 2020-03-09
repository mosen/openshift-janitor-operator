package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

// SweepSpec defines the desired state of Sweep, which is a single scan of all Projects/Namespaces for Old and Unused
// Resources.
type SweepSpec struct {
	IgnoreProjects    []string          `json:"ignoreProjects,omitempty"`
	IgnoreAnnotations map[string]string `json:"ignoreAnnotation,omitempty"`

	// +kubebuilder:validation:Minimum=1
	WarnAgeDays int `json:"warnAgeDays"`

	// +kubebuilder:validation:Minimum=1
	DeleteAgeDays int `json:"deleteAgeDays"`
}

// SweepStatus defines the state of a Sweep operation
type SweepStatus struct {
	Active bool `json:"active"`

	Started  *metav1.Time `json:"started,omitempty"`
	Finished *metav1.Time `json:"finished,omitempty"`

	ProjectsDeleted []string `json:"projectsDeleted,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sweep is the Schema for the sweeps API.
// A "Sweep" is a single operation which scans the entire cluster for outdated/unused Projects
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
