package valueObject

import "strconv"

type ContainerSpecs struct {
	CpuCores    float64 `json:"cpuCores"`
	MemoryBytes Byte    `json:"memoryBytes"`
}

func NewContainerSpecs(cpuCores float64, memoryBytes Byte) ContainerSpecs {
	return ContainerSpecs{
		CpuCores:    cpuCores,
		MemoryBytes: memoryBytes,
	}
}

func (specs ContainerSpecs) GetCpuCores() float64 {
	return specs.CpuCores
}

func (specs ContainerSpecs) GetMemoryBytes() Byte {
	return specs.MemoryBytes
}

func (specs ContainerSpecs) GetCpuCoresAsString() string {
	return strconv.FormatFloat(specs.CpuCores, 'f', 1, 64)
}

func (specs ContainerSpecs) GetMemoryAsString() string {
	return strconv.FormatInt(specs.MemoryBytes.Get(), 10)
}
