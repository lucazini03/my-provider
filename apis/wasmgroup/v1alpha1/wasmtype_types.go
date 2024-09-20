package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// EDIT: Aggiungi qui i dettagli del Custom Resource per il modulo Wasm

// WASMTypeSpec definisce i parametri desiderati per il modulo Wasm.
type WASMTypeSpec struct {
    xpv1.ResourceSpec `json:",inline"`
    
    // Modulo Wasm container image
    Image string `json:"image"`

    // Nome del runtime class per eseguire il modulo Wasm
    RuntimeClassName string `json:"runtimeClassName"`
}

// WASMTypeStatus rappresenta lo stato osservato del modulo Wasm.
type WASMTypeStatus struct {
    xpv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// WASMType è il CR per il modulo Wasm.
type WASMType struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   WASMTypeSpec   `json:"spec"`
    Status WASMTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WASMTypeList è una lista di oggetti WASMType.
type WASMTypeList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []WASMType `json:"items"`
}

