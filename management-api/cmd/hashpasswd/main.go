package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alexedwards/argon2id"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		// CreateHash returns a Argon2id hash of a plain-text password using the
		// provided algorithm parameters. The returned hash follows the format used
		// by the Argon2 reference C implementation and looks like this:
		// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG
		hash, err := argon2id.CreateHash(text, argon2id.DefaultParams)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(hash)
	}
}
