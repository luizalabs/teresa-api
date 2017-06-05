package storage

import (
	"io"
)

type storageType string

const (
	S3Type storageType = "s3"
)

type Config struct {
	Type                storageType `envconfig:"type" default:"s3"`
	AwsKey              string      `envconfig:"aws_key"`
	AwsSecret           string      `envconfig:"aws_secret"`
	AwsRegion           string      `envconfig:"aws_region"`
	AwsBucket           string      `envconfig:"aws_bucket"`
	AwsEndpoint         string      `envconfig:"aws_endpoint" default:""`
	AwsDisableSSL       bool        `envconfig:"aws_disable_ssl" default:"false"`
	AwsS3ForcePathStyle bool        `envconfig:"aws_s3_force_path_style" default:"false"`
}

type Storage interface {
	K8sSecretName() string
	AccessData() map[string][]byte
	UploadFile(path string, file io.ReadSeeker) error
	Type() string
}

func New(conf *Config) (Storage, error) {
	if conf.Type != S3Type {
		return nil, ErrInvalidStorageType
	}
	return newS3(conf), nil
}