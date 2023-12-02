package dtool

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
)

type Visitor struct{}

func (v Visitor) Visit(key string, val interface{}) error {
	fmt.Printf("%s: %s\n", key, val)
	return nil
}

func JwtInfo(ctx context.Context, b []byte, jwksEndpoint *string, aud, iss string) {
	// parse
	tok, err := jwt.Parse(b, jwt.WithVerify(false))
	if err != nil {
		log.Println(err)
	} else {
		v := Visitor{}
		tok.Walk(ctx, v)
	}

	// verify
	if jwksEndpoint == nil {
		log.Println("no endpoint")
	} else {
		keySet, err := jwk.Fetch(ctx, *jwksEndpoint)
		_, err = jws.Verify(b, jws.WithKeySet(keySet))
		if err != nil {
			log.Println(err)
		} else {
			log.Println("verify ok")
		}
	}

	// validate
	if tok == nil {
		log.Println("no token")
		return
	}
	err = jwt.Validate(tok, jwt.WithAudience(aud), jwt.WithIssuer(iss))
	if err != nil {
		log.Println(err)
	} else {
		log.Println("validate ok")
	}

}
