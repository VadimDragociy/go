package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/VadimDragociy/go/client"
	"github.com/VadimDragociy/go/server"
)

func main() {
	str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))
	srv := server.NewServer(":8080")

	if err := srv.Start(); err != nil {
		fmt.Println(err)
		return
	}

	client := client.NewClient("http://localhost:8080")
	body, getVersion_err := client.GetVersion()

	if getVersion_err != nil {
		fmt.Println(getVersion_err)
		return
	}
	fmt.Println(string(body))

	decodedString, postDecode_err := client.PostDecode(str)

	if postDecode_err != nil {
		fmt.Println(postDecode_err)
		return
	}
	fmt.Println(decodedString)

	status, code, getHardOp_err := client.GetHardOp()

	if getHardOp_err != nil {
		fmt.Println(getHardOp_err)
		// return
	}
	if status {
		fmt.Printf("%t, %d\n", status, code)
		// return
	}
	fmt.Printf("%t\n", status)

	// srv.Shutdown(context.Background())

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// // <-quit
	// // <-quit
	// fmt.Println("Shutdown signal recieved")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %s\n", err)
	}
	fmt.Println("Server shutdown successfully")
}
