package dnsserver

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/sparrowhe/dockerdns/common"
)

func PickIpAddr(name string) (string, error) {
	name = strings.TrimRight(name, ".clipd")
	common.Logger.Info("Lookup for: ", name)
	ctx := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	dockerClient.NegotiateAPIVersion(ctx)
	defer dockerClient.Close()
	containerList, err := dockerClient.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		common.Logger.Warn(err)
		return "", err
	}
	for _, ctr := range containerList {
		if len(ctr.Names) > 0 && strings.TrimLeft(ctr.Names[0], "/") == name {
			ipAddrs := ctr.NetworkSettings.Networks
			for _, networkSettings := range ipAddrs {
				return networkSettings.IPAddress, nil
			}
		}
	}
	return "", nil
}
