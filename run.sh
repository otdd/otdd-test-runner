sudo iptables -t nat -A OUTPUT -p tcp -m owner --uid-owner 1987 -j ACCEPT
sudo iptables -t nat -A OUTPUT -p tcp -j REDIRECT --to-port 18746
./otdd-test-runner -h $OTDD_SERVER_HOST -p $OTDD_SERVER_PORT -u $USERNAME -t $TAG
