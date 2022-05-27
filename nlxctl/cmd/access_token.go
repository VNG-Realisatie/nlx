package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AuthMethod string

// String is used both by fmt.Print and by Cobra in help text
func (e *AuthMethod) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *AuthMethod) Set(v string) error {
	switch v {
	case string(basicAuth), string(oidc):
		*e = AuthMethod(v)
		return nil
	default:
		return errors.New("must be one of 'basic-auth' or 'oidc'")
	}
}

// Type is only used in help text
func (e *AuthMethod) Type() string {
	return "AuthMethod"
}

const (
	basicAuth AuthMethod = "basic-auth"
	oidc      AuthMethod = "oidc"

	accessToken string = "access-token"
)

var loginOptions struct {
	authMethod             AuthMethod
	authorizationServerURL string
	clientID               string
	clientSecret           string
	username               string
	password               string
}

//nolint:gochecknoinits // recommended way to use Cobra
func init() {
	rootCmd.AddCommand(loginCommand)

	loginCommand.Flags().Var(&loginOptions.authMethod, "auth-method", "authorization method ('basic-auth' or 'oidc')")
	loginCommand.Flags().StringVarP(&loginOptions.clientID, "client-id", "c", "", "client id")
	loginCommand.Flags().StringVarP(&loginOptions.clientSecret, "client-secret", "s", "", "client secret")
	loginCommand.Flags().StringVarP(&loginOptions.authorizationServerURL, "authorization-server-url", "a", "", "authorization server URL")
	loginCommand.Flags().StringVarP(&loginOptions.username, "username", "u", "", "username")
	loginCommand.Flags().StringVarP(&loginOptions.password, "password", "p", "", "password")

	err := loginCommand.MarkFlagRequired("auth-method")
	if err != nil {
		panic(err)
	}

	if loginOptions.authMethod == oidc {
		err = loginCommand.MarkFlagRequired("client-id")
		if err != nil {
			panic(err)
		}

		err = loginCommand.MarkFlagRequired("client-secret")
		if err != nil {
			panic(err)
		}

		err = loginCommand.MarkFlagRequired("authorization-server-url")
		if err != nil {
			panic(err)
		}
	}

	err = loginCommand.MarkFlagRequired("username")
	if err != nil {
		panic(err)
	}

	err = loginCommand.MarkFlagRequired("password")
	if err != nil {
		panic(err)
	}
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

//nolint:dupl // inway command looks like service command
var loginCommand = &cobra.Command{
	Use:   "login",
	Short: "login to oidc provider",
	Run: func(cmd *cobra.Command, arg []string) {
		data, err := os.ReadFile(configLocation)
		if err != nil {
			if !strings.Contains(err.Error(), "no such file or directory") {
				log.Fatal(err)
			}
		}

		var token string

		switch loginOptions.authMethod {
		case basicAuth:
			token = loginUsingBasicAuth()
		case oidc:
			token, err = loginUsingOIDC()
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatalf("invalid auth method: %s", loginOptions.authMethod)
		}

		err = viper.ReadConfig(bytes.NewBuffer(data))
		if err != nil {
			log.Fatal(err)
		}

		viper.Set(accessToken, token)

		err = viper.WriteConfigAs(configLocation)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("login successful")
	},
}

func loginUsingBasicAuth() string {
	return "Basic " + encodeBasicAuth(loginOptions.username, loginOptions.password)
}

func loginUsingOIDC() (string, error) {
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", loginOptions.username)
	form.Add("password", loginOptions.password)
	form.Add("scope", "openid email")

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token", loginOptions.authorizationServerURL), strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(loginOptions.clientID, loginOptions.clientSecret)

	client := http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to log in. status: %d, message: %s", response.StatusCode, body)
	}

	tokenResponse := &TokenResponse{}

	err = json.Unmarshal(body, tokenResponse)
	if err != nil {
		log.Panic(err)
	}

	return "Bearer " + tokenResponse.AccessToken, nil
}

func encodeBasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
