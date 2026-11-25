package io

import (
	"bufio"
	"os"
)

func TakeUserInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	if scanner.Scan(){
		return scanner.Text(), nil
	}

	return "", scanner.Err()
}
