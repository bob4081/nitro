package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("php-fpm", "--fpm-config", "/etc/php-fpm.d/www.conf.default")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
