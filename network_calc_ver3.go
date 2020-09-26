package main

import (
	"fmt"
	"math"
	"strconv"
)

//Функция для расчета ip адресов в диапазоне, без учёта широковещательных
func maxHosts(mask int) int {

	max_ip := 1
	//max кол. хостов: 2^(32-маску)
	for i := 0; i < (32 - mask); i++ { // и минус 2 широковещ.
		max_ip = max_ip * 2
	}

	return max_ip - 2
}

//Функция для перевода маски из вида 24 в вид 255.255.255.0
func maskToMask(mask int, mask_b *[]byte) {

	var x int = mask

	for j := x; j > 0; j-- {

		(*mask_b)[0] = (*mask_b)[0] + byte(math.Pow(2, (8-float64(j))))
		//2^(x-1) + 2^(x-2) ... 2^(x-8) перевод из числа в байты
		if (*mask_b)[0] >= 255 { //если один байт уже 255 переходим к следующему
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

//функция берет стринг "192.168.0.1/24" и извлекает байты и маску
func stringToInt(ipstr *string, mask *int, ip_b *[]byte) {

	var str string
	str = *ipstr
	a := []byte(str) //получаем из string срез байтов, один символ один байт

	var b1, b2, b3, b4, s_mask string

	for i := 0; i < len(str); i++ {

		if a[i] == 46 { // 46 это "."
			break //если встретили точку прекращаем заполнять b1
		}

		b1 = b1 + string(a[i])
	}

	for i := len(b1) + 1; i < len(str); i++ { // заполнение b2 начинаем с кол. символов b1 + точка

		if a[i] == 46 {
			break
		}

		b2 = b2 + string(a[i])
	}

	for i := len(b1) + len(b2) + 2; i < len(str); i++ {

		if a[i] == 46 {
			break
		}

		b3 = b3 + string(a[i])
	}

	for i := len(b1) + len(b2) + len(b3) + 3; i < len(str); i++ {

		if a[i] == 46 {
			break
		}
		if a[i] == 47 { // 47 это "/"
			break
		}

		b4 = b4 + string(a[i])
	}

	for i := 0; i < len(str); i++ {

		if a[i] == 47 {
			s_mask = string(a[i+1]) + string(a[i+2])
		}
	}

	ip1, _ := strconv.Atoi(b1) //переводим string в int
	ip2, _ := strconv.Atoi(b2)
	ip3, _ := strconv.Atoi(b3)
	ip4, _ := strconv.Atoi(b4)
	n_mask, _ := strconv.Atoi(s_mask)

	*mask = n_mask

	(*ip_b)[0] = byte(ip1)
	(*ip_b)[1] = byte(ip2)
	(*ip_b)[2] = byte(ip3)
	(*ip_b)[3] = byte(ip4)

}

func main() {
	var ipstr string
	fmt.Println("Введите адрес и маску сети: (пример 192.168.0.1/24)")
	fmt.Scanf("%s", &ipstr)

	mask := 24

	ip_b := []byte{192, 168, 0, 1}

	stringToInt(&ipstr, &mask, &ip_b)

	max_ip := maxHosts(mask)

	mask_b := []byte{0, 0, 0, 0} // маска в виде байтов

	maskToMask(mask, &mask_b)

	first_ip := []byte{0, 0, 0, 0}
	last_ip := []byte{0, 0, 0, 0}

	first_ip[0] = ip_b[0] & mask_b[0] //сетвой адрес: побитовое AND между адресом и маской
	first_ip[1] = ip_b[1] & mask_b[1]
	first_ip[2] = ip_b[2] & mask_b[2]
	first_ip[3] = ip_b[3]&mask_b[3] + 1

	last_ip[0] = first_ip[0] | ^mask_b[0]
	last_ip[1] = first_ip[1] | ^mask_b[1] //широковещ. адрес: побитовое OR между сетевым адресом
	last_ip[2] = first_ip[2] | ^mask_b[2] //и перевернутой маской
	last_ip[3] = first_ip[3] | ^mask_b[3] - 1

	fmt.Println(ipstr)
	fmt.Printf("начальный адрес: %v.%v.%v.%v\n", first_ip[0], first_ip[1], first_ip[2], first_ip[3])
	fmt.Printf("конечный адрес: %v.%v.%v.%v\n", last_ip[0], last_ip[1], last_ip[2], last_ip[3])
	fmt.Println("ip адресов:", max_ip)

}
