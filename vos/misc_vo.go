package vos

import (
	"net"

	"encoding/json"

	"github.com/docker/docker/api/types/swarm"
)

type DockerInfo struct {
	ID                 string
	Containers         int
	ContainersRunning  int
	ContainersPaused   int
	ContainersStopped  int
	Images             int
	Driver             string
	DriverStatus       [][2]string
	SystemStatus       [][2]string
	Plugins            PluginsInfo
	MemoryLimit        bool
	SwapLimit          bool
	KernelMemory       bool
	CPUCfsPeriod       bool `json:"CpuCfsPeriod"`
	CPUCfsQuota        bool `json:"CpuCfsQuota"`
	CPUShares          bool
	CPUSet             bool
	IPv4Forwarding     bool
	BridgeNfIptables   bool
	BridgeNfIP6tables  bool `json:"BridgeNfIp6tables"`
	Debug              bool
	OomKillDisable     bool
	ExperimentalBuild  bool
	NFd                int
	NGoroutines        int
	SystemTime         string
	ExecutionDriver    string
	LoggingDriver      string
	CgroupDriver       string
	NEventsListener    int
	KernelVersion      string
	OperatingSystem    string
	OSType             string
	Architecture       string
	IndexServerAddress string
	RegistryConfig     *ServiceConfig
	SecurityOptions    []string
	NCPU               int
	MemTotal           int64
	DockerRootDir      string
	HTTPProxy          string `json:"HttpProxy"`
	HTTPSProxy         string `json:"HttpsProxy"`
	NoProxy            string
	Name               string
	Labels             []string
	ServerVersion      string
	ClusterStore       string
	ClusterAdvertise   string
	Isolation          string
	InitBinary         string
	DefaultRuntime     string
	LiveRestoreEnabled bool
	Swarm              swarm.Info
}

type PluginsInfo struct {
	Volume []string
	Network []string
	Authorization []string
}

type ServiceConfig struct {
	InsecureRegistryCIDRs []*NetIPNet
	IndexConfigs          map[string]*IndexInfo
	Mirrors               []string
}

type NetIPNet net.IPNet

func (ipnet *NetIPNet) MarshalJSON() ([]byte, error) {
	return json.Marshal((*net.IPNet)(ipnet).String())
}

func (ipnet *NetIPNet) UnmarshalJSON(b []byte) (err error) {
	var ipnetStr string
	if err = json.Unmarshal(b, &ipnetStr); err == nil {
		var cidr *net.IPNet
		if _, cidr, err = net.ParseCIDR(ipnetStr); err == nil {
			*ipnet = NetIPNet(*cidr)
		}
	}
	return
}

type IndexInfo struct {
	Name     string
	Mirrors  []string
	Secure   bool
	Official bool
}
