package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const inputPrompt = "To %s migrations choose %s"
const dockerMigrationCmd = "docker exec -i ddd.postgres psql -U postgres go-ddd < ./%s"

func main() {
	pattern := "./*." + getAction() + ".sql"

	matches, err := filepath.Glob(pattern)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range matches {
		runInDocker(file)
	}
}

func getAction() (action string) {
	for {
		fmt.Println(fmt.Sprintf(inputPrompt, "up", "1"))
		fmt.Println(fmt.Sprintf(inputPrompt, "down", "2"))

		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		i := input.Text()
		if i == "1" {
			return "up"
		}
		if i == "2" {
			return "down"
		}
	}
}

func runInDocker(migration string) {
	cmd := fmt.Sprintf(dockerMigrationCmd, migration)
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
