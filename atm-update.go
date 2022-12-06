//Program ATM Sederhana dengan menggunakan bahasa Golang
//Program-Author : Muhammad Fajar Baihaqi
/*
.> Deskripsi Program :
Didalam Program ATM ini terdapat 6 Buah fitur yang biasanya dijumpai pada atm umumnya seperti Cek Saldo Pengguna, Tarik Tunai, Setor Tunai , Transfer Antar pengguna
,Ubah PIN dan Keluar. Pada program ini juga terdapat kasus dimana terdapat 5 orang pengguna ATM beserta data yang dibutuhkan.

Untuk kekurangan dari program ini mohon dimaafkan
Sekian Terima Kasih - Author :)
*/

package main

import (
	"fmt"
	"strings"
)

const (
	PIN_LENGTH = 4
)

type Account struct {
	name    string
	balance int
	pin     string
}

func main() {
	accounts := []Account{
		Account{
			name:    "Abdurahman Aziz",
			balance: 1000000,
			pin:     "1234",
		},
		Account{
			name:    "Bernadhetta Ira Tri Lita",
			balance: 500000,
			pin:     "4321",
		},
		Account{
			name:    "Firman Nurcahyo",
			balance: 500000,
			pin:     "5678",
		},
		Account{
			name:    "Muhammad Fajar Baihaqi",
			balance: 500000,
			pin:     "8765",
		},
		Account{
			name:    "Rahadani Syifariani",
			balance: 500000,
			pin:     "9012",
		},
	}

	// Menerima input user
	var userInput string
	fmt.Print("Masukkan PIN ATM Anda : ")
	fmt.Scanln(&userInput)

	// Mencari account yang cocok dengan PIN yang dimasukkan user
	var account Account
	for _, acc := range accounts {
		if acc.pin == userInput {
			account = acc
			break
		}
	}

	// Jika account tidak ditemukan, keluar dari program
	if account.pin == "" {
		fmt.Println("Invalid PIN")
		return
	}

	// Menampilkan menu utama
	for {
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("Selamat datang di ATM ABC, " + account.name + "!")
		fmt.Println("Silahkan Pilih menu transaksi yang anda inginkan : ")
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Tarik Tunai")
		fmt.Println("3. Setor Tunai")
		fmt.Println("4. Transfer")
		fmt.Println("5. Ubah PIN")
		fmt.Println("6. Keluar")
		fmt.Println("-----------------------------------------------------------")
		fmt.Print("Masukkan nomor pilihan menu transaksi : ")
		fmt.Scanln(&userInput)

		switch strings.TrimSpace(userInput) {
		case "1":
			// Menampilkan balance
			fmt.Println("Sisa Saldo anda: Rp", account.balance)

		case "2":
			// Mengambil uang
			var amount int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik : Rp")
			fmt.Scanln(&amount)

			if amount > account.balance {
				fmt.Println("Maaf, Saldo anda tidak cukup")
				continue
			}

			// Hanya mengizinkan penarikan uang jika nominalnya kelipatan 50000
			if amount%50000 != 0 {
				fmt.Println("Anda salah memasukan nominal kelipatan, Nominal uang yang ingin ditarik wajib kelipatan Rp50.000")
				continue
			}

			account.balance -= amount
			fmt.Println("Saldo anda: Rp", account.balance)

		case "3":
			// Menyetor uang
			var amount int
			fmt.Print("Masukkan jumlah uang yang ingin disetorkan : Rp")
			fmt.Scanln(&amount)

			account.balance += amount
			fmt.Println("Saldo anda: Rp", account.balance)

		case "4":
			// Mentransfer uang
			var recipientAccount Account
			var recipientPin string

			fmt.Print("Enter recipient's PIN: ")
			fmt.Scanln(&recipientPin)

			// Mencari account penerima yang cocok dengan PIN yang dimasukkan user
			for _, acc := range accounts {
				if acc.pin == recipientPin {
					recipientAccount = acc
					break
				}
			}

			// Jika account penerima tidak ditemukan, kembali ke menu utama
			if recipientAccount.pin == "" {
				fmt.Println("Rekening penerima tidak ditemukan !")
				continue
			}

			var amount int
			fmt.Print("Enter amount to transfer: ")
			fmt.Scanln(&amount)

			if amount > account.balance {
				fmt.Println("Insufficient balance")
				continue
			}

			account.balance -= amount
			recipientAccount.balance += amount
			fmt.Println("Transfer successful")
			fmt.Println("Your new balance is: ", account.balance)

		case "5":
			// Mengganti PIN
			var newPin string
			fmt.Print("Masukan Pin baru anda : ")
			fmt.Scanln(&newPin)

			if len(newPin) != PIN_LENGTH {
				fmt.Println("PIN harus sebanyak", PIN_LENGTH, " digit")
				continue
			}

			account.pin = newPin
			fmt.Println("Pin anda telah berhasil diubah")
			fmt.Println("Berikut ini adalah pin baru anda : ", newPin)

		case "6":
			// Keluar dari program
			fmt.Println("Terima kasih telah menggunakan layanan ATM ABC ")
			fmt.Println("Semoga harimu menyenangkan :) ")
			return

		}
	}
}
