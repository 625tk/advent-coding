package main

import (
	"context"
	"fmt"
	"github.com/625tk/advent-coding/dtool"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

func main() {
	endpoint := getEnv("JWKS_ENDPOINT", "https://www.googleapis.com/oauth2/v3/certs")
	aud := getEnv("JWT_AUD", "my_aud")
	iss := getEnv("JWT_ISS", "my_iss")
	if terminal.IsTerminal(0) {
		fmt.Print("input jwt >>>")
	}

	ctx := context.Background()
	b, _ := io.ReadAll(os.Stdin)

	dtool.JwtInfo(ctx, b, &endpoint, aud, iss)
}

func getEnv(k, v string) string {
	if x := os.Getenv(k); x != "" {
		return x
	}
	return v
}
