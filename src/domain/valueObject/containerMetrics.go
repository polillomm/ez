package valueObject

type ContainerMetrics struct {
	ContainerId        ContainerId `json:"-"`
	CurrentCpuPercent  float64     `json:"currentCpuPercent"`
	AverageCpuPercent  float64     `json:"avgCpuPercent"`
	MemoryBytes        Byte        `json:"memoryBytes"`
	MemoryPercent      float64     `json:"memoryPercent"`
	StorageInputBytes  Byte        `json:"storageInputBytes"`
	StorageOutputBytes Byte        `json:"storageOutputBytes"`
	StorageSpaceBytes  Byte        `json:"storageSpaceBytes"`
	StorageInodesCount uint64      `json:"storageInodesCount"`
	NetInputBytes      Byte        `json:"netInputBytes"`
	NetOutputBytes     Byte        `json:"netOutputBytes"`
}

func NewContainerMetrics(
	containerId ContainerId,
	currentCpuPercent, averageCpuPercent float64,
	memoryBytes Byte,
	memoryPercent float64,
	storageInputBytes, storageOutputBytes, storageSpaceBytes Byte,
	storageInodesCount uint64,
	netInputBytes, netOutputBytes Byte,
) ContainerMetrics {
	return ContainerMetrics{
		ContainerId:        containerId,
		CurrentCpuPercent:  currentCpuPercent,
		AverageCpuPercent:  averageCpuPercent,
		MemoryBytes:        memoryBytes,
		MemoryPercent:      memoryPercent,
		StorageInputBytes:  storageInputBytes,
		StorageOutputBytes: storageOutputBytes,
		StorageSpaceBytes:  storageSpaceBytes,
		StorageInodesCount: storageInodesCount,
		NetInputBytes:      netInputBytes,
		NetOutputBytes:     netOutputBytes,
	}
}

func NewBlankContainerMetrics(containerId ContainerId) ContainerMetrics {
	return ContainerMetrics{
		ContainerId:        containerId,
		CurrentCpuPercent:  0,
		AverageCpuPercent:  0,
		MemoryBytes:        0,
		MemoryPercent:      0,
		StorageInputBytes:  0,
		StorageOutputBytes: 0,
		StorageSpaceBytes:  0,
		StorageInodesCount: 0,
		NetInputBytes:      0,
		NetOutputBytes:     0,
	}
}
