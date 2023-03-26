package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitlab.privy.id/order_service/internal/appctx"
	"gitlab.privy.id/order_service/internal/consts"
	"gitlab.privy.id/order_service/internal/entity"
)

type Auth struct {
	config *appctx.Config
}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			resp := appctx.NewResponse().WithMsgKey(consts.RespUnAuthorized).Generate()

			w.Header().Set("Content-Type", consts.HeaderContentTypeJSON)
			w.WriteHeader(resp.Code)
			w.Write(resp.Byte())
			return
		}

		user, err := userServices(username, password)
		if err != nil {
			fmt.Println(err)
		}
		if !user {
			resp := appctx.NewResponse().WithMsgKey(consts.RespUnAuthorized).Generate()

			w.Header().Set("Content-Type", consts.HeaderContentTypeJSON)
			w.WriteHeader(resp.Code)
			w.Write(resp.Byte())
			return
		}
		// call the next handler if authorized
		next.ServeHTTP(w, r)
	})
}

func userServices(username string, password string) (bool, error) {
	var (
		client = &http.Client{}
		cfg    = appctx.NewConfig()
		data   entity.UserServiceResponseEntity
		body   = entity.UserServiceRequestEntity{}
	)
	body.Username = username
	body.Password = password

	payload, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error parse json : ", err)
		return false, err
	}
	fmt.Println(cfg.App.UserServiceUrl)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/users/login", cfg.App.UserServiceUrl), bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Println("errror : ", err)
		return false, err
	}
	log.Println("Masuk lagi")
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("error parse json : ", err)
		return false, err
	}
	if data.Data {
		return true, nil
	} else {
		return false, nil
	}
}
