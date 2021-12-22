package collectMethod

import (
	"context"
	netCommon "net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/threading"
)

func NewCollector() *collector {
	return &collector{
		DataRecord: &DataRecord{},
	}
}
func (collector *collector) collectHostInfo() {
	hostName, err := os.Hostname()
	if err != nil {
		logx.Errorf("get host name failed: %s", err)
		return
	}
	collector.DataRecord.HostName = hostName
	collector.DataRecord.OSName = runtime.GOOS
}

func (collector *collector) collectNetInfo() {
	addrs, err := netCommon.InterfaceAddrs()
	if err != nil {
		logx.Errorf("get address information failed: %s", err)
		return
	}

	var addrsList string
	for _, addr := range addrs {
		addrsList += addr.String()
	}
	collector.DataRecord.AddrsList = addrsList

	netIOInfo, err := net.IOCounters(false)
	if err != nil {
		logx.Errorf("get net io information failed: %s", err)
		return
	}

	for _, netIO := range netIOInfo {
		collector.DataRecord.NetinPackets = strconv.FormatUint(netIO.PacketsRecv, 10)
		collector.DataRecord.NetoutPackets = strconv.FormatUint(netIO.PacketsSent, 10)
		collector.DataRecord.NetinBytes = strconv.FormatUint(netIO.BytesRecv, 10)
		collector.DataRecord.NetoutBytes = strconv.FormatUint(netIO.BytesSent, 10)
	}
}

func (collector *collector) collectCPUInfo() {
	cpuPhysicalCount, err := cpu.CountsWithContext(context.Background(), false)
	if err != nil {
		logx.Errorf("get cpuPhysicalCount failed: %s", err)
		return
	}
	collector.DataRecord.CPUPhysicalCount = strconv.Itoa(cpuPhysicalCount)

	cpuLogicalCount, err := cpu.CountsWithContext(context.Background(), true)
	if err != nil {
		logx.Errorf("get cpuLogicalCount failed: %s", err)
	}
	collector.DataRecord.CPULogicalCount = strconv.Itoa(cpuLogicalCount)

	cpuStat, err := cpu.Info()
	if err != nil {
		logx.Errorf("get cpu information failed: %s", err)
		return
	}

	if len(cpuStat) <= 0 {
		return
	}

	collector.DataRecord.CPUName = cpuStat[0].ModelName

	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		logx.Errorf("get cpu usage failed: %s", err)
		return
	}

	if len(cpuUsage) <= 0 {
		return
	}

	cpuString := strconv.FormatFloat(cpuUsage[0], 'E', -1, 64)
	collector.DataRecord.CPUUsage = cpuString
}

func (collector *collector) collectMemInfo() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		logx.Errorf("get memory information failed: %s", err)
		return
	}
	collector.DataRecord.MemTotal = strconv.FormatUint(memInfo.Total, 10)
	collector.DataRecord.MemUsed = strconv.FormatUint(memInfo.Used, 10)

	swapInfo, err := mem.SwapMemory()
	if err != nil {
		logx.Errorf("get swap memory failed: %s", err)
		return
	}
	collector.DataRecord.SwapTotal = strconv.FormatUint(swapInfo.Total, 10)
}

func (collector *collector) collectDiskInfo() {
	cmd := []string{"/bin/df", "-B1"}
	cmdObj := exec.CommandContext(context.Background(), cmd[0], cmd[1:]...)
	output, err := cmdObj.Output()
	if err != nil {
		logx.Errorf("get disk information failed: %s", err)
		return
	}

	regInfo, err := regexp.Compile(`(/dev/[a-z]{3}\d?)\s+(\S+)\s+(\S+)`)
	if err != nil {
		logx.Errorf("create reg failed: %s", err)
		return
	}

	var diskTotal, diskUsed float64
	disksInfo := regInfo.FindAllStringSubmatch(string(output), -1)
	for _, diskInfo := range disksInfo {
		diskEachTotal, _ := strconv.ParseFloat(diskInfo[2], 64)
		diskTotal += diskEachTotal
		diskEachUsed, _ := strconv.ParseFloat(diskInfo[3], 64)
		diskUsed += diskEachUsed
	}

	collector.DataRecord.DiskTotal = strconv.FormatFloat(diskTotal, 'E', -1, 64)
	collector.DataRecord.DiskUsed = strconv.FormatFloat(diskUsed, 'E', -1, 64)
}

func (collector *collector) CollectData() {
	collector.DataRecord.Time =  strconv.FormatInt(time.Now().Unix(),10)

	pool := threading.NewTaskRunner(coroutinesNum)
	var wg sync.WaitGroup
	wg.Add(coroutinesNum)

	pool.Schedule(func() {
		defer wg.Done()
		collector.collectHostInfo()
	})

	pool.Schedule(func() {
		defer wg.Done()
		collector.collectNetInfo()
	})

	pool.Schedule(func() {
		defer wg.Done()
		collector.collectCPUInfo()
	})

	pool.Schedule(func() {
		defer wg.Done()
		collector.collectMemInfo()
	})

	pool.Schedule(func() {
		defer wg.Done()
		collector.collectDiskInfo()
	})

	wg.Wait()
}
