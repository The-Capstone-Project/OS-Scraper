package main

import (
	"flag"
	"fmt"
	"os"

	"os-scrapper/agents/RunningServicesAndDaemons"
	"os-scrapper/agents/env_info"
	"os-scrapper/agents/kernel_probe"
	"os-scrapper/agents/osinfo"
	"os-scrapper/agents/shellhistory"
	"os-scrapper/agents/shellinfo"
)

func main() {
	// Define flags
	all := flag.Bool("all", false, "Run all modules")
	osinfoFlag := flag.Bool("osinfo", false, "Run OS info module")
	kernelProbeFlag := flag.Bool("kernel_probe", false, "Run kernel probe module")
	envInfoFlag := flag.Bool("env_info", false, "Run environment info module")
	servicesFlag := flag.Bool("services", false, "Run RunningServicesAndDaemons module")
	shellHistoryFlag := flag.Bool("shellhistory", false, "Run shell history module")
	shellInfoFlag := flag.Bool("shellinfo", false, "Run shell info module")
	commandsFile := flag.String("f", "agents/osinfo/commands.json", "Path to the commands JSON file")

	flag.Parse()

	// Run all modules if --all is set
	if *all {
		fmt.Println("[OUTPUT] Running all modules:")
		env_info.Run()
		kernel_probe.Run()
		osinfo.Run(*commandsFile)
		RunningServicesAndDaemons.Run()
		shellhistory.Run()
		shellinfo.Run()
		return
	}

	// Run specific modules based on flags
	if *osinfoFlag {
		fmt.Println("[OUTPUT] osinfo:")
		osinfo.Run(*commandsFile)
	}
	if *kernelProbeFlag {
		fmt.Println("[OUTPUT] kernel_probe:")
		kernel_probe.Run()
	}
	if *envInfoFlag {
		fmt.Println("[OUTPUT] env_info:")
		env_info.Run()
	}
	if *servicesFlag {
		fmt.Println("[OUTPUT] RunningServicesAndDaemons:")
		RunningServicesAndDaemons.Run()
	}
	if *shellHistoryFlag {
		fmt.Println("[OUTPUT] shellhistory:")
		shellhistory.Run()
	}
	if *shellInfoFlag {
		fmt.Println("[OUTPUT] shellinfo:")
		shellinfo.Run()
	}

	// If no flags provided, show usage
	if !*all && !*osinfoFlag && !*kernelProbeFlag && !*envInfoFlag && !*servicesFlag && !*shellHistoryFlag && !*shellInfoFlag {
		fmt.Println("Usage: runner --all or runner --<module>")
		flag.Usage()
		os.Exit(1)
	}
}
