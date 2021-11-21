package internal

import (
	"github.com/pulumi/pulumi-cloudinit/sdk/go/cloudinit"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// cloudConfigContentType cannot be a constant, because we cannot use pointers to constants in Go
var cloudConfigContentType = "text/cloud-config"

// CloudConfig extends the pulumi cloud-config with an ID
type CloudConfig struct {
	ID          string
	CloudConfig *cloudinit.LookupConfigResult
}

// NewCloudConfigs reads all cloud-config files in a given path and returns
// a slice of CloudConfig
func NewCloudConfigs(ctx *pulumi.Context, path string) ([]CloudConfig, error) {
	var cloudConfigs []CloudConfig

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		config, err := ioutil.ReadFile(filepath.Join(path, f.Name()))
		if err != nil {
			return nil, err
		}
		cloudConfig, err := cloudinit.LookupConfig(ctx, &cloudinit.LookupConfigArgs{
			Base64Encode: BoolPtr(false),
			Gzip:         BoolPtr(false),
			Parts: []cloudinit.GetConfigPart{
				{
					Content:     string(config),
					ContentType: &cloudConfigContentType,
					Filename:    StringPtr(f.Name()),
				},
			},
		})
		if err != nil {
			return nil, err
		}
		cloudConfigs = append(cloudConfigs, CloudConfig{
			ID:          strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())),
			CloudConfig: cloudConfig,
		})
	}
	return cloudConfigs, nil
}
