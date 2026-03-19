# Port Analyzer

## My idea

I am currently building an efficient port scanner based on TCP-protocol. The current port scanner in main.go attempts a 3-way handshake with every port using a worker pool of 40 workers. 

***

This means I send a ***SYN*** packet, the port sends back a ***SYN/ACK***, and I then send an ***ACK*** again; thus, the connection is established. 

However, this is very inefficient. First of all, we only know that a port is closed once the timeout has expired, which causes the program to lose a lot of time.

Therefore, my idea is to develop a program that sends a ***SYN*** packet and subsequently sniffs for a ***SYN/ACK***. If this arrives, the port is immediately marked as open without building a 3-way handshake connection. However, if a ***RST*** comes back, then this means the port is closed; if nothing comes back, then the port is filtered (meaning it is being blocked). This way, my program can become much faster. Therefore:

***

## Step 1: Build a Sniffer (Almost finished)

Build a sniffer that sniffs all packets where they come from and where they need to go. This allows us to easily catch ***SYN/ACK*** packets. We can already test this with our basic port analyzer: since it performs a 3-way handshake, we should be able to see directly with the sniffer that a ***SYN/ACK*** is coming to us.

***

## Step 2: Send a Raw Packet

Next, we want to send a simple ***SYN*** packet to the port in order to receive the ***SYN/ACK*** packet. To do this, we must be able to construct such a packet ourselves.

***

For all these steps, the **gopacket** library is perfect, and I will be working with it.
