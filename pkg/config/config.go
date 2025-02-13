package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rancher/opni-monitoring/pkg/config/meta"
	"github.com/rancher/opni-monitoring/pkg/config/v1beta1"
	"github.com/rancher/opni-monitoring/pkg/logger"
	"go.uber.org/zap"
	"sigs.k8s.io/yaml"
)

var (
	ErrConfigNotFound        = errors.New("config not found")
	ErrUnsupportedApiVersion = errors.New("unsupported api version")
)

type Unmarshaler interface {
	Unmarshal(into interface{}) error
}

var configLog = logger.New().Named("config")

type GatewayConfig = v1beta1.GatewayConfig

func LoadObjectsFromFile(path string) (meta.ObjectList, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	objects := []meta.Object{}
	documents := bytes.Split(data, []byte("\n---\n"))
	for i, document := range documents {
		lg := configLog.With(
			"path", path,
			"documentIndex", i,
		)
		if len(strings.TrimSpace(string(document))) == 0 {
			continue
		}
		typeMeta := meta.TypeMeta{}
		if err := yaml.Unmarshal(document, &typeMeta); err != nil {
			lg.With(
				zap.Error(err),
			).Error("object has missing or invalid TypeMeta")
			continue
		}
		if typeMeta.APIVersion == "" || typeMeta.Kind == "" {
			lg.Error("object has missing or invalid TypeMeta")
			continue
		}
		object, err := decodeObject(typeMeta, document)
		if err != nil {
			lg.With(
				zap.Error(err),
			).Error("failed to decode object")
			continue
		}
		objects = append(objects, object)
	}
	return objects, nil
}

func decodeObject(typeMeta meta.TypeMeta, document []byte) (meta.Object, error) {
	switch typeMeta.APIVersion {
	case v1beta1.APIVersion:
		return v1beta1.DecodeObject(typeMeta.Kind, document)
	default:
		return nil, ErrUnsupportedApiVersion
	}
}

func FindConfig() (string, error) {
	pathsToSearch := []string{
		".",
		"/etc/opni-monitoring",
	}
	filenamesToSearch := []string{
		"gateway.yaml",
		"gateway.yml",
		"gateway.json",
		"agent.yaml",
		"agent.yml",
		"agent.json",
		"config.yaml",
		"config.yml",
		"config.json",
	}

	for _, path := range pathsToSearch {
		for _, filename := range filenamesToSearch {
			p, err := filepath.Abs(filepath.Join(path, filename))
			if err != nil {
				return "", err
			}
			if f, err := os.Open(p); err == nil {
				f.Close()
				return p, nil
			}
		}
	}

	return "", ErrConfigNotFound
}
