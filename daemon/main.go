package main

import (
	"context"
	"firebase.google.com/go/v4"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"google.golang.org/api/option"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getHostId() string {
	hostId, _ := host.HostID()
	return hostId
}

func getHostname() string {
	hostInfo, _ := host.Info()
	return hostInfo.Hostname
}

func getUptime() string {
	hostUptime, _ := host.Uptime()
	return strconv.FormatUint(hostUptime, 10)
}

func getOs() string {
	hostInfo, _ := host.Info()
	return hostInfo.OS
}

func getPlatform() string {
	hostInfo, _ := host.Info()
	return hostInfo.Platform
}

func getPlatformVersion() string {
	hostInfo, _ := host.Info()
	return hostInfo.PlatformVersion
}

func getCpuModel() string {
	cpuInfo, _ := cpu.Info()
	return cpuInfo[0].ModelName
}

func getCpuCores() string {
	cpuCores, _ := cpu.Counts(true)
	return strconv.Itoa(cpuCores)
}

func getCpuPercent() string {
	percent, _ := cpu.Percent(time.Second, false)
	return strconv.FormatFloat(percent[0], 'f', 2, 64)
}

func getMemPercent() string {
	memInfo, _ := mem.VirtualMemory()
	return strconv.FormatFloat(memInfo.UsedPercent, 'f', 2, 64)
}

func getSwapPercent() string {
	swapInfo, _ := mem.SwapMemory()
	return strconv.FormatFloat(swapInfo.UsedPercent, 'f', 2, 64)
}

func getDiskPercent() string {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return strconv.FormatFloat(diskInfo.UsedPercent, 'f', 2, 64)
}

func getLoad() string {
	loadInfo, _ := load.Avg()
	return strings.Join([]string{strconv.FormatFloat(loadInfo.Load1, 'f', 2, 64), strconv.FormatFloat(loadInfo.Load5, 'f', 2, 64), strconv.FormatFloat(loadInfo.Load15, 'f', 2, 64)}, " ")
}

func getIpInfo() string {
	req, _ := http.Get("http://ip-api.com/json/")
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	return string(body)
}

type Node struct {
	FriendlyName    string
	Timestamp       string
	Hostname        string
	Uptime          string
	OS              string
	Platform        string
	PlatformVersion string
	CPUModel        string
	CPUCores        string
	CPUPercent      string
	MemPercent      string
	SwapPercent     string
	DiskPercent     string
	Load            string
	IPInfo          string
}

var dbURL = flag.String("db", "https://node-hubs-default-rtdb.firebaseio.com", "Firebase Realtime Database URL")
var credentialsPath = flag.String("auth", "/root/.nodehubs/serviceAccountKey.json", "Firebase Service Account Key Path")
var friendlyName = flag.String("name", getHostname(), "Friendly Name for the Server")

func main() {
	flag.Parse()
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: *dbURL,
	}
	opt := option.WithCredentialsFile(*credentialsPath)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		fmt.Printf("Error initializing firebase app: %v\n", err)
		return
	}
	client, err := app.Database(ctx)
	if err != nil {
		fmt.Printf("Error initializing realtime database client: %v\n", err)
		return
	}
	ref := client.NewRef("/nodes")
	nodeRef := ref.Child(getHostId())
	if _, err := nodeRef.Push(ctx, &Node{
		FriendlyName:    *friendlyName,
		Timestamp:       strconv.FormatInt(time.Now().Unix(), 10),
		Hostname:        getHostname(),
		Uptime:          getUptime(),
		OS:              getOs(),
		Platform:        getPlatform(),
		PlatformVersion: getPlatformVersion(),
		CPUModel:        getCpuModel(),
		CPUCores:        getCpuCores(),
		CPUPercent:      getCpuPercent(),
		MemPercent:      getMemPercent(),
		SwapPercent:     getSwapPercent(),
		DiskPercent:     getDiskPercent(),
		Load:            getLoad(),
		IPInfo:          getIpInfo(),
	}); err != nil {
		fmt.Printf("Failed to report data to firebase: %v\n", err)
		return
	}
	if err := nodeRef.Update(ctx, map[string]interface{}{
		"LastUpdate": strconv.FormatInt(time.Now().Unix(), 10),
	}); err != nil {
		fmt.Printf("Failed to report data to firebase: %v\n", err)
		return
	}
	fmt.Printf("Successfully reported data to firebase\n")
}
