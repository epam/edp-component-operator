package client

import (
	"context"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	apiV1 "github.com/epam/edp-component-operator/api/v1"
)

type EDPComponentInterface interface {
	Create(*apiV1.EDPComponent) (*apiV1.EDPComponent, error)
	Get(name string, options metaV1.GetOptions) (*apiV1.EDPComponent, error)
}

type edpComponents struct {
	client rest.Interface
	ns     string
}

func newEdpComponents(c *EDPComponentV1Client, namespace string) *edpComponents {
	return &edpComponents{
		client: c.restClient,
		ns:     namespace,
	}
}

func (e edpComponents) Create(c *apiV1.EDPComponent) (res *apiV1.EDPComponent, err error) {
	res = &apiV1.EDPComponent{}
	err = e.client.Post().
		Namespace(e.ns).
		Resource("edpcomponents").
		Body(c).
		Do(context.TODO()).
		Into(res)

	return
}

func (e edpComponents) Get(name string, options metaV1.GetOptions) (res *apiV1.EDPComponent, err error) {
	res = &apiV1.EDPComponent{}
	err = e.client.Get().
		Namespace(e.ns).
		Resource("edpcomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(res)

	return
}
