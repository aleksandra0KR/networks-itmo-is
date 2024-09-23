package network

import "errors"

type Network struct {
	UsersAmount int
}

func FromInts(amounts []int) ([]Network, error) {
	networks := make([]Network, 0, len(amounts))
	for _, amount := range amounts {
		if amount < 0 {
			return nil, errors.New("amount can't be negative")
		}

		networks = append(networks, Network{UsersAmount: amount})
	}

	return networks, nil
}
