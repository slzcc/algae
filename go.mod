module github.com/slzcc/algae

go 1.15

require (
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/gin-gonic/gin v1.6.3
	gopkg.in/igm/sockjs-go.v2 v2.1.0
	gopkg.in/yaml.v2 v2.3.0
	istio.io/client-go v1.8.1
	k8s.io/api v0.20.0
	k8s.io/apimachinery v0.20.0
	k8s.io/client-go v0.20.0
	k8s.io/metrics v0.20.0
)

replace (
	k8s.io/api => github.com/kubernetes/api v0.20.0
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.20.0
	k8s.io/client-go => github.com/kubernetes/client-go v0.20.0
)
