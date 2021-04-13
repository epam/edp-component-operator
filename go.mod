module github.com/epam/edp-component-operator

go 1.14

replace git.apache.org/thrift.git => github.com/apache/thrift v0.12.0

require (
	github.com/appscode/jsonpatch v0.0.0-20190108182946-7c0e3b262f30 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gophercloud/gophercloud v0.3.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
	sigs.k8s.io/testing_frameworks v0.1.1 // indirect
)
