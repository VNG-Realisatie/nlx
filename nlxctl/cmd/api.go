package cmd

import (
	"crypto/tls"
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/management-api/api"
)

func getManagementClient() api.ManagementClient {
	ca, err := orgtls.LoadRootCert(viper.GetString("ca-path"))
	if err != nil {
		log.Fatal(err)
	}

	keyPair, err := tls.LoadX509KeyPair(viper.GetString("cert-path"), viper.GetString("key-path"))
	if err != nil {
		log.Fatal(err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{keyPair},
		RootCAs:      ca,
	})

	c, err := grpc.Dial(viper.GetString("api-address"), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	return api.NewManagementClient(c)
}
