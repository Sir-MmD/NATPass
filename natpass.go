//MmD
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
    GREEN = "\033[32m"
    RED   = "\033[31m"
    BLUE  = "\033[34m"
    GOLD  = "\033[33m"
    CYAN  = "\033[36m"
    NC    = "\033[0m"
)

var (
	host  string
	port  int
	bind  int
	token string
	log   string
)

func banner() {
	asciiArt := RED+"    _   _____  __________                 \n" +
		"   / | / /   |/_  __/ __ \\____ ___________\n" +
		"  /  |/ / /| | / / / /_/ / __ `/ ___/ ___/\n" +
		" / /|  / ___ |/ / / ____/ /_/ (__  |__  ) \n" +
		"/_/ |_/_/  |_/_/ /_/    \\__,_/____/____/  \n" +
		GOLD+"       NAT Traversal with Xray-Core \n"+NC

	fmt.Println(asciiArt)
}

func help() {
	banner()
	fmt.Println(`Usage: app [OPTIONS]

Options:
  -s, --server        Run in server mode
  -c, --client        Run in client mode
  -h, --host <host>   Specify the host (required for both modes)
  -p, --port <port>   Specify the port (1-65535)
  -b, --bind <port>   Specify the bind port (1-65535)
  -t, --token <token> Specify the authentication token
  -l, --log <level>   Specify the log level (none, debug, info, warning, error). Default is warning.
      --help          Show this help message
`)
}

// Update server.json.sample with user input
func server_config(host string, port int, bind int, token string, log string) error {
	// Read the contents of server.json.sample
	sample_path := "assets/server.json.sample"
	content, err := ioutil.ReadFile(sample_path)
	if err != nil {
		return fmt.Errorf(RED+"Could not read sample file '%s': %v"+NC, sample_path, err)
	}

	// Replace placeholders with user input
	updated_value := string(content)
	updated_value = strings.Replace(updated_value, "SERVER_HOST", host, -1)
	updated_value = strings.Replace(updated_value, "SERVER_PORT", fmt.Sprintf("%d", port), -1)
	updated_value = strings.Replace(updated_value, "SERVER_BIND", fmt.Sprintf("%d", bind), -1)
	updated_value = strings.Replace(updated_value, "SERVER_TOKEN", token, -1)
	updated_value = strings.Replace(updated_value, "SERVER_LOG", log, -1)

	// Write the updated content to a new file (server.json)
	new_json := "assets/server.json"
	err = ioutil.WriteFile(new_json, []byte(updated_value), 0644)
	if err != nil {
		return fmt.Errorf(RED+"Could not write updated file '%s': %v"+NC, new_json, err)
	}

	return nil
}

// Update client.json.sample with user input
func client_config(host string, port int, token string, log string) error {
	// Read the contents of client.json.sample
	sample_path := "assets/client.json.sample"
	content, err := ioutil.ReadFile(sample_path)
	if err != nil {
		return fmt.Errorf(RED+"Could not read sample file '%s': %v"+NC, sample_path, err)
	}

	// Replace placeholders with user input
	updated_value := string(content)
	updated_value = strings.Replace(updated_value, "CLIENT_HOST", host, -1)
	updated_value = strings.Replace(updated_value, "CLIENT_PORT", fmt.Sprintf("%d", port), -1)
	updated_value = strings.Replace(updated_value, "CLIENT_BIND", fmt.Sprintf("%d", bind), -1)
	updated_value = strings.Replace(updated_value, "CLIENT_TOKEN", token, -1)
	updated_value = strings.Replace(updated_value, "CLIENT_LOG", log, -1)

	// Write the updated content to a new file (client.json)
	new_json := "assets/client.json"
	err = ioutil.WriteFile(new_json, []byte(updated_value), 0644)
	if err != nil {
		return fmt.Errorf(RED+"Could not write updated file '%s': %v"+NC, new_json, err)
	}

	return nil
}

// Run xray
func run_xray(mode string) error {
	// OS check
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		if mode == "server" {
			cmd = exec.Command("assets/xray.exe", "run", "-c", "assets/server.json")
		} else {
			cmd = exec.Command("assets/xray.exe", "run", "-c", "assets/client.json")
		}
	} else {
		if mode == "server" {
			cmd = exec.Command("assets/xray", "run", "-c", "assets/server.json")
		} else {
			cmd = exec.Command("assets/xray", "run", "-c", "assets/client.json")
		}
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(CYAN + "")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(RED+"Failed to run Xray-Core: %v"+NC, err)
	}
	return nil
}

func main() {
	server_switch := flag.Bool("server", false, "Run in server mode")
	client_switch := flag.Bool("client", false, "Run in client mode")

	flag.StringVar(&host, "host", "", "Specify the host")
	flag.IntVar(&port, "port", 0, "Specify the port (1-65535)")
	flag.IntVar(&bind, "bind", 0, "Specify the bind port (1-65535)")
	flag.StringVar(&token, "token", "", "Specify the authentication token")
	flag.StringVar(&log, "log", "warning", "Specify the log level (none, debug, info, warning, error)")

	flag.BoolVar(server_switch, "s", false, "Run in server mode")
	flag.BoolVar(client_switch, "c", false, "Run in client mode")
	flag.StringVar(&host, "h", "", "Specify the host")
	flag.IntVar(&port, "p", 0, "Specify the port")
	flag.IntVar(&bind, "b", 0, "Specify the bind port")
	flag.StringVar(&token, "t", "", "Specify the authentication token")
	flag.StringVar(&log, "l", "warning", "Specify the log level (none, debug, info, warning, error)")

	flag.Usage = help

	flag.Parse()

	if len(os.Args) == 1 {
		help()
		os.Exit(0)
	}

	if *server_switch && *client_switch {
		fmt.Println("Error: Cannot run in both server and client mode.")
		os.Exit(1)
	}
	if !*server_switch && !*client_switch {
		fmt.Println("Error: Must specify either server (-s) or client (-c) mode.")
		os.Exit(1)
	}

	if host == "" {
		fmt.Println("Error: Must specify a host using -h <host> for both modes.")
		os.Exit(1)
	}

	if port < 1 || port > 65535 {
		fmt.Println("Error: Port must be in range 1-65535.")
		os.Exit(1)
	}
	if bind < 1 || bind > 65535 {
		fmt.Println("Error: Bind port must be in range 1-65535.")
		os.Exit(1)
	}

	if token == "" {
		fmt.Println("Error: Must specify an authentication token using -t <token>.")
		os.Exit(1)
	}

	validLogLevels := map[string]bool{
		"none":    true,
		"debug":   true,
		"info":    true,
		"warning": true,
		"error":   true,
	}
	if !validLogLevels[log] {
		fmt.Println("Error: Invalid log level. Must be one of: none, debug, info, warning, error.")
		os.Exit(1)
	}

	if *server_switch {
		banner()

		fmt.Printf(GREEN+"Running in server mode - Host: %s, Port: %d, Bind Port: %d, Token: %s, Log Level: %s\n"+NC,
			host, port, bind, token, log)

		err := server_config(host, port, bind, token, log)
		if err != nil {
			fmt.Printf(RED+"Error updating server configuration: %v\n"+NC, err)
			os.Exit(1)
		}
		fmt.Println("Server configuration updated and saved to assets/server.json")

		err = run_xray("server")
		if err != nil {
			fmt.Printf(RED+"Error running Xray-Core: %v\n"+NC, err)
			os.Exit(1)
		}
		fmt.Println("Xray-Core is running successfully")
	}

	if *client_switch {
		banner()

		fmt.Printf(GREEN+"Running in client mode - Host: %s, Port: %d, Bind Port: %d, Token: %s, Log Level: %s\n"+NC,
			host, port, bind, token, log)

		err := client_config(host, port, token, log)
		if err != nil {
			fmt.Printf(RED+"Error updating client configuration: %v\n"+NC, err)
			os.Exit(1)
		}
		fmt.Println("Client configuration updated and saved to assets/client.json")

		err = run_xray("client")
		if err != nil {
			fmt.Printf(RED+"Error running Xray-Core: %v\n"+NC, err)
			os.Exit(1)
		}
	}
}
