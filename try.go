package main

import (
	"encoding/json"
	"fmt"
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	fmt.Printf("Hello \n")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s \n", container.ID[:10], container.Image)
		stats, err := cli.ContainerStats(context.Background(), container.ID, false)

		if err != nil {
			panic(err)
		}

		var v *types.StatsJSON

		dec := json.NewDecoder(stats.Body)
		dec.Decode(&v)

		fmt.Printf("No of cpu's: %d \n", len(v.CPUStats.CPUUsage.PercpuUsage))
		fmt.Println("")
		fmt.Printf("System usage: %d \n", v.CPUStats.SystemUsage)
	}
}
