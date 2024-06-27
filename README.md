# Docker DNS (internal as ClipDNS)

A simple DNS server to provide Docker containers' IP address.


Usage:

1. Rename `config.example.yml` to `config.yml`
2. Start the server
3. Now you can lookup `<container_name>.<suffix>` to get containers' IP address.

Limitation:

This project current support IPv4 (Type A Questions) **Only**.

This project needs `docker` user group privileges to access Docker Engine socket.