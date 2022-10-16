package onedrive

import (
	"log"
	"os"
	"os/exec"
)

func OneDriver(patch string, url string) {
	cmd := exec.Command("wget", url+"&download=1", "-O", patch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
