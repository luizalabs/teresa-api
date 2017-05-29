package k8s

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned"
)

type Config struct {
	Host               string `required:"true"`
	Username           string `required:"true"`
	Password           string `required:"true"`
	Insecure           bool   `default:"false"`
	DefaultServiceType string `default:"LoadBalancer"`
}

type Client interface {
	Namespaces() unversioned.NamespaceInterface
	LimitRanges(namespace string) unversioned.LimitRangeInterface
	Secrets(namespace string) unversioned.SecretsInterface
	HorizontalPodAutoscalers(namespace string) unversioned.HorizontalPodAutoscalerInterface
}

func New(conf *Config) (Client, error) {
	validServiceTypes := map[api.ServiceType]bool{
		api.ServiceTypeLoadBalancer: true,
		api.ServiceTypeNodePort:     true,
		api.ServiceTypeClusterIP:    true,
	}
	serviceType := api.ServiceType(conf.DefaultServiceType)
	if _, ok := validServiceTypes[serviceType]; !ok {
		return nil, fmt.Errorf(
			"invalid default service type: %s",
			conf.DefaultServiceType,
		)
	}
	k8sConf := &restclient.Config{
		Host:     conf.Host,
		Username: conf.Username,
		Password: conf.Password,
		Insecure: conf.Insecure,
	}
	return unversioned.New(k8sConf)
}
