package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Email struct {
	MessageID string
	Date      string
	From      string
	To        string
	Subject   string
	Body      strings.Builder
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./indexer [path_to_email_database]")
		os.Exit(1)
	}
	root := os.Args[1]

	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking through files: %v\n", err)
		os.Exit(1)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error visiting path %s: %v\n", path, err)
		return nil
	}

	if f == nil || f.IsDir() {
		return nil
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", path, err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024 //Evitar erroes por exceder el buffer del scanner
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	email := Email{}
	headersEnded := false
	bodyStarted := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && headersEnded {
			bodyStarted = true
			continue
		}

		if !bodyStarted && !strings.HasPrefix(line, "X-") {
			headersEnded = true
		}

		if !bodyStarted {
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				assignEmailField(&email, parts[0], parts[1])
			}
		} else {
			email.Body.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", path, err)
		return nil // Continúa con el siguiente archivo
	}

	if err := indexEmail(email); err != nil {
		fmt.Fprintf(os.Stderr, "Error indexing email: %v\n", err)
		// Decidir si continuar o no dependiendo de la criticidad del error
	}

	return nil
}

func assignEmailField(email *Email, key, value string) {
	switch key {
	case "Message-ID":
		email.MessageID = value
	case "Date":
		email.Date = value
	case "From":
		email.From = value
	case "To":
		email.To = value
	case "Subject":
		email.Subject = value
	}
}

func indexEmail(email Email) error {
	jsonData, err := json.Marshal(email)
	if err != nil {
		return fmt.Errorf("error marshalling email: %v", err)
	}

	url := "http://localhost:4080/api/test_emails/_doc"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request to ZincSearch: %v", err)
	}
	resp.Body.Close()

	return nil
}
