package main

import (
	"fmt"
)

//сколько ip адресов в диапазоне, без учёта широковещательных
func maxHosts(mask int) int {

	max_ip := 1

	for i := 0; i < (32 - mask); i++ {
		max_ip = max_ip * 2
	}

	return max_ip - 2
}

func main() {
	ip_b := []byte{192, 168, 0, 1}

	mask := 24

	mask_b := []byte{255, 255, 255, 0}

	first_ip := []byte{0, 0, 0, 0}
	last_ip := []byte{0, 0, 0, 0}

	first_ip[0] = ip_b[0] & mask_b[0]
	first_ip[1] = ip_b[1] & mask_b[1]
	first_ip[2] = ip_b[2] & mask_b[2]
	first_ip[3] = ip_b[3]&mask_b[3] + 1

	last_ip[0] = first_ip[0] | ^mask_b[0]
	last_ip[1] = first_ip[1] | ^mask_b[1]
	last_ip[2] = first_ip[2] | ^mask_b[2]
	last_ip[3] = first_ip[3] | ^mask_b[3] - 1

	fmt.Println("ip адресов:", maxHosts(mask))
	fmt.Println("начальный адрес:", first_ip)
	fmt.Println("конечный адрес:", last_ip)

}
