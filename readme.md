# netcheck ;-)
simple tool that will check if a given protocol:host:port is reachable. can check multiple times an will exit(1) if not reachable. i created it to be able to block a script until a resource finished starting inside a docker network without the need to link ports to the host machine and write a loop.

## usage
    -host string
        destination host
    -port string
        destination port
    -protocol string
        protocol ('tcp' (default), 'tcp4' (IPv4-only), 'tcp6' (IPv6-only), 'udp', 'udp4' (IPv4-only), 'udp6' (IPv6-only), 'ip', 'ip4' (IPv4-only), 'ip6' (IPv6-only), 'unix', 'unixgram' and 'unixpacket') (default "tcp")
    -times int
            repeat this many times (default 3)
        
## dockerize
you can build a mini container (<4mb) of this tool simply by running `make`. it will build the go app using go:alpine (so you do not need to have go installed) and after that build a docker image that you can run like this:
    
    docker run --net=docker_default --rm --name netcheck netcheck --host=foo-db --port=3306 --times=13