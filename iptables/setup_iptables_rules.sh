#iptables -t nat -I OUTPUT -p tcp -d 127.0.0.1 -j ACCEPT
iptables -t nat -A OUTPUT -p tcp -m owner --uid-owner otdd-test-runner -j ACCEPT
iptables -t nat -A OUTPUT -p tcp -m owner --gid-owner otdd -j REDIRECT --to-port 18746
