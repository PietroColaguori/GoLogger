package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Define regex expression for credentials
var patterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)password[:=\s]*(\S+)`),
	regexp.MustCompile(`(?i)pass[:=\s]*(\S+)`),
	regexp.MustCompile(`(?i)pwd[:=\s]*(\S+)`),
	regexp.MustCompile(`(?i)username[:=\s]*(\S+)`),
	regexp.MustCompile(`(?i)email[:=\s]*(\S+)`),
	regexp.MustCompile(`(?i)(\S+)@(\S+)\.(\S+)[=:]\s*(\S+)`),
	regexp.MustCompile(`(?i)admin[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)root[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)ftp[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)ssh[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)github_user[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)aws_access_key_id[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)aws_secret_access_key[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)docker_user[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)kubernetes_token[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)mysql_user[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)slack_token[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)zoom_user[:\s]*(\S+)`),
	regexp.MustCompile(`(?i)netflix_user[:\s]*(\S+)`),
}

// Define keywords to look for
var services = []string{
	// Social Media and Communication
	"facebook",
	"gmail",
	"email",
	"twitter",
	"linkedin",
	"instagram",
	"github",
	"slack",
	"zoom",
	"whatsapp",
	"signal",
	"telegram",
	"webex",

	// Cloud Storage and File Sharing
	"dropbox",
	"icloud",
	"google drive",
	"onedrive",
	"box",

	// Financial and E-commerce
	"paypal",
	"amazon",
	"shopify",
	"stripe",
	"square",
	"bank",

	// Development and Hosting Platforms
	"aws",
	"azure",
	"oracle",
	"salesforce",
	"wordpress",
	"vimeo",
	"heroku",
	"digitalocean",

	// Streaming and Entertainment
	"netflix",
	"spotify",
	"twitch",

	// Travel and Services
	"uber",
	"airbnb",
	"expedia",
	"etsy",

	// CLI and Technical Services
	"ssh",
	"ftp",
	"telnet",
	"sftp",
	"scp",
	"rlogin",
	"vnc",
	"docker",
	"kubernetes",

	// Miscellaneous
	"microsoft",
	"alibaba",
	"bing",
	"teams",
}

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func main() {
	filePath := "../attacker/capture.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		// fmt.Println("Current line:", line)

		// Check for patterns
		for _, pattern := range patterns {
			matches := pattern.FindAllStringSubmatch(line, -1)
			if len(matches) > 0 {
				for _, match := range matches {
					// fmt.Println("Match found: ", match)
					service := detectService(line)
					fmt.Printf("Potential leak for service %s%s%s found on line %d: %s%s%s\n", Red, service, Reset, lineNumber, Red, match[0], Reset)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// detectService identifies the service based on the line content
func detectService(line string) string {
	for _, service := range services {
		if strings.Contains(strings.ToLower(line), service) {
			return service
		}
	}
	return "unknown"
}
