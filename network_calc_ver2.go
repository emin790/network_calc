package main

import 	"fmt"
import  "math"

//Функция для расчета ip адресов в диапазоне, без учёта широковещательных
func maxHosts(mask int) int {

	max_ip := 1

	for i := 0; i < (32 - mask); i++ {
		max_ip = max_ip * 2
	}

	return max_ip - 2
}

//Функция для перевода маски из вида 24 в вид 255.255.255.0
func maskToMask (mask int, mask_b *[]byte){

	var x int = mask
	
	for j := x; j > 0; j-- {
	
	(*mask_b)[0] = (*mask_b)[0] + byte(math.Pow(2, (8-float64(j))))
	
	if (*mask_b)[0] >= 255 {
	   	x = x - 8
	   	for j := x; j > 0; j-- {
		(*mask_b)[1] = (*mask_b)[1] + byte(math.Pow(2, (8-float64(j))))
		}
			
	if (*mask_b)[1] >= 255 {
		x = x - 8
		for j := x; j > 0; j-- {
			(*mask_b)[2] = (*mask_b)[2] + byte(math.Pow(2, (8-float64(j))))
		}
			
	if (*mask_b)[2] >= 255 {
		x = x - 8
		for j := x; j > 0; j-- {
			(*mask_b)[3] = (*mask_b)[3] + byte(math.Pow(2, (8-float64(j))))
		}
	  }
	}
      }
   }
}

func main() {
	ip_b := []byte{192, 168, 0, 1}

	mask := 24
	max_ip := maxHosts(mask)
	
	mask_b := []byte{0, 0, 0, 0}
	
	maskToMask(mask, &mask_b)
	
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

	fmt.Println("ip адресов:", max_ip)
	fmt.Println("начальный адрес:", first_ip)
	fmt.Println("конечный адрес:", last_ip)

}

ip адресов: 254
начальный адрес: [192 168 0 1]
конечный адрес: [192 168 0 254]