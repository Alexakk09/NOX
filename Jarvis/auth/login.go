package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetCredentials() (string, string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Username > ")
	if !scanner.Scan() {
		return "", "", scanner.Err()
	}

	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Password > ")
	if !scanner.Scan() {
		return "", "", scanner.Err()
	}

	password := strings.TrimSpace(scanner.Text())

	return username, password, nil
}