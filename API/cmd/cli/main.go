package main

import (
	"bufio"
	
	"fmt"
	
	"os"
	"strings"
)




func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║            VeerAI            ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Println("Commands:")
	fmt.Println("  /pdf <path>  Upload and extract a PDF")
	fmt.Println("  exit         Quit")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Username > ")
	if !scanner.Scan() {
		fmt.Println("Could not read username.")
		return
	}
	username := strings.TrimSpace(scanner.Text())

	fmt.Print("Password > ")
	if !scanner.Scan() {
		fmt.Println("Could not read password.")
		return
	}
	password := strings.TrimSpace(scanner.Text())

	token, err := login(username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	conversationID, err := createConversation(token)
	if err != nil {
		fmt.Println("Failed to create conversation:", err)
		return
	}

	fmt.Println("Conversation ID:", conversationID)
	fmt.Println()

	for {
		fmt.Print("You > ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			fmt.Println("Goodbye!")
			return
		}

		if input == "/pdf" || input == `/pdf ""` {
			fmt.Println("Error: PDF path is required.")
			fmt.Println("Usage: /pdf <path>")
			fmt.Println()
			continue
		}

		if input == "/pdf explain" || input == `/pdf explain ""` {
			fmt.Println("Error: PDF path is required.")
			fmt.Println("Usage: /pdf explain <path>")
			fmt.Println()
			continue
		}

		if strings.HasPrefix(input, "/pdf summarize ") {
			filePath := strings.TrimSpace(strings.TrimPrefix(input, "/pdf summarize "))
			filePath = strings.Trim(filePath, `"`)

			err := summarizePDF(filePath)
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println()
			continue
		}

		if strings.HasPrefix(input, "/pdf explain ") {
			filePath := strings.TrimSpace(strings.TrimPrefix(input, "/pdf explain "))
			filePath = strings.Trim(filePath, `"`)

			err := explainPDF(filePath)
			if err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println()
			continue
		}

		sendChat(conversationID,token, input)
		fmt.Println()
	}
}

