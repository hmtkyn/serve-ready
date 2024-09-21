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
	// Bold fonksiyonu oluşturuluyor
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("=== Project Requirements ==="))

	// Sistem bilgilerini al ve göster
	displaySystemInfo()

	// Boşluk veya ayrım çizgisi ekleyelim
	fmt.Println("\n", color.New(color.FgWhite).SprintFunc()("---------------------------------------------------"))

	// Dinamik seçenekler oluşturuluyor
	selectedFramework, selectedDatabase, selectedCache, selectedWebserver, err := cli.GetSelections()
	if err != nil {
		log.Fatalf("Error during selection: %v", err)
	}

	// Gereksinimlerin kontrolü
	allPassed := checkRequirementsWithSpinner(selectedFramework, selectedDatabase, selectedCache, selectedWebserver)

	if allPassed {
		fmt.Println("\n", color.GreenString("Result: All requirements are met. (true)"))
	} else {
		fmt.Println("\n", color.RedString("Result: Requirements are not met. (false)"))
	}
}

// Sistem bilgilerini almak ve göstermek için fonksiyon
func displaySystemInfo() {
	// Renkli ve bold stilleri oluşturuluyor
	labelColor := color.New(color.FgMagenta).SprintFunc() // Magenta renk etiketler için
	valueColor := color.New(color.FgCyan).SprintFunc()    // Cyan renk değerler için
	bold := color.New(color.Bold).SprintFunc()            // Bold stili oluşturma

	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	vMem, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	interfaces, _ := net.Interfaces()

	// Doğru formatta printf kullanımıyla hataları düzeltiyoruz
	fmt.Printf("%s: %s (%s)\n", labelColor(bold("Operating System")), valueColor(hostInfo.OS), valueColor(hostInfo.Platform))
	fmt.Printf("%s: %s\n", labelColor(bold("Kernel Version")), valueColor(hostInfo.KernelVersion))
	fmt.Printf("%s: %d seconds\n", labelColor(bold("Uptime")), hostInfo.Uptime)

	fmt.Printf("\n%s: %s\n", labelColor(bold("Processor")), valueColor(cpuInfo[0].ModelName))
	fmt.Printf("%s: %d\n", labelColor(bold("Cores")), cpuInfo[0].Cores)
	fmt.Printf("%s: %.2f GB / %.2f GB used\n", labelColor(bold("RAM")), float64(vMem.Used)/1e9, float64(vMem.Total)/1e9)
	fmt.Printf("%s: %.2f GB / %.2f GB used\n", labelColor(bold("Disk Usage")), float64(diskInfo.Used)/1e9, float64(diskInfo.Total)/1e9)

	// Ağ arayüzlerini düzgün formatta gösteriyoruz
	fmt.Println("\n", labelColor(bold("Network Interfaces")), ":")
	for _, iface := range interfaces {
		fmt.Printf("%s: %s, %s: %v\n", labelColor(bold("Name")), valueColor(iface.Name), labelColor(bold("IP")), iface.Addrs)
	}
}

// Gereksinim kontrolü sırasında spinner ve animasyonlar ekleyen fonksiyon
func checkRequirementsWithSpinner(framework, database, cache, webserver string) bool {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond) // Yeni bir spinner yaratılıyor
	s.Start()                                                    // Spinner başlatılıyor
	s.Suffix = " Checking requirements..."                       // Spinner'ın yanında gösterilecek metin

	// Gereksinimlerin kontrolü
	allPassed := requirements.CheckRequirements(framework, database, cache, webserver)

	s.Stop() // Gereksinim kontrolü bittiğinde spinner durduruluyor

	// Gereksinimlerin sonuçlarını ekranda güzel bir formatla göster
	labelColor := color.New(color.FgYellow).SprintFunc()
	valueColor := color.New(color.FgCyan).SprintFunc()      // Renkleri ekliyoruz
	checkMark := color.New(color.FgGreen).SprintFunc()("✔") // Check işareti için yeşil renk
	crossMark := color.New(color.FgRed).SprintFunc()("✘")   // Çarpı işareti için kırmızı renk

	// Bold fonksiyonunu kullanarak Requirements Check yazısı
	bold := color.New(color.Bold).SprintFunc() // Bold fonksiyonunu burada da kullanıyoruz
	fmt.Println("\n", bold("--- Requirements Check ---"))

	// Framework kontrolü
	fmt.Printf("%s %s %s: %s\n\n", labelColor("Framework"), valueColor(framework), checkMark, "Selected framework")

	// PHP veya diğer kontroller
	if !allPassed {
		fmt.Printf("%s %s %s: %s\n", labelColor("PHP"), crossMark, color.RedString("✘"), "PHP version mismatch or not installed")
	}

	// Database kontrolü
	fmt.Printf("%s %s %s: %s\n", labelColor("Database"), valueColor(database), checkMark, "Selected database")

	// Cache kontrolü
	fmt.Printf("%s %s %s: %s\n", labelColor("Cache"), valueColor(cache), checkMark, "Selected cache")

	// Web Server kontrolü
	fmt.Printf("%s %s %s: %s\n", labelColor("Web Server"), valueColor(webserver), checkMark, "Selected web server")

	// Sonuç çıktısı
	if allPassed {
		fmt.Println("\n", color.GreenString("--- All requirements for %s met successfully! ---", framework))
	} else {
		fmt.Println("\n", color.RedString("--- Some requirements for %s are missing! ---", framework))
	}

	return allPassed
}
