package deploy

import (
	"fmt"
	"strconv"

	"github.com/luizalabs/teresa-api/pkg/server/app"
	st "github.com/luizalabs/teresa-api/pkg/server/storage"
)

const (
	slugBuilderImage = "luizalabs/slugbuilder:v2.4.9"
	slugRunnerImage  = "luizalabs/slugrunner:v2.2.4"
	DefaultPort      = 5000
)

type PodVolumeMountsSpec struct {
	Name      string
	MountPath string
	ReadOnly  bool
}

type PodVolumeSpec struct {
	Name       string
	SecretName string
}

type PodSpec struct {
	Name         string
	Namespace    string
	Image        string
	Env          map[string]string
	VolumeMounts []*PodVolumeMountsSpec
	Volume       []*PodVolumeSpec
}

type DeploySpec struct {
	PodSpec
	TeresaYaml
	RevisionHistoryLimit int
	Description          string
	SlugURL              string
	Args                 []string
}

func newPodSpec(name, image string, a *app.App, envVars map[string]string, fileStorage st.Storage) *PodSpec {
	ps := &PodSpec{
		Name:      name,
		Namespace: a.Name,
		Image:     image,
		VolumeMounts: []*PodVolumeMountsSpec{{
			Name:      "storage-key",
			MountPath: "/var/run/secrets/deis/objectstore/creds",
			ReadOnly:  true,
		}},
		Volume: []*PodVolumeSpec{{
			Name:       "storage-keys",
			SecretName: fileStorage.K8sSecretName(),
		}},
		Env: envVars,
	}
	for _, e := range a.EnvVars {
		ps.Env[e.Key] = e.Value
	}
	return ps
}

func newBuildSpec(a *app.App, deployId, tarBallLocation, buildDest string, fileStorage st.Storage) *PodSpec {
	return newPodSpec(
		fmt.Sprintf("build-%s", deployId),
		slugBuilderImage,
		a,
		map[string]string{
			"TAR_PATH":        tarBallLocation,
			"PUT_PATH":        buildDest,
			"BUILDER_STORAGE": fileStorage.Type(),
		},
		fileStorage,
	)
}

func newDeploySpec(a *app.App, tYaml *TeresaYaml, fileStorage st.Storage, description, slugURL, processType string) *DeploySpec {
	ps := newPodSpec(
		a.Name,
		slugRunnerImage,
		a,
		map[string]string{
			"APP":             a.Name,
			"PORT":            strconv.Itoa(DefaultPort),
			"SLUG_URL:":       slugURL,
			"BUILDER_STORAGE": fileStorage.Type(),
		},
		fileStorage,
	)
	ds := &DeploySpec{
		Args:        []string{"start", processType},
		Description: description,
		SlugURL:     slugURL,
		PodSpec:     *ps,
		TeresaYaml: TeresaYaml{
			HealthCheck:   tYaml.HealthCheck,
			RollingUpdate: tYaml.RollingUpdate,
		},
	}

	return ds
}
