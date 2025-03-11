# NATPass
### NAT Traversal with Xray-Core
![App Screenshot](https://raw.githubusercontent.com/Sir-MmD/NATPass/refs/heads/main/pics/natpass.png)

## What is NATPass ?
NATPass is a NAT Traversal tool that uses Xray-Core's Reverse Tunnel feature to bypass NAT limitation on a system.

It supports both TCP and UDP

## Usage
```
  app [OPTIONS]

  -s, --server        Run in server mode
  -c, --client        Run in client mode
  -h, --host <host>   Specify the host (required for both modes)
  -p, --port <port>   Specify the port (1-65535)
  -b, --bind <port>   Specify the bind port (1-65535)
  -t, --token <token> Specify the authentication token
  -l, --log <level>   Specify the log level (none, debug, info, warning, error). Default is warning.
      --help          Show this help message
```

## Donate
🔹USDT-TRC20: ```TEUsjU22rAb3TZNaecEBJAhZXQsFZYU7vv```

🔹TRX: ```TEUsjU22rAb3TZNaecEBJAhZXQsFZYU7vv```

🔹LTC: ```LS7rJ6nMWwgw9FWpMjSnYa1bAPdzK7bJLM```

🔹BTC: ```1D4cSHY95FoHExSicKYmjeVzdLLxhjTTqs```

🔹ETH: ```0x03fde84612e0d572db7a18efeeec590ad3fa5dfb```

## Supported OS

| OS  | CPU Type         | ROOT Required |
|:--------------:|:----------------:|:----------------:
|![App Screenshot](https://raw.githubusercontent.com/Sir-MmD/NATPass/refs/heads/main/pics/windows.png)<div align="center">**Windows**| x86_64 | NO
| ![App Screenshot](https://raw.githubusercontent.com/Sir-MmD/NATPass/refs/heads/main/pics/linux.png)<div align="center">**Linux** | x86_64 / aarch64 | NO
| ![App Screenshot](https://raw.githubusercontent.com/Sir-MmD/NATPass/refs/heads/main/pics/android.png)                                            <div align="center">**Android** | aarch64 / aarch64 | NO

## Diagram
![App Screenshot](https://raw.githubusercontent.com/Sir-MmD/NATPass/refs/heads/main/pics/diagram.png)

## Bridge Server Example
```
$ netpass -s -h 0.0.0.0 -p 9090 -b 9191 -t MySecurePassword
```
### -h: 
Listening ip addresses
### -p: 
Listening port for Reverse Tunnel
### -b: 
Exposed port to access System Behind NAT
### -t: 
A secure password 

## Client Example (System behind NAT)
```
$ netpass -c -h SERVER_IP -p 9090 -b 80 -t MySecurePassword
```
### -h: 
Public IP address of Bridge Server 
### -p: 
Same port entered on Bridge Server
### -b: 
Desired port to expose
### -t: 
Same secure password entered on Bridge Server

### Log
By default, NETPass uses "none" as log level but you can change it with ```"-l``` switch to:

- debug
- info
- warning
- error

## Tunnel Tweaking 
By default, NATPass uses "VLESS + WS" as for the Reverse Tunnel, You can modify it to your desired protocol and transport inside ```assets/server.json.sample``` and ```assets/client.json.sample```
## Build
```bash
git clone https://github.com/Sir-MmD/NATPass && cd NATPass
```
```bash
go build netpass.go
```
After compilation, you need to download Xray-Core and extract it inside ```assets``` folder: https://github.com/XTLS/Xray-core
