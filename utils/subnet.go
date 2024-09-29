package utils

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// GenerateRandomIPsFromSubnets selects random IP addresses from the provided subnets
func GenerateRandomIPsFromSubnets(subnets []string, count int) ([]string, error) {
	var allIPs []string

	for _, subnet := range subnets {
		_, ipnet, err := net.ParseCIDR(subnet)
		if err != nil {
			return nil, fmt.Errorf("failed to parse subnet: %v", err)
		}

		// Collect all IPs in the range
		for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
			allIPs = append(allIPs, ip.String())
		}
	}

	// Create a new local random generator with the current time as the seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Shuffle and select random IPs
	rng.Shuffle(len(allIPs), func(i, j int) { allIPs[i], allIPs[j] = allIPs[j], allIPs[i] })

	if len(allIPs) < count {
		count = len(allIPs) // Adjust if there are fewer IPs than requested
	}

	return allIPs[:count], nil
}

// incrementIP increments an IP address by one
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
