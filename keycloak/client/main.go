package main

import (
	"context"
	oidc "github.com/coreos/go-oidc"
)

var(
	clientID = "app"
	clientSecret = "7e36940b-641e-4c1e-99d0-dbf8105d5ce1"
)

func main(){
	ctx:= context.Background()

	provider, err := oidc.NewProvider(ctx, Issuer:"http://localhost:8080/auth/realms/demo")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.config{
		clientID: clientID,
		clientSecret: clientSecret,
		EndPoint: provider.EndPoint(),
		RedirectURL: "http://localhost:8081/auth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "magica"

	http.HandlerFunc( pattern: "/", func(writer http.ResponseWriter, request *http.Request){
		http.Redirect(writer, request, config.AuthCodeURL(state), http.statusFound)

	})

	http.HandlerFunc( pattern: "/auth/callback", func(w http.ResponseWriter, r *http.Request){
		if r.URL.Query().Get(key: "state") != state {
			http.Error(w, error: "State didi not match", http.statusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get(key: "code"))
		if err != nil {
			http.Error(w, error: "Failed to exchange token", http.statusBadRequest)
			return
		}

		rawIDToken, ok := oauth2Token.Extra(key: "id_token").(string)
		if !ok {
			http.Error(w, error: "no id_token", http.statusBadRequest)
			return	
		}

		resp :
			struct {
				OAuth2Token *oauth2.Token
				RawIDToken string
			}{
				OAuth2Token: oauth2Token, RawIDToken:rawIDToken,
			}

		data, err := json.MarshalIndent(resp, prefix:"", indent:"    ")	
		if err != nil {
			http.Error(w, err.Error(), http.statusBadRequest)
			return	
		}

		w.Write(data)

	})	

	log.Fatal(http.ListenAndServe(addr: "8081",handler: nill))

}