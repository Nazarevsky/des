package cli

import (
	"Lab1_des/des"
	"bufio"
	"fmt"
	"os"
)

func desTesting() {
	scanner := bufio.NewScanner(os.Stdin)
	var errText = ""

	for {
		clearConsole()
		if errText != "" {
			println(errText)
			errText = ""
		}

		println("Enter some text to encrypt:")
		scanner.Scan()
		text := []byte(scanner.Text())

		println("Enter key:")
		scanner.Scan()
		key := []byte(scanner.Text())

		encr, entrEncr, errEncr := des.Encrypt(text, key)

		if errEncr != nil {
			errText = fmt.Sprintf("Error: %s", errEncr.Error())
			continue
		}

		tryPrintEntropy("Encrypt entropy:", entrEncr)

		decr, entrDecr, _ := des.Decrypt(encr, key)
		tryPrintEntropy("Decrypt entropy:", entrDecr)

		fmt.Printf("Test results:\n")
		fmt.Printf("Is key weak: %v\n", des.IsKeyWeak(key))
		fmt.Printf("Encrypted bytes: %v\n", encr)
		fmt.Printf("Key bytes: %v\n", key)
		fmt.Printf("Decrypted text: %s\n", string(decr))

		if string(text) == string(decr) {
			println("Test passed")
		} else {
			println("Test is not passed")
		}

		println()

		println("Enter 0 to back to the menu or eny key to try one more time...")
		scanner.Scan()
		switch scanner.Text() {
		case "0":
			return
		}
	}
}
