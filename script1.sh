#!/usr/bin/env bash

get_interface_info() {
  echo "Net interface info:"

  for interface in $(ip -brief address show | awk '{print $1;}'); do
    echo "Interface: $interface"

    echo "Model: $(ethtool -i "$interface" | grep 'driver' | awk '{print $2}')"
    echo "Speed: $(ethtool "$interface" | grep 'Speed' | awk '{print $2}')"
    echo "Duplex mode: $(ethtool "$interface" | grep 'Duplex' | awk '{print $2}')"
    echo "Link: $(ethtool "$interface" | grep 'Link detected' | awk '{print $3}')"
    echo "MAC address: $(cat /sys/class/net/"$interface"/address)"
    echo ""
  done
}

get_ipv4_info() {
  echo "IPv4 info:"

  for interface in $(ip -brief address show | awk '{print $1;}'); do
    echo "Interface: $interface"

    ip -4 addr show |
      grep inet |
      awk '{print $2}' |
      while read -r line; do
        echo "IP/mask: $line"
      done

    echo -e "\nDefault gateway:"
    ip route | grep default | awk '{print $3}'

    echo -e "\nDNS servers:"
    grep nameserver /etc/resolv.conf | awk '{print $2}'
    echo ""
  done
}

opts=(interface_info ipv4_info quit)

while true; do
  PS3="Select from following: "

  select opt in "${opts[@]}"; do
    case $opt in
    interface_info)
      get_interface_info
      ;;
    ipv4_info)
      get_ipv4_info
      ;;
    quit)
      exit 0
      ;;
    *)
      error "Unexpected expression $opt"
      ;;
    esac
  done
done
