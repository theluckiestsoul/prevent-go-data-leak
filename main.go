package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type MyData struct {
	Field     string
	sensitive string
}

func (MyData) Format(f fmt.State, verb rune) {
	f.Write([]byte("x"))
}

func (MyData) String() string {
	return "x"
}

func (MyData) LogValue() slog.Value {
	return slog.StringValue("x")
}

func (MyData) MarshalText() ([]byte, error) {
	return []byte("x"), nil
}

func main() {
	d := MyData{Field: "Hello world!", sensitive: "sensitive data"}

	fmt.Println(d)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("message", "data", d)

	b, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
