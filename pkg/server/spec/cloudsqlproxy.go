package spec

import (
	"fmt"
	"path"

	"github.com/luizalabs/teresa/pkg/server/app"
	"github.com/pkg/errors"
)

const (
	cloudSQLProxyDefaultCPULimit    = "100m"
	cloudSQLProxyDefaultMemoryLimit = "256Mi"
)

type CloudSQLProxy struct {
	Instances      string
	CredentialFile string `yaml:"credentialFile"`
	Image          string `yaml:"-"`
}

func NewCloudSQLProxy(img string, t *TeresaYaml, a *app.App) (*CloudSQLProxy, error) {
	if t == nil {
		return nil, nil
	}
	v, found := t.SideCars["cloudsql-proxy"]
	if !found {
		return nil, nil
	}
	csp := new(CloudSQLProxy)
	if err := v.Unmarshal(csp); err != nil {
		return nil, errors.Wrap(err, "failed to build cloudsql proxy")
	}
	csp.Image = img
	csp.Instances = newCloudSQLProxyContainerInstanceString(csp, a)
	return csp, nil
}

func NewCloudSQLProxyContainer(csp *CloudSQLProxy, a *app.App) *Container {
	args := newCloudSQLProxyContainerArgs(csp)
	mpath := newCloudSQLProxyContainerMountPath(csp)
	env := map[string]string{}
	for _, e := range a.EnvVars {
		env[e.Key] = e.Value
	}
	return NewContainerBuilder("cloudsql-proxy", csp.Image).
		WithCommand([]string{"/cloud_sql_proxy"}).
		WithArgs(args).
		WithEnv(env).
		WithLimits(cloudSQLProxyDefaultCPULimit, cloudSQLProxyDefaultMemoryLimit).
		WithVolumeMount(AppSecretName, mpath, path.Base(mpath)).
		Build()
}

func newCloudSQLProxyContainerArgs(csp *CloudSQLProxy) []string {
	return []string{
		fmt.Sprintf("-instances=%s", csp.Instances),
		fmt.Sprintf("-credential_file=%s", newCloudSQLProxyContainerMountPath(csp)),
	}
}

func newCloudSQLProxyContainerMountPath(csp *CloudSQLProxy) string {
	return path.Join("/secrets/cloudsql", path.Base(csp.CredentialFile))
}

func newCloudSQLProxyContainerInstanceString(csp *CloudSQLProxy, a *app.App) string {
	instanceName := "%s:%s:%s=tcp:3306"
	if csp.Instances == "" {
		dbProject := ""
		dbZone := ""
		dbName := ""
		for _, envVar := range a.EnvVars {
			switch envVar.Key {
			case "DB_PROJECT":
				dbProject = envVar.Value
			case "DB_ZONE":
				dbZone = envVar.Value
			case "DB_NAME":
				dbName = envVar.Value
			}
		}
		instanceName = fmt.Sprintf("%s:%s:%s=tcp:3306", dbProject, dbZone, dbName)
	} else {
		instanceName = csp.Instances
	}
	return instanceName
}
