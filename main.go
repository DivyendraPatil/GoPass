package main

import (
	"bufio"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"net/http"
	"os"
	"strings"
	"sync"
	"unicode"
)

const haveIBeenPwnedAPIURL = "https://api.pwnedpasswords.com/range/"

func main() {
	// Parse command-line arguments
	length := flag.Int("length", 12, "password length")
	uppercase := flag.Bool("uppercase", true, "include uppercase letters")
	lowercase := flag.Bool("lowercase", true, "include lowercase letters")
	numbers := flag.Bool("numbers", true, "include numbers")
	specials := flag.Bool("specials", true, "include special characters")
	numPasswords := flag.Int("num-passwords", 1, "number of passwords to generate")
	saveToFile := flag.Bool("save-to-file", false, "save passwords to a file")
	checkBreaches := flag.Bool("check-breaches", false, "check if password has been breached before")
	flag.Parse()

	// Determine the character set to use for password generation
	charset := ""
	if *uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if *lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if *numbers {
		charset += "0123456789"
	}
	if *specials {
		charset += "!@#$%^&*()-_=+[]{}\\|;:'\",.<>/?"
	}

	// Generate the specified number of passwords in parallel
	var wg sync.WaitGroup
	for i := 0; i < *numPasswords; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			password, strength := generatePassword(*length, charset)
			if *checkBreaches {
				if isBreached(password) {
					fmt.Printf("WARNING: password '%s' has been breached before\n", password)
				}
			}
			fmt.Printf("Password: %s\nStrength: %s\n", password, strength)
			if *saveToFile {
				savePasswordToFile(password)
			}
		}()
	}
	wg.Wait()
}

func generatePassword(length int, charset string) (string, int) {
	// Generate a random password with the specified length and character set
	var password strings.Builder
	var strength int
	types := 0
	prev := ""
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		char := string(charset[index.Int64()])
		password.WriteString(char)
		if unicode.IsUpper(rune(char[0])) {
			types |= 1
		} else if unicode.IsLower(rune(char[0])) {
			types |= 2
		} else if unicode.IsDigit(rune(char[0])) {
			types |= 4
		} else {
			types |= 8
		}
		if char == prev {
			strength--
		}
		prev = char
	}
	// Calculate password strength based on number of characters, types of characters, and repetition
	strength += int(math.Log2(float64(length))) * 4
	strength += bits.OnesCount(uint(types)) * 2
	if length > 8 {
		strength += bits.OnesCount(uint(types)) * 2
	}
	if length > 12 {
		strength += bits.OnesCount(uint(types)) * 2
	}
	if length > 16 {
		strength += bits.OnesCount(uint(types)) * 2
	}
	if length > 20 {
		strength += bits.OnesCount(uint(types)) * 2
	}
	if length > 24 {
		strength += bits.OnesCount(uint(types)) * 2
	}
	return password.String(), strength
}

func savePasswordToFile(password string) {
	// Save the password to a file named "passwords.txt" in the current directory
	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error saving password:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	if _, err := file.WriteString(password + "\n"); err != nil {
		fmt.Println("Error saving password:", err)
	}
}

func isBreached(password string) bool {
	// Check if the password has been breached before using the HaveIBeenPwned API
	hash := sha1.Sum([]byte(password))
	hashString := strings.ToUpper(hex.EncodeToString(hash[:]))

	prefix := hashString[0:5]
	suffix := hashString[5:]

	client := http.DefaultClient
	response, err := client.Get(fmt.Sprintf("%s%s", haveIBeenPwnedAPIURL, prefix))
	if err != nil {
		fmt.Println("Error checking if password has been breached:", err)
		return false
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if parts[0] == suffix {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error checking if password has been breached:", err)
	}

	return false
}
