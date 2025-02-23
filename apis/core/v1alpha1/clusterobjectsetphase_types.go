package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ClusterObjectSetPhase is an internal API, allowing a ClusterObjectSet to delegate a single phase to another custom controller.
// ClusterObjectSets will create subordinate ClusterObjectSetPhases when `.class` is set within the phase specification.
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type ClusterObjectSetPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterObjectSetPhaseSpec   `json:"spec,omitempty"`
	Status ClusterObjectSetPhaseStatus `json:"status,omitempty"`
}

// ClusterObjectSetPhaseList contains a list of ClusterObjectSetPhases.
// +kubebuilder:object:root=true
type ClusterObjectSetPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterObjectSetPhase `json:"items"`
}

// ClusterObjectSetPhaseSpec defines the desired state of a ClusterObjectSetPhase.
type ClusterObjectSetPhaseSpec struct {
	// Specifies the lifecycle state of the ClusterObjectSetPhase.
	// +kubebuilder:default="Active"
	// +kubebuilder:validation:Enum=Active;Paused;Archived
	LifecycleState ObjectSetLifecycleState `json:"lifecycleState,omitempty"`

	// Immutable fields below

	// Revision of the parent ObjectSet to use during object adoption.
	Revision int64 `json:"revision"`

	// Previous revisions of the ClusterObjectSet to adopt objects from.
	Previous []PreviousRevisionReference `json:"previous,omitempty"`

	// Availability Probes check objects that are part of the package.
	// All probes need to succeed for a package to be considered Available.
	// Failing probes will prevent the reconciliation of objects in later phases.
	AvailabilityProbes []ObjectSetProbe `json:"availabilityProbes"`

	ObjectSetTemplatePhase `json:",inline"`
}

// ClusterObjectSetPhaseStatus defines the observed state of a ClusterObjectSetPhase.
type ClusterObjectSetPhaseStatus struct {
	// Conditions is a list of status conditions ths object is in.
	// +example=[{type: "Available", status: "True"}]
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ClusterObjectSetPhase{}, &ClusterObjectSetPhaseList{})
}
