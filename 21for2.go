package main

import (
	"bufio"
	"fmt"
	"os"
)

// GİRİŞ PROGTAMI

func main() {
	k_adi := "alpt"
	k_sifre := "12345"

	giris_hak := 2

	fmt.Print("***login****")

	for true {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("kullanıcı adınızı giriniz =  ")
		scanner.Scan()
		giris_kadi := scanner.Text()

		fmt.Println("şifrenizi giriniz =   ")
		scanner.Scan()
		giris_sifre := scanner.Text()

		if k_adi != giris_kadi && k_sifre != giris_sifre {
			fmt.Println("kullanıcı adiniz ve şifreniz yanlış ")
			giris_hak--
		} else if k_adi == giris_kadi && k_sifre != giris_sifre {
			fmt.Println("şifreniz yanlış ")
			giris_hak--
		} else if k_adi != giris_kadi && k_sifre == giris_sifre {
			fmt.Println("kullanıcı adınız yanlış")
		} else {
			fmt.Println("hosgeldiniz")
			break
		}

		if giris_hak == 0 {
			fmt.Println("giris hakkınız bitti ")
			break

		}
	}

}
