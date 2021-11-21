package main

import (
	"fmt"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"infra/internal"
	"os/user"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// retrieve current user
		user, err := user.Current()
		if err != nil {
			return err
		}

		key, err := internal.ReadSSHPublicKey(fmt.Sprintf("%s/.ssh/id_rsa.pub", user.HomeDir))
		if err != nil {
			return err
		}

		fmt.Println(key.Comment)
		fmt.Println(string(key.Bytes))
		sshKey, err := hcloud.NewSshKey(ctx, key.Comment, &hcloud.SshKeyArgs{
			Name:      pulumi.String(key.Comment),
			PublicKey: pulumi.String(key.Bytes),
		})

		cloudConfigs, err := internal.NewCloudConfigs(ctx, "assets/cloud-config")
		if err != nil {
			return err
		}

		for _, config := range cloudConfigs {
			_, err := hcloud.NewServer(ctx, config.ID, &hcloud.ServerArgs{
				Image:      pulumi.String("fedora-34"),
				Name:       pulumi.String(config.ID),
				ServerType: pulumi.String("cx11"),
				SshKeys: pulumi.StringArray{
					sshKey.Name,
				},
				UserData: pulumi.String(config.CloudConfig.Rendered),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}
