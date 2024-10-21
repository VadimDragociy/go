package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/VadimDragociy/go/client"
	"github.com/VadimDragociy/go/server"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))
	srv := server.NewServer(":8080")

	if err := srv.Start(); err != nil {
		fmt.Println(err)
	}

	client := client.NewClient("http://localhost:8080")
	body, err := client.GetVersion()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	decodedString, err := client.PostDecode(str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodedString)

	status, code, err := client.GetHardOp()

	if err != nil {
		fmt.Println(err)
	}
	if status {
		fmt.Printf("%t, %d\n", status, code)
		return
	}
	fmt.Printf("%t\n", status)

	<-ctx.Done()

	fmt.Println("Shutdown signal recieved")

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
		return
	}
	fmt.Println("Server shutdown successfully")
}
