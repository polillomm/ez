package o11yInfra

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/speedianet/control/src/domain/entity"
	"github.com/speedianet/control/src/domain/valueObject"
	"github.com/speedianet/control/src/infra/db"
	infraEnvs "github.com/speedianet/control/src/infra/envs"
	infraHelper "github.com/speedianet/control/src/infra/helper"
)

const PublicIpTransientKey string = "PublicIp"

type O11yQueryRepo struct {
	transientDbSvc *db.TransientDatabaseService
}

func NewO11yQueryRepo(
	transientDbSvc *db.TransientDatabaseService,
) *O11yQueryRepo {
	return &O11yQueryRepo{
		transientDbSvc: transientDbSvc,
	}
}

func (repo *O11yQueryRepo) getUptime() (uint64, error) {
	sysinfo := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(sysinfo); err != nil {
		return 0, err
	}

	return uint64(sysinfo.Uptime), nil
}

func (repo *O11yQueryRepo) ReadServerPublicIpAddress() (
	ipAddress valueObject.IpAddress, err error,
) {
	cachedIpAddressStr, err := repo.transientDbSvc.Get(PublicIpTransientKey)
	if err == nil {
		return valueObject.NewIpAddress(cachedIpAddressStr)
	}

	serverPublicIpAddress, err := infraHelper.ReadServerPublicIpAddress()
	if err != nil {
		return ipAddress, errors.New("ReadServerPublicIpAddressError: " + err.Error())
	}

	err = repo.transientDbSvc.Set(PublicIpTransientKey, serverPublicIpAddress.String())
	if err != nil {
		return ipAddress, errors.New("PersistPublicIpFailed: " + err.Error())
	}

	return serverPublicIpAddress, nil
}

func (repo *O11yQueryRepo) getStorageUnitInfos() ([]valueObject.StorageUnitInfo, error) {
	storageInfos := []valueObject.StorageUnitInfo{}

	initialStats, err := disk.IOCounters()
	if err != nil {
		log.Printf("GetInitialStorageStatsFailed: %v", err)
		return storageInfos, errors.New("GetInitialStorageStatsFailed")
	}

	time.Sleep(time.Second)

	finalStats, err := disk.IOCounters()
	if err != nil {
		log.Printf("GetFinalStorageStatsFailed: %v", err)
		return storageInfos, errors.New("GetFinalStorageStatsFailed")
	}

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Printf("GetPartitionsFailed: %v", err)
		return storageInfos, errors.New("GetPartitionsFailed")
	}

	desireableFileSystems := []string{
		"xfs", "btrfs", "ext4", "ext3", "ext2", "zfs", "vfat", "ntfs",
	}
	scannedDevices := []string{}
	for _, partition := range partitions {
		if !slices.Contains(desireableFileSystems, partition.Fstype) {
			continue
		}

		if slices.Contains(scannedDevices, partition.Device) {
			continue
		}

		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		initialStats := initialStats[partition.Device]
		finalStats := finalStats[partition.Device]

		deviceName, err := valueObject.NewDeviceName(partition.Device)
		if err != nil {
			continue
		}

		mountPoint, err := valueObject.NewUnixFilePath(partition.Mountpoint)
		if err != nil {
			continue
		}

		fileSystem, err := valueObject.NewUnixFileSystem(partition.Fstype)
		if err != nil {
			continue
		}

		readBytes := finalStats.ReadBytes - initialStats.ReadBytes
		readOpsCount := finalStats.ReadCount - initialStats.ReadCount
		writeBytes := finalStats.WriteBytes - initialStats.WriteBytes
		writeOpsCount := finalStats.WriteCount - initialStats.WriteCount

		storageUnitInfo := valueObject.NewStorageUnitInfo(
			deviceName,
			mountPoint,
			fileSystem,
			valueObject.Byte(usageStat.Total),
			valueObject.Byte(usageStat.Free),
			valueObject.Byte(usageStat.Used),
			infraHelper.RoundFloat(usageStat.UsedPercent),
			usageStat.InodesTotal,
			usageStat.InodesFree,
			usageStat.InodesUsed,
			infraHelper.RoundFloat(usageStat.InodesUsedPercent),
			valueObject.Byte(readBytes),
			readOpsCount,
			valueObject.Byte(writeBytes),
			writeOpsCount,
		)

		storageInfos = append(storageInfos, storageUnitInfo)
		scannedDevices = append(scannedDevices, partition.Device)
	}

	return storageInfos, nil
}

func (repo *O11yQueryRepo) getMemLimit() (uint64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, errors.New("GetMemInfoFailed")
	}

	return memInfo.Total, nil
}

func (repo *O11yQueryRepo) getHardwareSpecs() (
	hardwareSpecs valueObject.HardwareSpecs, err error,
) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return hardwareSpecs, errors.New("GetCpuInfoFailed")
	}

	if len(cpuInfo) == 0 {
		return hardwareSpecs, errors.New("CpuInfoEmpty")
	}

	cpuModel, err := valueObject.NewCpuModelName(cpuInfo[0].ModelName)
	if err != nil {
		return hardwareSpecs, errors.New("GetCpuModelNameFailed")
	}

	cpuFrequency := infraHelper.RoundFloat(cpuInfo[0].Mhz)

	cpuCoresCount := float64(len(cpuInfo))

	memoryLimit, err := repo.getMemLimit()
	if err != nil {
		return valueObject.HardwareSpecs{}, err
	}

	return valueObject.NewHardwareSpecs(
		cpuModel,
		cpuCoresCount,
		cpuFrequency,
		valueObject.Byte(memoryLimit),
	), nil
}

func (repo *O11yQueryRepo) getCpuUsagePercent() (float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, errors.New("GetCpuUsageFailed")
	}

	return infraHelper.RoundFloat(cpuPercent[0]), nil
}

func (repo *O11yQueryRepo) getMemUsagePercent() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, errors.New("GetMemInfoFailed")
	}

	return infraHelper.RoundFloat(memInfo.UsedPercent), nil
}

func (repo *O11yQueryRepo) getNetInfos() ([]valueObject.NetInterfaceInfo, error) {
	var netInfos []valueObject.NetInterfaceInfo

	initialStats, err := net.IOCounters(true)
	if err != nil {
		log.Printf("GetInitialNetStatsFailed: %v", err)
		return netInfos, errors.New("GetInitialNetStatsFailed")
	}

	time.Sleep(time.Second)

	finalStats, err := net.IOCounters(true)
	if err != nil {
		log.Printf("GetFinalNetStatsFailed: %v", err)
		return netInfos, errors.New("GetFinalNetStatsFailed")
	}

	for i, interfaceStat := range finalStats {
		if interfaceStat.Name != initialStats[i].Name {
			continue
		}

		if interfaceStat.Name == "lo" {
			continue
		}

		deviceName, err := valueObject.NewDeviceName(interfaceStat.Name)
		if err != nil {
			continue
		}

		recvBytes := interfaceStat.BytesRecv - initialStats[i].BytesRecv
		recvPackets := interfaceStat.PacketsRecv - initialStats[i].PacketsRecv
		recvDropPackets := interfaceStat.Dropin - initialStats[i].Dropin
		recvErrs := interfaceStat.Errin - initialStats[i].Errin

		sentBytes := interfaceStat.BytesSent - initialStats[i].BytesSent
		sentPackets := interfaceStat.PacketsSent - initialStats[i].PacketsSent
		sentDropPackets := interfaceStat.Dropout - initialStats[i].Dropout
		sentErrs := interfaceStat.Errout - initialStats[i].Errout

		netInfo := valueObject.NewNetInterfaceInfo(
			deviceName,
			valueObject.Byte(recvBytes),
			recvPackets,
			recvDropPackets,
			recvErrs,
			valueObject.Byte(sentBytes),
			sentPackets,
			sentDropPackets,
			sentErrs,
		)

		netInfos = append(netInfos, netInfo)
	}

	return netInfos, nil
}

type HostResourceUsageResult struct {
	cpuUsagePercent float64
	memUsagePercent float64
	storageInfos    []valueObject.StorageUnitInfo
	netInfos        []valueObject.NetInterfaceInfo
	err             error
}

func (repo *O11yQueryRepo) getHostResourceUsage() (
	hostResourceUsage valueObject.HostResourceUsage,
	err error,
) {
	cpuChan := make(chan HostResourceUsageResult)
	memChan := make(chan HostResourceUsageResult)
	storageChan := make(chan HostResourceUsageResult)
	netChan := make(chan HostResourceUsageResult)

	go func() {
		cpuUsagePercent, err := repo.getCpuUsagePercent()
		cpuChan <- HostResourceUsageResult{cpuUsagePercent: cpuUsagePercent, err: err}
	}()

	go func() {
		memUsagePercent, err := repo.getMemUsagePercent()
		memChan <- HostResourceUsageResult{memUsagePercent: memUsagePercent, err: err}
	}()

	go func() {
		storageInfos, err := repo.getStorageUnitInfos()
		storageChan <- HostResourceUsageResult{storageInfos: storageInfos, err: err}
	}()

	go func() {
		netInfos, err := repo.getNetInfos()
		netChan <- HostResourceUsageResult{netInfos: netInfos, err: err}
	}()

	cpuResult := <-cpuChan
	if cpuResult.err != nil {
		return hostResourceUsage, errors.New("ReadCpuInfoFailed: " + cpuResult.err.Error())
	}

	memResult := <-memChan
	if memResult.err != nil {
		return hostResourceUsage, errors.New("ReadMemInfoFailed: " + memResult.err.Error())
	}

	storageResult := <-storageChan
	if storageResult.err != nil {
		return hostResourceUsage, errors.New("ReadStorageInfoFailed: " + storageResult.err.Error())
	}
	if len(storageResult.storageInfos) == 0 {
		return hostResourceUsage, errors.New("ReadStorageInfoResultEmpty")
	}

	netResult := <-netChan
	if netResult.err != nil {
		return hostResourceUsage, errors.New("ReadNetInfoFailed: " + netResult.err.Error())
	}

	cpuUsagePercentStr := strconv.FormatFloat(cpuResult.cpuUsagePercent, 'f', 0, 64)
	memUsagePercentStr := strconv.FormatFloat(memResult.memUsagePercent, 'f', 0, 64)
	userDataStorageInfo := storageResult.storageInfos[0]
	for _, storageInfo := range storageResult.storageInfos {
		if storageInfo.MountPoint.String() != infraEnvs.UserDataDirectory {
			continue
		}

		userDataStorageInfo = storageInfo
		break
	}

	return valueObject.NewHostResourceUsage(
		cpuResult.cpuUsagePercent, cpuUsagePercentStr,
		memResult.memUsagePercent, memUsagePercentStr,
		userDataStorageInfo, storageResult.storageInfos,
		netResult.netInfos,
	), nil
}

func (repo *O11yQueryRepo) ReadOverview() (overview entity.O11yOverview, err error) {
	serverHostname, err := infraHelper.ReadServerHostname()
	if err != nil {
		return overview, errors.New("ReadHostnameFailed: " + err.Error())
	}

	uptime, err := repo.getUptime()
	if err != nil {
		uptime = 0
	}

	publicIpAddress, err := repo.ReadServerPublicIpAddress()
	if err != nil {
		publicIpAddress, _ = valueObject.NewIpAddress("0.0.0.0")
	}

	hardwareSpecs, err := repo.getHardwareSpecs()
	if err != nil {
		return overview, errors.New("ReadHardwareSpecsFailed: " + err.Error())
	}

	currentResourceUsage, err := repo.getHostResourceUsage()
	if err != nil {
		return overview, errors.New("ReadHostResourceUsageFailed: " + err.Error())
	}

	return entity.NewO11yOverview(
		serverHostname, uptime, publicIpAddress, hardwareSpecs, currentResourceUsage,
	), nil
}
