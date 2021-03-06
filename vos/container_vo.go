package vos

import (
	"strings"
	"time"
)

type PortVo struct {
	PrivatePort int64
	PublicPort  int64
	Type        string
	IP          string
}

type ContainersVo struct {
	ID      string `json:"Id"`
	Image   string
	Command string
	Created int64
	State   string
	Status  string
	Ports   []PortVo
	Names   []string
}

type ContainerNetworkVo struct {
	Aliases             []string
	MacAddress          string
	GlobalIPv6PrefixLen int
	GlobalIPv6Address   string
	IPv6Gateway         string
	IPPrefixLen         int
	IPAddress           string
	Gateway             string
	EndpointID          string
	NetworkID           string
}

type PortBindingVo struct {
	HostIP   string `json:"HostIp,omitempty"`
	HostPort string `json:"HostPort,omitempty"`
}

type Port string

func (p Port) Port() string {
	return strings.Split(string(p), "/")[0]
}

func (p Port) Protocol() string {
	parts := strings.Split(string(p), "/")
	if len(parts) == 1 {
		return "tcp"
	}
	return parts[1]
}

type NetworkSettingsVo struct {
	Networks               map[string]ContainerNetworkVo
	IPAddress              string
	IPPrefixLen            int
	MacAddress             string
	Gateway                string
	Bridge                 string
	Ports                  map[Port][]PortBindingVo
	NetworkID              string
	EndpointID             string
	SandboxKey             string
	GlobalIPv6Address      string
	GlobalIPv6PrefixLen    int
	IPv6Gateway            string
	LinkLocalIPv6Address   string
	LinkLocalIPv6PrefixLen int
	SecondaryIPAddresses   []string
	SecondaryIPv6Addresses []string
}

type ContainerDetailVo struct {
	ID              string `json:"Id"`
	Created         time.Time
	Path            string
	Args            []string
	Image           string
	NetworkSettings *NetworkSettingsVo
	Name            string
	Driver          string
	Volumes         map[string]string
}

type ContainerProcesses struct {
	Titles    []string
	Processes [][]string
}

type ContainerStats struct {
	Read      time.Time `json:"read,omitempty" yaml:"read,omitempty" toml:"read,omitempty"`
	PreRead   time.Time `json:"preread,omitempty" yaml:"preread,omitempty" toml:"preread,omitempty"`
	NumProcs  uint32    `json:"num_procs" yaml:"num_procs" toml:"num_procs"`
	PidsStats struct {
		Current uint64 `json:"current,omitempty" yaml:"current,omitempty"`
	} `json:"pids_stats,omitempty" yaml:"pids_stats,omitempty" toml:"pids_stats,omitempty"`
	Network     NetworkStats            `json:"network,omitempty" yaml:"network,omitempty" toml:"network,omitempty"`
	Networks    map[string]NetworkStats `json:"networks,omitempty" yaml:"networks,omitempty" toml:"networks,omitempty"`
	MemoryStats struct {
		Stats struct {
			TotalPgmafault          uint64 `json:"total_pgmafault,omitempty" yaml:"total_pgmafault,omitempty" toml:"total_pgmafault,omitempty"`
			Cache                   uint64 `json:"cache,omitempty" yaml:"cache,omitempty" toml:"cache,omitempty"`
			MappedFile              uint64 `json:"mapped_file,omitempty" yaml:"mapped_file,omitempty" toml:"mapped_file,omitempty"`
			TotalInactiveFile       uint64 `json:"total_inactive_file,omitempty" yaml:"total_inactive_file,omitempty" toml:"total_inactive_file,omitempty"`
			Pgpgout                 uint64 `json:"pgpgout,omitempty" yaml:"pgpgout,omitempty" toml:"pgpgout,omitempty"`
			Rss                     uint64 `json:"rss,omitempty" yaml:"rss,omitempty" toml:"rss,omitempty"`
			TotalMappedFile         uint64 `json:"total_mapped_file,omitempty" yaml:"total_mapped_file,omitempty" toml:"total_mapped_file,omitempty"`
			Writeback               uint64 `json:"writeback,omitempty" yaml:"writeback,omitempty" toml:"writeback,omitempty"`
			Unevictable             uint64 `json:"unevictable,omitempty" yaml:"unevictable,omitempty" toml:"unevictable,omitempty"`
			Pgpgin                  uint64 `json:"pgpgin,omitempty" yaml:"pgpgin,omitempty" toml:"pgpgin,omitempty"`
			TotalUnevictable        uint64 `json:"total_unevictable,omitempty" yaml:"total_unevictable,omitempty" toml:"total_unevictable,omitempty"`
			Pgmajfault              uint64 `json:"pgmajfault,omitempty" yaml:"pgmajfault,omitempty" toml:"pgmajfault,omitempty"`
			TotalRss                uint64 `json:"total_rss,omitempty" yaml:"total_rss,omitempty" toml:"total_rss,omitempty"`
			TotalRssHuge            uint64 `json:"total_rss_huge,omitempty" yaml:"total_rss_huge,omitempty" toml:"total_rss_huge,omitempty"`
			TotalWriteback          uint64 `json:"total_writeback,omitempty" yaml:"total_writeback,omitempty" toml:"total_writeback,omitempty"`
			TotalInactiveAnon       uint64 `json:"total_inactive_anon,omitempty" yaml:"total_inactive_anon,omitempty" toml:"total_inactive_anon,omitempty"`
			RssHuge                 uint64 `json:"rss_huge,omitempty" yaml:"rss_huge,omitempty" toml:"rss_huge,omitempty"`
			HierarchicalMemoryLimit uint64 `json:"hierarchical_memory_limit,omitempty" yaml:"hierarchical_memory_limit,omitempty" toml:"hierarchical_memory_limit,omitempty"`
			TotalPgfault            uint64 `json:"total_pgfault,omitempty" yaml:"total_pgfault,omitempty" toml:"total_pgfault,omitempty"`
			TotalActiveFile         uint64 `json:"total_active_file,omitempty" yaml:"total_active_file,omitempty" toml:"total_active_file,omitempty"`
			ActiveAnon              uint64 `json:"active_anon,omitempty" yaml:"active_anon,omitempty" toml:"active_anon,omitempty"`
			TotalActiveAnon         uint64 `json:"total_active_anon,omitempty" yaml:"total_active_anon,omitempty" toml:"total_active_anon,omitempty"`
			TotalPgpgout            uint64 `json:"total_pgpgout,omitempty" yaml:"total_pgpgout,omitempty" toml:"total_pgpgout,omitempty"`
			TotalCache              uint64 `json:"total_cache,omitempty" yaml:"total_cache,omitempty" toml:"total_cache,omitempty"`
			InactiveAnon            uint64 `json:"inactive_anon,omitempty" yaml:"inactive_anon,omitempty" toml:"inactive_anon,omitempty"`
			ActiveFile              uint64 `json:"active_file,omitempty" yaml:"active_file,omitempty" toml:"active_file,omitempty"`
			Pgfault                 uint64 `json:"pgfault,omitempty" yaml:"pgfault,omitempty" toml:"pgfault,omitempty"`
			InactiveFile            uint64 `json:"inactive_file,omitempty" yaml:"inactive_file,omitempty" toml:"inactive_file,omitempty"`
			TotalPgpgin             uint64 `json:"total_pgpgin,omitempty" yaml:"total_pgpgin,omitempty" toml:"total_pgpgin,omitempty"`
			HierarchicalMemswLimit  uint64 `json:"hierarchical_memsw_limit,omitempty" yaml:"hierarchical_memsw_limit,omitempty" toml:"hierarchical_memsw_limit,omitempty"`
			Swap                    uint64 `json:"swap,omitempty" yaml:"swap,omitempty" toml:"swap,omitempty"`
		} `json:"stats,omitempty" yaml:"stats,omitempty" toml:"stats,omitempty"`
		MaxUsage          uint64 `json:"max_usage,omitempty" yaml:"max_usage,omitempty" toml:"max_usage,omitempty"`
		Usage             uint64 `json:"usage,omitempty" yaml:"usage,omitempty" toml:"usage,omitempty"`
		Failcnt           uint64 `json:"failcnt,omitempty" yaml:"failcnt,omitempty" toml:"failcnt,omitempty"`
		Limit             uint64 `json:"limit,omitempty" yaml:"limit,omitempty" toml:"limit,omitempty"`
		Commit            uint64 `json:"commitbytes,omitempty" yaml:"commitbytes,omitempty" toml:"privateworkingset,omitempty"`
		CommitPeak        uint64 `json:"commitpeakbytes,omitempty" yaml:"commitpeakbytes,omitempty" toml:"commitpeakbytes,omitempty"`
		PrivateWorkingSet uint64 `json:"privateworkingset,omitempty" yaml:"privateworkingset,omitempty" toml:"privateworkingset,omitempty"`
	} `json:"memory_stats,omitempty" yaml:"memory_stats,omitempty" toml:"memory_stats,omitempty"`
	BlkioStats struct {
		IOServiceBytesRecursive []BlkioStatsEntry `json:"io_service_bytes_recursive,omitempty" yaml:"io_service_bytes_recursive,omitempty" toml:"io_service_bytes_recursive,omitempty"`
		IOServicedRecursive     []BlkioStatsEntry `json:"io_serviced_recursive,omitempty" yaml:"io_serviced_recursive,omitempty" toml:"io_serviced_recursive,omitempty"`
		IOQueueRecursive        []BlkioStatsEntry `json:"io_queue_recursive,omitempty" yaml:"io_queue_recursive,omitempty" toml:"io_queue_recursive,omitempty"`
		IOServiceTimeRecursive  []BlkioStatsEntry `json:"io_service_time_recursive,omitempty" yaml:"io_service_time_recursive,omitempty" toml:"io_service_time_recursive,omitempty"`
		IOWaitTimeRecursive     []BlkioStatsEntry `json:"io_wait_time_recursive,omitempty" yaml:"io_wait_time_recursive,omitempty" toml:"io_wait_time_recursive,omitempty"`
		IOMergedRecursive       []BlkioStatsEntry `json:"io_merged_recursive,omitempty" yaml:"io_merged_recursive,omitempty" toml:"io_merged_recursive,omitempty"`
		IOTimeRecursive         []BlkioStatsEntry `json:"io_time_recursive,omitempty" yaml:"io_time_recursive,omitempty" toml:"io_time_recursive,omitempty"`
		SectorsRecursive        []BlkioStatsEntry `json:"sectors_recursive,omitempty" yaml:"sectors_recursive,omitempty" toml:"sectors_recursive,omitempty"`
	} `json:"blkio_stats,omitempty" yaml:"blkio_stats,omitempty" toml:"blkio_stats,omitempty"`
	CPUStats     CPUStats `json:"cpu_stats,omitempty" yaml:"cpu_stats,omitempty" toml:"cpu_stats,omitempty"`
	PreCPUStats  CPUStats `json:"precpu_stats,omitempty"`
	StorageStats struct {
		ReadCountNormalized  uint64 `json:"read_count_normalized,omitempty" yaml:"read_count_normalized,omitempty" toml:"read_count_normalized,omitempty"`
		ReadSizeBytes        uint64 `json:"read_size_bytes,omitempty" yaml:"read_size_bytes,omitempty" toml:"read_size_bytes,omitempty"`
		WriteCountNormalized uint64 `json:"write_count_normalized,omitempty" yaml:"write_count_normalized,omitempty" toml:"write_count_normalized,omitempty"`
		WriteSizeBytes       uint64 `json:"write_size_bytes,omitempty" yaml:"write_size_bytes,omitempty" toml:"write_size_bytes,omitempty"`
	} `json:"storage_stats,omitempty" yaml:"storage_stats,omitempty" toml:"storage_stats,omitempty"`
}

type NetworkStats struct {
	RxDropped uint64 `json:"rx_dropped,omitempty" yaml:"rx_dropped,omitempty" toml:"rx_dropped,omitempty"`
	RxBytes   uint64 `json:"rx_bytes,omitempty" yaml:"rx_bytes,omitempty" toml:"rx_bytes,omitempty"`
	RxErrors  uint64 `json:"rx_errors,omitempty" yaml:"rx_errors,omitempty" toml:"rx_errors,omitempty"`
	TxPackets uint64 `json:"tx_packets,omitempty" yaml:"tx_packets,omitempty" toml:"tx_packets,omitempty"`
	TxDropped uint64 `json:"tx_dropped,omitempty" yaml:"tx_dropped,omitempty" toml:"tx_dropped,omitempty"`
	RxPackets uint64 `json:"rx_packets,omitempty" yaml:"rx_packets,omitempty" toml:"rx_packets,omitempty"`
	TxErrors  uint64 `json:"tx_errors,omitempty" yaml:"tx_errors,omitempty" toml:"tx_errors,omitempty"`
	TxBytes   uint64 `json:"tx_bytes,omitempty" yaml:"tx_bytes,omitempty" toml:"tx_bytes,omitempty"`
}

type CPUStats struct {
	CPUUsage struct {
		PercpuUsage       []uint64 `json:"percpu_usage,omitempty" yaml:"percpu_usage,omitempty" toml:"percpu_usage,omitempty"`
		UsageInUsermode   uint64   `json:"usage_in_usermode,omitempty" yaml:"usage_in_usermode,omitempty" toml:"usage_in_usermode,omitempty"`
		TotalUsage        uint64   `json:"total_usage,omitempty" yaml:"total_usage,omitempty" toml:"total_usage,omitempty"`
		UsageInKernelmode uint64   `json:"usage_in_kernelmode,omitempty" yaml:"usage_in_kernelmode,omitempty" toml:"usage_in_kernelmode,omitempty"`
	} `json:"cpu_usage,omitempty" yaml:"cpu_usage,omitempty" toml:"cpu_usage,omitempty"`
	SystemCPUUsage uint64 `json:"system_cpu_usage,omitempty" yaml:"system_cpu_usage,omitempty" toml:"system_cpu_usage,omitempty"`
	ThrottlingData struct {
		Periods          uint64 `json:"periods,omitempty"`
		ThrottledPeriods uint64 `json:"throttled_periods,omitempty"`
		ThrottledTime    uint64 `json:"throttled_time,omitempty"`
	} `json:"throttling_data,omitempty" yaml:"throttling_data,omitempty" toml:"throttling_data,omitempty"`
}

type BlkioStatsEntry struct {
	Major uint64 `json:"major,omitempty" yaml:"major,omitempty" toml:"major,omitempty"`
	Minor uint64 `json:"minor,omitempty" yaml:"minor,omitempty" toml:"minor,omitempty"`
	Op    string `json:"op,omitempty" yaml:"op,omitempty" toml:"op,omitempty"`
	Value uint64 `json:"value,omitempty" yaml:"value,omitempty" toml:"value,omitempty"`
}

type HealthConfig struct {
	Test []string `json:"Test,omitempty" yaml:"Test,omitempty" toml:"Test,omitempty"`

	Interval    time.Duration `json:"Interval,omitempty" yaml:"Interval,omitempty" toml:"Interval,omitempty"`          // Interval is the time to wait between checks.
	Timeout     time.Duration `json:"Timeout,omitempty" yaml:"Timeout,omitempty" toml:"Timeout,omitempty"`             // Timeout is the time to wait before considering the check to have hung.
	StartPeriod time.Duration `json:"StartPeriod,omitempty" yaml:"StartPeriod,omitempty" toml:"StartPeriod,omitempty"` // The start period for the container to initialize before the retries starts to count down.

	Retries int `json:"Retries,omitempty" yaml:"Retries,omitempty" toml:"Retries,omitempty"`
}

type HealthCheck struct {
	Start    time.Time `json:"Start,omitempty" yaml:"Start,omitempty" toml:"Start,omitempty"`
	End      time.Time `json:"End,omitempty" yaml:"End,omitempty" toml:"End,omitempty"`
	ExitCode int       `json:"ExitCode,omitempty" yaml:"ExitCode,omitempty" toml:"ExitCode,omitempty"`
	Output   string    `json:"Output,omitempty" yaml:"Output,omitempty" toml:"Output,omitempty"`
}

type Health struct {
	Status        string        `json:"Status,omitempty" yaml:"Status,omitempty" toml:"Status,omitempty"`
	FailingStreak int           `json:"FailingStreak,omitempty" yaml:"FailingStreak,omitempty" toml:"FailingStreak,omitempty"`
	Log           []HealthCheck `json:"Log,omitempty" yaml:"Log,omitempty" toml:"Log,omitempty"`
}

type Mount struct {
	Name        string
	Source      string
	Destination string
	Driver      string
	Mode        string
	RW          bool
}

type Config struct {
	Hostname          string              `json:"Hostname,omitempty" yaml:"Hostname,omitempty" toml:"Hostname,omitempty"`
	Domainname        string              `json:"Domainname,omitempty" yaml:"Domainname,omitempty" toml:"Domainname,omitempty"`
	User              string              `json:"User,omitempty" yaml:"User,omitempty" toml:"User,omitempty"`
	Memory            int64               `json:"Memory,omitempty" yaml:"Memory,omitempty" toml:"Memory,omitempty"`
	MemorySwap        int64               `json:"MemorySwap,omitempty" yaml:"MemorySwap,omitempty" toml:"MemorySwap,omitempty"`
	MemoryReservation int64               `json:"MemoryReservation,omitempty" yaml:"MemoryReservation,omitempty" toml:"MemoryReservation,omitempty"`
	KernelMemory      int64               `json:"KernelMemory,omitempty" yaml:"KernelMemory,omitempty" toml:"KernelMemory,omitempty"`
	CPUShares         int64               `json:"CpuShares,omitempty" yaml:"CpuShares,omitempty" toml:"CpuShares,omitempty"`
	CPUSet            string              `json:"Cpuset,omitempty" yaml:"Cpuset,omitempty" toml:"Cpuset,omitempty"`
	PortSpecs         []string            `json:"PortSpecs,omitempty" yaml:"PortSpecs,omitempty" toml:"PortSpecs,omitempty"`
	ExposedPorts      map[Port]struct{}   `json:"ExposedPorts,omitempty" yaml:"ExposedPorts,omitempty" toml:"ExposedPorts,omitempty"`
	PublishService    string              `json:"PublishService,omitempty" yaml:"PublishService,omitempty" toml:"PublishService,omitempty"`
	StopSignal        string              `json:"StopSignal,omitempty" yaml:"StopSignal,omitempty" toml:"StopSignal,omitempty"`
	StopTimeout       int                 `json:"StopTimeout,omitempty" yaml:"StopTimeout,omitempty" toml:"StopTimeout,omitempty"`
	Env               []string            `json:"Env,omitempty" yaml:"Env,omitempty" toml:"Env,omitempty"`
	Cmd               []string            `json:"Cmd" yaml:"Cmd" toml:"Cmd"`
	Healthcheck       *HealthConfig       `json:"Healthcheck,omitempty" yaml:"Healthcheck,omitempty" toml:"Healthcheck,omitempty"`
	DNS               []string            `json:"Dns,omitempty" yaml:"Dns,omitempty" toml:"Dns,omitempty"` // For Docker API v1.9 and below only
	Image             string              `json:"Image,omitempty" yaml:"Image,omitempty" toml:"Image,omitempty"`
	Volumes           map[string]struct{} `json:"Volumes,omitempty" yaml:"Volumes,omitempty" toml:"Volumes,omitempty"`
	VolumeDriver      string              `json:"VolumeDriver,omitempty" yaml:"VolumeDriver,omitempty" toml:"VolumeDriver,omitempty"`
	WorkingDir        string              `json:"WorkingDir,omitempty" yaml:"WorkingDir,omitempty" toml:"WorkingDir,omitempty"`
	MacAddress        string              `json:"MacAddress,omitempty" yaml:"MacAddress,omitempty" toml:"MacAddress,omitempty"`
	Entrypoint        []string            `json:"Entrypoint" yaml:"Entrypoint" toml:"Entrypoint"`
	SecurityOpts      []string            `json:"SecurityOpts,omitempty" yaml:"SecurityOpts,omitempty" toml:"SecurityOpts,omitempty"`
	OnBuild           []string            `json:"OnBuild,omitempty" yaml:"OnBuild,omitempty" toml:"OnBuild,omitempty"`
	Mounts            []Mount             `json:"Mounts,omitempty" yaml:"Mounts,omitempty" toml:"Mounts,omitempty"`
	Labels            map[string]string   `json:"Labels,omitempty" yaml:"Labels,omitempty" toml:"Labels,omitempty"`
	AttachStdin       bool                `json:"AttachStdin,omitempty" yaml:"AttachStdin,omitempty" toml:"AttachStdin,omitempty"`
	AttachStdout      bool                `json:"AttachStdout,omitempty" yaml:"AttachStdout,omitempty" toml:"AttachStdout,omitempty"`
	AttachStderr      bool                `json:"AttachStderr,omitempty" yaml:"AttachStderr,omitempty" toml:"AttachStderr,omitempty"`
	ArgsEscaped       bool                `json:"ArgsEscaped,omitempty" yaml:"ArgsEscaped,omitempty" toml:"ArgsEscaped,omitempty"`
	Tty               bool                `json:"Tty,omitempty" yaml:"Tty,omitempty" toml:"Tty,omitempty"`
	OpenStdin         bool                `json:"OpenStdin,omitempty" yaml:"OpenStdin,omitempty" toml:"OpenStdin,omitempty"`
	StdinOnce         bool                `json:"StdinOnce,omitempty" yaml:"StdinOnce,omitempty" toml:"StdinOnce,omitempty"`
	NetworkDisabled   bool                `json:"NetworkDisabled,omitempty" yaml:"NetworkDisabled,omitempty" toml:"NetworkDisabled,omitempty"`

	// This is no longer used and has been kept here for backward
	// compatibility, please use HostConfig.VolumesFrom.
	VolumesFrom string `json:"VolumesFrom,omitempty" yaml:"VolumesFrom,omitempty" toml:"VolumesFrom,omitempty"`
}

