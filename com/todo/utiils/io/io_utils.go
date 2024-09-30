package io

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	TwelveHourFormat TimeInputFormat = iota
	TwentyFourHourFormat
	MilitaryFormat
)

func BoolInput(varName string) (*bool, error) {
	var bString string
	var b bool
	fmt.Println(varName, "(Enter 'Y' for Yes, 'N' for No.)?")
	_, err := fmt.Scanln(&bString)
	if err != nil {
		return nil, err
	}
	for bString != "Y" && bString != "N" {
		fmt.Println("Invalid input. Please enter a valid input, (Enter 'Y' for Yes, 'N' for No.)")
		_, err = fmt.Scanln(&bString)
		if err != nil {
			return nil, err
		}
	}
	b = bString == "Y"
	return &b, nil
}

func StringInput(varName string, optional bool) (*string, error) {
	return StringInputOfLength(varName, optional, ^0<<1)
}

func StringInputOfLength(varName string, optional bool, length int) (*string, error) {
	var input = ""
	var runes = make([]rune, 0)
	if optional {
		fmt.Println("Enter", varName, "(optional, Press Enter to skip):")
	} else {
		fmt.Println("Enter", varName)
	}
	reader := bufio.NewReader(os.Stdin)
	last := false
	for len(runes) < length {
		r, _, err := reader.ReadRune()
		if err != nil || r == '\n' {
			input = string(runes)
			return &input, err
		}
		if r == '\\' {
			last = true
		} else if r == 'e' {
			if last {
				runes = runes[:len(runes)-1]
				break
			}
		} else {
			last = false
		}
		runes = append(runes, r)
	}
	input = string(runes)
	return &input, nil
}

func IntegerInput[T IntegerInputType](varName string) (*T, error) {
	var input T
	fmt.Println("Enter", varName)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}

func ChoiceInput(varName string, choices []string) (*int, error) {
	var input = -1
	fmt.Println("Choose", varName)
	fmt.Print("Choices: [")
	for i := range choices {
		fmt.Print(i+1, ":", choices[i])
		if i < len(choices)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return &input, err
	}
	for input < 0 || input > len(choices) {
		_, err = fmt.Scanln(&input)
		if err != nil {
			return &input, err
		}
	}
	input--
	return &input, nil
}

func TimeInput(varName string, tif TimeInputFormat) (*time.Time, error) {
	var input string
	if tif == TwelveHourFormat {
		fmt.Println("Enter", varName, "in a twelve hour format (a valid time looks like, 11:59AM, 05:36PM .. etc")
	} else if tif == TwentyFourHourFormat {
		fmt.Println("Enter", varName, "in a twenty-four hour format (a valid time looks like, 11:59, 17:36 .. etc")
	} else if tif == MilitaryFormat {
		fmt.Println("Enter", varName, "in a military time format (a valid time looks like, 1159, 1736 .. etc")
	} else {
		return nil, errors.New("invalid time input format")
	}
	_, err := fmt.Scanln(&input)
	if err != nil {
		return nil, err
	}
	parsedTime, er := parseTime(input, tif)
	if er != nil {
		return nil, er
	}
	return parsedTime, nil
}

type IntegerInputType interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint
}

type TimeInputFormat int

func parseTime(input string, tif TimeInputFormat) (*time.Time, error) {
	var runes = []rune(input)
	var year = time.Now().Year()
	var month = time.Now().Month()
	var day = time.Now().Day()
	var parsedTime time.Time
	var err error = nil
	if tif == TwelveHourFormat {
		hour, er := strconv.Atoi(string(runes[0]) + string(runes[1]))
		if er != nil {
			err = er
		}
		minute, er := strconv.Atoi(string(runes[3]) + string(runes[4]))
		if er != nil {
			err = er
		}
		timeOfDay := string(runes[5]) + string(runes[6])
		if timeOfDay == "PM" {
			hour += 12
			if hour == 24 {
				hour = 0
			}
		}
		parsedTime = time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	} else if tif == TwentyFourHourFormat {
		hour, er := strconv.Atoi(string(runes[0]) + string(runes[1]))
		if er != nil {
			err = er
		}
		minute, er := strconv.Atoi(string(runes[3]) + string(runes[4]))
		if er != nil {
			err = er
		}
		parsedTime = time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	} else if tif == MilitaryFormat {
		hour, er := strconv.Atoi(string(runes[0]) + string(runes[1]))
		if er != nil {
			err = er
		}
		minute, er := strconv.Atoi(string(runes[2]) + string(runes[3]))
		if er != nil {
			err = er
		}
		parsedTime = time.Date(year, month, day, hour, minute, 0, 0, time.Local)
	} else {
		err = errors.New("unsupported time format")
	}
	return &parsedTime, err
}
