package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/gpu"
	"github.com/samber/lo"
)

func main() {
	gpuInfo, err := ghw.GPU()
	if err != nil {
		panic(err)
	}

	fmt.Println(int32(len(lo.Filter(gpuInfo.GraphicsCards, func(t *gpu.GraphicsCard, _ int) bool {
		return t.DeviceInfo != nil
	}))))

	fmt.Println(len(gpuInfo.GraphicsCards))
}
