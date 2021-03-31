# dns-exfil: coredns

The coredns conf file and zone file I currently run.

## iptables

It's best to run coredns as a normal user and redirect the privileged ports.

```bash
iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 53 -j REDIRECT --to-port 5353
ip6tables -t nat -A PREROUTING -i eth0 -p tcp --dport 53 -j REDIRECT --to-port 5353
iptables -t nat -A PREROUTING -i eth0 -p udp --dport 53 -j REDIRECT --to-port 5353
ip6tables -t nat -A PREROUTING -i eth0 -p udp --dport 53 -j REDIRECT --to-port 5353
```
