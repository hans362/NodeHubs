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
	"net"
	"strconv"
	"time"
)

func getHostname() string {
	hostInfo, _ := host.Info()
	return hostInfo.Hostname
}

func getUptime() string {
	hostInfo, _ := host.Info()
	return strconv.FormatUint(hostInfo.Uptime, 10)
}

func getPlatformVersion() string {
	hostInfo, _ := host.Info()
	return hostInfo.PlatformVersion
}

func getCpuModel() string {
	cpuInfo, _ := cpu.Info()
	return cpuInfo[0].ModelName
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
	return strconv.FormatFloat(loadInfo.Load1, 'f', 2, 64)
}

func getIpAddr() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

type Node struct {
	Timestamp       string
	Uptime          string
	PlatformVersion string
	CPUModel        string
	CPUPercent      string
	MemPercent      string
	SwapPercent     string
	DiskPercent     string
	Load            string
	IPAddr          string
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
		fmt.Printf("Error initializing app: %v\n", err)
		return
	}
	client, err := app.Database(ctx)
	if err != nil {
		fmt.Printf("Error initializing database client: %v\n", err)
		return
	}
	ref := client.NewRef("/")
	nodeRef := ref.Child(*friendlyName)
	if _, err := nodeRef.Push(ctx, &Node{
		Timestamp:       strconv.FormatInt(time.Now().Unix(), 10),
		Uptime:          getUptime(),
		PlatformVersion: getPlatformVersion(),
		CPUModel:        getCpuModel(),
		CPUPercent:      getCpuPercent(),
		MemPercent:      getMemPercent(),
		SwapPercent:     getSwapPercent(),
		DiskPercent:     getDiskPercent(),
		Load:            getLoad(),
		IPAddr:          getIpAddr(),
	}); err != nil {
		fmt.Printf("Error reporting data: %v\n", err)
	}
}
