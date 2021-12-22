package collectMethod

type collector struct {
	DataRecord *DataRecord
}

type DataRecord struct {
	Time             string  `json:"time"`
	HostName         string `json:"hostName"`
	OSName           string `json:"osName"`
	AddrsList        string `json:"addrsList"`
	NetinPackets     string `json:"netinPackets"`
	NetoutPackets    string `json:"netoutPackets"`
	NetinBytes       string `json:"netinBytes"`
	NetoutBytes      string `json:"netoutBytes"`
	CPUPhysicalCount string `json:"cpuPhysicalCount"`
	CPULogicalCount  string `json:"cpuLogicalCount"`
	CPUName          string `json:"cpuName"`
	CPUUsage         string `json:"cpuUsage"`
	MemTotal         string `json:"memTotal"`
	MemUsed          string `json:"memUsed"`
	SwapTotal        string `json:"swapTotal"`
	DiskTotal        string `json:"diskTotal"`
	DiskUsed         string `json:"diskUsed"`
}

var coroutinesNum = 5
