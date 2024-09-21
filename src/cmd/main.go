package main

import (
	"fmt"
	"log"
	"time"

	"serve-ready/src/internal/cli"
	"serve-ready/src/internal/requirements"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("=== Project Requirements ==="))

	displaySystemInfo()

	fmt.Println("\n", color.New(color.FgWhite).SprintFunc()("---------------------------------------------------"))

	selectedFramework, _, selectedCache, selectedWebserver, err := cli.GetSelections()
	if err != nil {
		log.Fatalf("Error during selection: %v", err)
	}

	allPassed := checkRequirementsWithSpinner(selectedFramework, selectedCache, selectedWebserver)

	if allPassed {
		fmt.Println("\n", color.GreenString("Result: All requirements are met. (true)"))
	} else {
		fmt.Println("\n", color.RedString("Result: Requirements are not met. (false)"))
	}
}

func displaySystemInfo() {
	labelColor := color.New(color.FgMagenta).SprintFunc()
	valueColor := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	vMem, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	interfaces, _ := net.Interfaces()

	fmt.Printf("%s: %s (%s)\n", labelColor(bold("Operating System")), valueColor(hostInfo.OS), valueColor(hostInfo.Platform))
	fmt.Printf("%s: %s\n", labelColor(bold("Kernel Version")), valueColor(hostInfo.KernelVersion))
	fmt.Printf("%s: %d seconds\n", labelColor(bold("Uptime")), hostInfo.Uptime)

	fmt.Printf("\n%s: %s\n", labelColor(bold("Processor")), valueColor(cpuInfo[0].ModelName))
	fmt.Printf("%s: %d\n", labelColor(bold("Cores")), cpuInfo[0].Cores)
	fmt.Printf("%s: %.2f GB / %.2f GB used\n", labelColor(bold("RAM")), float64(vMem.Used)/1e9, float64(vMem.Total)/1e9)
	fmt.Printf("%s: %.2f GB / %.2f GB used\n", labelColor(bold("Disk Usage")), float64(diskInfo.Used)/1e9, float64(diskInfo.Total)/1e9)

	fmt.Println("\n", labelColor(bold("Network Interfaces")), ":")
	for _, iface := range interfaces {
		fmt.Printf("%s: %s, %s: %v\n", labelColor(bold("Name")), valueColor(iface.Name), labelColor(bold("IP")), iface.Addrs)
	}
}

func checkRequirementsWithSpinner(framework, cache, webserver string) bool {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	s.Suffix = " Checking requirements..."

	allPassed := requirements.CheckRequirements(framework, webserver, cache)

	s.Stop()

	labelColor := color.New(color.FgYellow).SprintFunc()
	valueColor := color.New(color.FgCyan).SprintFunc()
	checkMark := color.New(color.FgGreen).SprintFunc()("✔")
	crossMark := color.New(color.FgRed).SprintFunc()("✘")
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println("\n", bold("--- Requirements Check ---"))

	fmt.Printf("%s %s %s: %s\n\n", labelColor("Framework"), valueColor(framework), checkMark, "Selected framework")

	if !allPassed {
		fmt.Printf("%s %s %s: %s\n", labelColor("PHP"), crossMark, color.RedString("✘"), "PHP version mismatch or not installed")
	}

	fmt.Printf("%s %s %s: %s\n", labelColor("Cache"), valueColor(cache), checkMark, "Selected cache")
	fmt.Printf("%s %s %s: %s\n", labelColor("Web Server"), valueColor(webserver), checkMark, "Selected web server")

	if allPassed {
		fmt.Println("\n", color.GreenString("--- All requirements for %s met successfully! ---", framework))
	} else {
		fmt.Println("\n", color.RedString("--- Some requirements for %s are missing! ---", framework))
	}

	return allPassed
}
