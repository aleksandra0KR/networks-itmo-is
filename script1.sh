#!/usr/bin/env bash

get_network_info() {
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

get_network_info
