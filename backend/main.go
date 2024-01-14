package main

import (
	"github.com/samoei/acromate/backend/api"
)

func main() {
	e := api.NewServer()
	e.Start(":42069")
}
