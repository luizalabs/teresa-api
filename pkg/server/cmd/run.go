package cmd

import (
	"crypto/rsa"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	"github.com/luizalabs/teresa-api/pkg/server"
	"github.com/luizalabs/teresa-api/pkg/server/database"
	"github.com/luizalabs/teresa-api/pkg/server/secrets"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  ``,
	Run:   runServer,
}

var port string

func init() {
	RunCmd.Flags().StringVarP(&port, "port", "p", "50051", "TCP port to create a listener")
}

func runServer(cmd *cobra.Command, args []string) {
	db, err := getDB()
	if err != nil {
		log.Fatal("Error on connect to Database: ", err)
	}

	privateKey, err := getPrivateKey()
	if err != nil {
		log.Fatal("Error on get signed key: ", err)
	}

	s, err := server.New(server.Options{
		Port:          port,
		RSAPrivateKey: privateKey,
		DB:            db})
	if err != nil {
		log.Fatal("Error on create Server: ", err)
	}

	log.Println("Start TERESA server on port ", port)
	log.Print(s.Run())
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	s, err := secrets.NewK8SSecrets()
	if err != nil {
		return nil, err
	}
	return s.PrivateKey()
}

func getDB() (*gorm.DB, error) {
	dbConf := new(database.Config)
	if err := envconfig.Process("teresadb", dbConf); err != nil {
		return nil, err
	}
	return database.New(*dbConf)
}
