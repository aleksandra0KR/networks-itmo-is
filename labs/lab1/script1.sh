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

scenario_one() {
  for interface in $(ip -brief address show | awk '{print $1;}'); do
    ip="10.100.0.2"
    mask="255.255.255.0"
    gate="10.100.0.1"
    dns="8.8.8.8"

    ip a add $ip/$mask dev "$interface"
    ip r delete default dev "$interface"
    ip r add default via $gate dev "$interface"
    echo "nameserver ${dns}" | tee /etc/resolv.conf
  done
}

undo_scenario_one() {
  for interface in $(ip -brief address show | awk '{print $1;}'); do
    ip="10.100.0.2"
    mask="255.255.255.0"
    old_gate="192.168.0.1"
    dns="8.8.8.8"

    ip a delete $ip/$mask dev "$interface"
    ip r delete default dev "$interface"
    ip r add default via $old_gate dev "$interface"
    true "" >/etc/resolv.conf
  done
}

scenario_two() {
  for interface in $(ip -brief address show | awk '{print \$1;}'); do
    nmcli dev set "\$interface" managed yes
    nmcli con add type ethernet ifname "\$interface" con-name "\$interface-dhcp"
    nmcli con modify "\$interface-dhcp" ipv4.method auto
    nmcli con up "\$interface-dhcp"
  done
}

opts=(interface_info ipv4_info scenario_one undo scenario_two quit)

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
    scenario_one)
      scenario_one
      ;;
    undo)
      undo_scenario
      ;;
    scenario_two)
      scenario_two
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
