package main

import (
	"errors"
	"fmt"
	"os"
	"ts-rpc-demo/server"

	tsrpc "github.com/millevolte/ts-rpc"
)

func main() {
	path := "GeneratedCode"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var tsSource = tsrpc.GetTSSource(tsrpc.TSConfig{Url: "http://localhost:8080", TsApi: nil, Path: "."})
	err := os.WriteFile("GeneratedCode/generatedTypescript.ts", []byte(tsSource), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server.Server()
}
