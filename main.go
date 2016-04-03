package main

import (
	"encoding/json"
	"flag"
	"log"
	//
	// This has been modified using magiconair/jose
	// as shown below
	//
	//"github.com/magiconair/jose/crypto"
	//"github.com/magiconair/jose/jws"
	"github.com/SermoDigital/jose/jwt"
)

/*
	Sample CLI app to test to token validity.
*/

func main() {
	var (
		err          error
		authToken    = flag.String("t", "", "JWT Token to validate")
		clientConfig = flag.String("c", "./keycloak-config.json", "OAuth2 client config")
	)
	flag.Parse()

	occ, err := LoadKeycloakClientConfig(*clientConfig)
	if err == nil {
		var kc *KeycloakClient
		if kc, err = NewKeycloakClient(occ); err == nil {
			var j jwt.JWT
			if j, err = kc.ValidateToken(*authToken); err == nil {
				c, _ := json.MarshalIndent(j.Claims(), " ", "  ")
				log.Printf("JWT Token: %s\n", c)
				return
			}
		}
	}

	log.Print(err)

}
