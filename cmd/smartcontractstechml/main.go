// cmd/smartcontractstechml/main.go
package main

import (
"flag"
"log"
"os"

"smartcontractstechml/internal/smartcontractstechml"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := smartcontractstechml.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
