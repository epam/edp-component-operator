package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"

	apiV1 "github.com/epam/edp-component-operator/pkg/apis/v1/v1"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "v1.edp.epam.com", Version: "v1"}

type EDPComponentV1Client struct {
	restClient *rest.RESTClient
}

func (c *EDPComponentV1Client) EDPComponents(namespace string) EDPComponentInterface {
	return newEdpComponents(c, namespace)
}

func NewForConfig(c *rest.Config) (*EDPComponentV1Client, error) {
	if err := setConfigDefault(c); err != nil {
		return nil, err
	}
	rc, err := rest.RESTClientFor(c)
	if err != nil {
		return nil, err
	}
	return &EDPComponentV1Client{restClient: rc}, nil
}

func setConfigDefault(cfg *rest.Config) error {
	scheme := runtime.NewScheme()
	sb := runtime.NewSchemeBuilder(addKnownTypes)
	if err := sb.AddToScheme(scheme); err != nil {
		return err
	}
	cfg.GroupVersion = &SchemeGroupVersion
	cfg.APIPath = "/apis"
	cfg.ContentType = runtime.ContentTypeJSON
	cfg.NegotiatedSerializer = serializer.NewCodecFactory(scheme)
	return nil
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&apiV1.EDPComponent{},
		&apiV1.EDPComponentList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
