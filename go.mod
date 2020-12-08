module github.com/slzcc/algae

go 1.15

require (
	github.com/emicklei/go-restful v2.11.1+incompatible
	github.com/gin-gonic/gin v1.6.3
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/igm/sockjs-go.v2 v2.1.0
	gopkg.in/yaml.v2 v2.3.0
	istio.io/client-go v1.8.0
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	k8s.io/metrics v0.19.4
)

replace (
	k8s.io/api => github.com/kubernetes/api v0.19.4
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.19.4
	k8s.io/client-go => github.com/kubernetes/client-go v0.19.4
)
