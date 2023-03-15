/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// sysInfoCmd represents the sysInfo command
var sysInfoCmd = &cobra.Command{
	Use:     "sysInfo",
	Aliases: []string{"system", "sysinf"},
	Short:   "Prints out information on the Operating System",
	Long: `Prints out information on the Operating System
A specifier is avilable: --spec [OPTIONS]

Available options for the specifier --spec:
	memory - print out the memory information (alias: mem)
	oshost - print out the host information (alias: host)
	all - print out the OS information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("System information called")

		spec, _ := cmd.Flags().GetString("spec")
		memStat, _ := mem.VirtualMemory()
		oshost, _ := host.Info()

		switch spec {
		case "host", "oshost":
			OSHostInformation(oshost)
		case "memory", "mem":
			OSMemoryInformation(memStat)
		case "all":
			OSHostInformation(oshost)
			OSMemoryInformation(memStat)
		default:
			log.Fatal("Unknown specifier...")
		}
	},
}

func init() {
	rootCmd.AddCommand(sysInfoCmd)
	sysInfoCmd.Flags().String("spec", "all", "specific information on the Operating System")
}

func OSHostInformation(oshost *host.InfoStat) error {
	fmt.Println("\nHost:")
	fmt.Printf("HostID: %v\n", oshost.HostID)
	fmt.Printf("Hostname: %v\n", oshost.Hostname)
	fmt.Printf("Procs: %v\n", oshost.Procs)
	fmt.Printf("Uptime: %v\n", oshost.Uptime)
	fmt.Printf("OS: %v\n", oshost.OS)
	fmt.Printf("\t..Platform: %v\n", oshost.Platform)
	fmt.Printf("\t..Platform Family: %v\n", oshost.PlatformFamily)
	fmt.Printf("\t..Platform Version: %v\n", oshost.PlatformVersion)

	return nil
}

func OSMemoryInformation(memStat *mem.VirtualMemoryStat) error {
	fmt.Println("\nOS Memory:")
	fmt.Printf("Total amount of RAM: %v\n", memStat.Total)
	fmt.Printf("RAM available: %v\n", memStat.Available)
	fmt.Printf("RAM used: %v\n\t..in percents: %.2f\n", memStat.Used, memStat.UsedPercent)

	return nil
}
