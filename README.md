# Port Analyzer

## Overview

This project is a custom-built TCP SYN port scanner written in Go.

It is designed to efficiently detect open ports without completing a full TCP three-way handshake.

Instead of establishing full connections, the scanner sends raw **SYN** packets and analyzes responses using a packet sniffer.

***
## How it works

The scanner follows the logic of a classic **SYN** scan:

We send a raw **SYN** packet to a specific port of the target.
Then we catch the response of the target-port and look if it was a **SYN-ACK** packet.

If this is true then the port is open.


## Efficiency    

To reduce running time, the scanner uses goroutines and worker-pools for parallel working. 
At the moment it scanns the most important ports of the target (53 ports).

***
## Architecture

```bash

port-analyzer/
├── main.go               # Main entry point: Worker pool, job queue, orchestrator
├── main                  # (Compiled binary after go build)
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── README.md             # Project description and instructions
├── network/              # Network-related code
│   ├── sender.go         # Sends raw TCP SYN packets
│   └── sniffer.go        # Sniffs incoming TCP packets for SYN/ACK detection
└── utils/                # Helper functions
    └── getIpAdress.go    # Automatically determines the local IP
    
```

***
# ⚠️ Disclaimer
This project was built for educational purposes only, mainly for my own learning.
