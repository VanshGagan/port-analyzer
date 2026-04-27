# Port Analyzer

A custom-built SYN-based port scanner written in Go, featuring packet sniffing, banner grabbing, and basic OS detection.

## Features
- SYN-based port scanning using raw packets
- Packet sniffing and SYN-ACK filtering
- Banner grabbing (1024 bytes)
- Basic OS detection using TTL analysis
- Detection of common ports and services

## Usage
### Port-Analyzer (Normal)
```bash
sudo go run main.go <target-ip>
```
### Port-Analyzer (with banner grabbing)
```bash
sudo go run main.go <target-ip> -b
```


## Example Output

### Port analysis with banner grabbing
#### Input (with the ip of: scanme.nmap.org)
```bash
sudo go run main.go 45.33.32.156 -b
```
#### Output
```bash        
Operating system is probably: Linux/Mac/Android
                          
┌──────────────────────────────────────────┐
│  [FOUND] Port: 21     Name: FTP          │
└──────────────────────────────────────────┐

┌──────────────────────────────────────────┐
│  [FOUND] Port: 80     Name: HTTP         │
└──────────────────────────────────────────┐

┌──────────────────────────────────────────┐
│  [FOUND] Port: 22     Name: SSH          │
└──────────────────────────────────────────┐


Banner of port 80:
HTTP/1.1 200 OK



Banner of port 22:
SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13

--- scan finished ---

scanned in 5.86 seconds

```
- Predicted OS
- Found open-ports (22, 21, 80)
- Duration was 5.86 seconds
- Grabbed SSH and HTTP banner
- Identified OpenSSH version 6.6.1p1
- Detected Ubuntu system

***

## Performance

- Uses concurrent workers to scan multiple ports in parallel.
- Improves scanning efficiency, especially for larger port ranges. 
- Designed with scalability in mind for future extensions.

## Architecture
```bash 

port-analyzer/
├── main.go               # Main entry point: Worker pool, job queue, orchestrator
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── README.md             # Project description and instructions
├── network/              # Network-related code
│   └── sender.go         # Sends raw TCP SYN packets
│   └── sniffer.go        # Sniffs incoming TCP packets for SYN/ACK detection and TTL analysis
└── banner/               # banner grabbing folder
│   └── grabber.go        # grabs the banner of open ports
└── utils/                # Helper functions
    └── getIpAdress.go    # Automatically determines the local IP
    └── ImportantPorts.go # list of important ports
    └── loading.go        # loading animation
    └── osDetection.go    # simple skript to determine the operating system with TTL
```


## Disclaimer

This project was built for educational purposes only. 
Only scan systems you own or have explicit permission to test.

