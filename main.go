package main

import (
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"infra/internal"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// read configuration file
		var pubKey internal.SSHPublicKey
		pulumiConf := config.New(ctx, "")                        // namespace "" refers to the project namespace "infra"
		pulumiConf.RequireObject("key", &pubKey)                 // read infra:key object
		cloudConfigPath := pulumiConf.Require("cloudConfigPath") // read infra:cloudConfigPath string

		// create Hetzner SSH Public Key
		sshKey, err := hcloud.NewSshKey(ctx, pubKey.ID, &hcloud.SshKeyArgs{
			Name:      pulumi.String(pubKey.ID),
			PublicKey: pulumi.String(pubKey.PublicKey),
		})

		// create cloud-configs
		cloudConfigs, err := internal.NewCloudConfigs(ctx, cloudConfigPath)
		if err != nil {
			return err
		}

		// use cloud-configs to initialize virtual machines
		for _, cloudConfig := range cloudConfigs {
			_, err := hcloud.NewServer(ctx, cloudConfig.ID, &hcloud.ServerArgs{
				Image:      pulumi.String("fedora-34"),
				Name:       pulumi.String(cloudConfig.ID),
				ServerType: pulumi.String("cx11"),
				SshKeys: pulumi.StringArray{
					sshKey.Name,
				},
				UserData: pulumi.String(cloudConfig.CloudConfig.Rendered),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}
