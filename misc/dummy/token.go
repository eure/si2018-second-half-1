package main

import (
	"fmt"
	"time"

	"github.com/eure/si2018-second-half-1/entities"
	"github.com/eure/si2018-second-half-1/repositories"
	"github.com/go-openapi/strfmt"
)

func dummyToken() {
	fmt.Printf("dummyToken")
	s := repositories.NewSession()
	r := repositories.NewUserTokenRepository(s)

	for i := firstUserID; i <= lastUserID; i++ {
		now := strfmt.DateTime(time.Now())
		token := entities.UserToken{
			UserID:    int64(i),
			Token:     fmt.Sprintf("USERTOKEN%v", i),
			CreatedAt: now,
			UpdatedAt: now,
		}
		r.Create(token)
		if i%100 == 0 {
			fmt.Println("dummyToken", i)
		}
	}
}
