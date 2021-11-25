module github.com/epam/edp-component-operator

go 1.14

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	k8s.io/api => k8s.io/api v0.20.7-rc.0
)

require (
	github.com/epam/edp-common v0.0.0-20211124100535-e54dcdf42879
	k8s.io/apimachinery v0.21.0-rc.0
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
)
