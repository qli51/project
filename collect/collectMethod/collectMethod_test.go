package collectMethod

import (
	"testing"
	"os"
	"errors"
	netCommon "net"
	"context"
	"time"
	"regexp"

	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func TestNewCollector(t *testing.T) {
	Convey("创建收集器", t, func() {
		res := NewCollector()
		So(res, ShouldNotBeNil)
	})
}

func TestCollectHostInfo(t *testing.T) {
	Convey("成功获取主机信息", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(os.Hostname, func() (string, error) {
			return "test", nil
		})
		defer patcheOne.Reset()

		collector.collectHostInfo()
		So(collector.DataRecord.HostName, ShouldEqual, "test")
	})
	Convey("获取失败情况", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(os.Hostname, func() (string, error) {
			return "wrong", errors.New("get hostname failed")
		})
		defer patcheOne.Reset()

		collector.collectHostInfo()
		So(collector.DataRecord.HostName, ShouldEqual, "")
	})
}

func TestCollectNetInfo(t *testing.T) {
	Convey("正确获取网络信息", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(netCommon.InterfaceAddrs, func() ([]netCommon.Addr, error) {
			return []netCommon.Addr{}, nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(net.IOCounters, func(bool) ([]net.IOCountersStat, error) {
			return []net.IOCountersStat{net.IOCountersStat{PacketsRecv: 1}}, nil
		})
		defer patcheTwo.Reset()

		collector.collectNetInfo()
		So(collector.DataRecord.NetinPackets, ShouldEqual, "1")
	})

	Convey("获取ip信息失败", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(netCommon.InterfaceAddrs, func() ([]netCommon.Addr, error) {
			return nil, errors.New("get ip failed")
		})
		defer patcheOne.Reset()

		collector.collectNetInfo()
		So(collector.DataRecord.AddrsList, ShouldEqual, "")
	})

	Convey("获取流速信息失败", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(netCommon.InterfaceAddrs, func() ([]netCommon.Addr, error) {
			return nil, errors.New("get ip failed")
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(net.IOCounters, func(bool) ([]net.IOCountersStat, error) {
			return nil, errors.New("get net io failed")
		})
		defer patcheTwo.Reset()

		collector.collectNetInfo()
		So(collector.DataRecord.NetinPackets, ShouldEqual, "")
	})
}

func TestCollectCPUInfo(t *testing.T) {
	Convey("正确获取cpu信息", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(cpu.CountsWithContext, func(context.Context, bool) (int, error) {
			return 1, nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(cpu.Info, func() ([]cpu.InfoStat, error) {
			return []cpu.InfoStat{cpu.InfoStat{ModelName:"test"}}, nil
		})
		defer patcheTwo.Reset()

		patcheThree := ApplyFunc(cpu.Percent, func(time.Duration, bool) ([]float64, error) {
			return []float64{0.5}, nil
		})
		defer patcheThree.Reset()

		collector.collectCPUInfo()
		So(collector.DataRecord.CPUPhysicalCount, ShouldEqual, "1")
		So(collector.DataRecord.CPUName, ShouldEqual, "test")
		So(collector.DataRecord.CPUUsage, ShouldEqual, "5E-01")
	})

	Convey("获取cpu核数错误", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(cpu.CountsWithContext, func(context.Context, bool) (int, error) {
			return 0, errors.New("get cpu core failed")
		})
		defer patcheOne.Reset()

		collector.collectCPUInfo()
		So(collector.DataRecord.CPUPhysicalCount, ShouldEqual, "")
	})

	Convey("获取cpu名称错误", t, func() {
		collector := NewCollector()
		patcheTwo := ApplyFunc(cpu.Info, func() ([]cpu.InfoStat, error) {
			return nil, errors.New("get cpu module name failed")
		})
		defer patcheTwo.Reset()

		collector.collectCPUInfo()
		So(collector.DataRecord.CPUName, ShouldEqual, "")
	})

	Convey("获取cpu使用率错误", t, func() {
		collector := NewCollector()
		patcheTwo := ApplyFunc(cpu.Info, func() ([]cpu.InfoStat, error) {
			return []cpu.InfoStat{cpu.InfoStat{ModelName:"test"}}, nil
		})
		defer patcheTwo.Reset()
		
		patcheThree := ApplyFunc(cpu.Percent, func(time.Duration, bool) ([]float64, error) {
			return nil, nil
		})
		defer patcheThree.Reset()

		collector.collectCPUInfo()
		So(collector.DataRecord.CPUUsage, ShouldEqual, "")
	})
}

func TestCollectMemInfo(t *testing.T) {
	Convey("正确获取内存信息", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(mem.VirtualMemory, func() (*mem.VirtualMemoryStat, error) {
			return &mem.VirtualMemoryStat{Total : 2, Used : 1}, nil
		})
		defer patcheOne.Reset()

		patcheTwo := ApplyFunc(mem.SwapMemory, func() (*mem.SwapMemoryStat, error) {
			return &mem.SwapMemoryStat{Total : 1}, nil
		})
		defer patcheTwo.Reset()


		collector.collectMemInfo()
		So(collector.DataRecord.MemTotal, ShouldEqual, "2")
		So(collector.DataRecord.MemUsed, ShouldEqual, "1")
		So(collector.DataRecord.SwapTotal, ShouldEqual, "1")
	})

	Convey("获取内存失败", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(mem.VirtualMemory, func() (*mem.VirtualMemoryStat, error) {
			return &mem.VirtualMemoryStat{}, errors.New("get memory failed")
		})
		defer patcheOne.Reset()

		collector.collectMemInfo()
		So(collector.DataRecord.MemTotal, ShouldEqual, "")
	})

	Convey("获取swap失败", t, func() {
		collector := NewCollector()
		patcheTwo := ApplyFunc(mem.SwapMemory, func() (*mem.SwapMemoryStat, error) {
			return &mem.SwapMemoryStat{}, errors.New("get swap failed")
		})
		defer patcheTwo.Reset()

		collector.collectMemInfo()
		So(collector.DataRecord.SwapTotal, ShouldEqual, "")
	})
}

func TestCollectDiskInfo(t *testing.T) {
	Convey("正确获取磁盘容量信息", t, func() {
		collector := NewCollector()
		collector.collectDiskInfo()
		So(collector.DataRecord.DiskTotal, ShouldNotBeNil)
	})
	Convey("生成匹配器失败", t, func() {
		collector := NewCollector()
		patcheOne := ApplyFunc(regexp.Compile, func(string) (*regexp.Regexp, error) {
			return &regexp.Regexp{}, errors.New("generate regexp failed")
		})
		defer patcheOne.Reset()
		collector.collectDiskInfo()
		So(collector.DataRecord.DiskTotal, ShouldEqual, "")
	})
}

func TestCollectData(t *testing.T) {
	Convey("开启协程执行任务", t, func() {
		collector := NewCollector()
		collector.CollectData()
	})
}