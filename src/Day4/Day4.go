package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	countryId      int
	birthYear      int
	issueYear      int
	expirationYear int
	height         string
	hairColor      string
	eyeColor       string
	passportId     string
	parsingError   error
}

func main() {

	passports, err := scanPassports("./Day4Input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	validPassports := len(passports)

	for _, pport := range passports {
		// fmt.Print(pport.countryId, ";", pport.birthYear, ";", pport.issueYear, ";", pport.expirationYear,
		// 	";", pport.height, ";", pport.hairColor, ";", pport.eyeColor, ";", pport.passportId)
		fmt.Print("Height: ", pport.height, " Hair: ", pport.hairColor)

		invalidText := ""
		heightIsValid := validHeight(pport.height)
		if !heightIsValid {
			invalidText = "height"
		}

		validHairColor := validHairColor(pport.hairColor)
		if !validHairColor {
			invalidText += " hairColor"
		}

		validEyeColor := validEyeColor(pport.eyeColor)
		validPassportNr := validPassportNr(pport.passportId)
		validBirthYear := pport.birthYear >= 1920 && pport.birthYear <= 2002

		if !validBirthYear ||
			pport.expirationYear < 2020 || pport.expirationYear > 2030 ||
			pport.issueYear < 2010 || pport.issueYear > 2020 ||
			!heightIsValid ||
			!validHairColor ||
			!validEyeColor ||
			!validPassportNr {
			validPassports--
			fmt.Printf("\t\tUgyldig! %s\n", invalidText)
		} else {
			fmt.Println()
		}
	}

	fmt.Println("Gyldige pass: ", validPassports, " av ", len(passports))
}

func validPassportNr(passnr string) bool {
	if passnr == "" {
		return false
	}
	nr, err := strconv.Atoi(passnr)
	if err != nil {
		return false
	}
	if nr <= 999999999 {
		return true
	}
	return false
}

func validEyeColor(color string) bool {
	if color == "" {
		return false
	}
	if color == "amb" || color == "blu" || color == "brn" ||
		color == "gry" || color == "grn" ||
		color == "hzl" || color == "oth" {
		return true
	}
	return false
}

func validHeight(height string) bool {
	if height == "" {
		return false
	}
	lenHeight := len(height)
	heightUnit := height[lenHeight-2 : lenHeight]
	heightNumber, heightErr := strconv.Atoi(height[:lenHeight-2])
	if heightErr != nil {
		return false
	}
	if (heightUnit == "cm" && (heightNumber < 150 || heightNumber > 193)) ||
		(heightUnit == "in" && (heightNumber < 59 || heightNumber > 76)) {
		return false
	}
	return true
}

func validHairColor(color string) bool {
	if color == "" || len(color) != 7 {
		return false
	}
	if color[0:1] != "#" {
		return false
	}
	_, err := hex.DecodeString(color[1:7])
	if err != nil {
		return false
	}
	return true
}

func scanPassports(path string) ([]Passport, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	passports := []Passport{}
	passports = append(passports, Passport{})

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	passportIndex := 0

	for scanner.Scan() {
		var line = scanner.Text()

		if line == "" {
			// new passport
			passports = append(passports, Passport{})
			passportIndex++
			continue
		}

		keyValues := strings.Split(line, " ")
		for keyValueIndex := 0; keyValueIndex < len(keyValues); keyValueIndex++ {
			keyValue := strings.Split(keyValues[keyValueIndex], ":")
			passportKey := keyValue[0]
			passportValue := keyValue[1]
			passportValueInt, err := strconv.Atoi(passportValue)

			switch passportKey {
			case "byr":
				passports[passportIndex].birthYear = passportValueInt
				passports[passportIndex].parsingError = err
			case "cid":
				passports[passportIndex].countryId = passportValueInt
				passports[passportIndex].parsingError = err
			case "ecl":
				passports[passportIndex].eyeColor = passportValue
			case "eyr":
				passports[passportIndex].expirationYear = passportValueInt
				passports[passportIndex].parsingError = err
			case "hgt":
				passports[passportIndex].height = passportValue
			case "hcl":
				passports[passportIndex].hairColor = passportValue
			case "iyr":
				passports[passportIndex].issueYear = passportValueInt
				passports[passportIndex].parsingError = err
			case "pid":
				passports[passportIndex].passportId = passportValue
			default:
				return nil, errors.New("Passportkey " + passportKey + " finnes ikke!!")
			}
		}
	}

	return passports, nil
}
