package cmd

import (
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/management-api/api"
)

func getManagementClient() api.ManagementClient {
	privateKeyPath := viper.GetString("key-path")
	if errValidate := common_tls.VerifyPrivateKeyPermissions(privateKeyPath); errValidate != nil {
		log.Printf("invalid private key permissions file: %s err: %s", privateKeyPath, errValidate)
	}

	certificate, err := common_tls.NewBundleFromFiles(viper.GetString("cert-path"), privateKeyPath, viper.GetString("ca-path"))
	if err != nil {
		log.Fatal(err)
	}

	creds := credentials.NewTLS(certificate.TLSConfig())

	c, err := grpc.Dial(viper.GetString("api-address"), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	return api.NewManagementClient(c)
}
