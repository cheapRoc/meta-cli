package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/authentication"
	"github.com/joyent/triton-go/compute"
)

var nameFlag string

type Config struct {
	keyID       string
	accountName string
	tritonURL   string
}

func envConfig() (*Config, error) {
	cfg := &Config{
		keyID:       os.Getenv("SDC_KEY_ID"),
		accountName: os.Getenv("SDC_ACCOUNT"),
		tritonURL:   os.Getenv("SDC_URL"),
	}

	// privateKey, err := ioutil.ReadFile(cfg.keyPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("can't find key file matching %s\n%s", cfg.keyID, err)
	// }
	// cfg.privateKey = privateKey
	return cfg, nil
}

func main() {
	flag.StringVar(&nameFlag, "name", "", "Glob name of instances to set metadata")
	flag.Parse()

	cfg, err := envConfig()
	if err != nil {
		log.Fatalf("can't configure triton using: %v+", cfg)
	}

	signer, err := authentication.NewSSHAgentSigner(cfg.keyID, cfg.accountName)
	if err != nil {
		log.Fatalf("can't configure NewSSHAgentSigner: %s", err)
	}

	// signer, err := authentication.NewPrivateKeySigner(cfg.keyID, cfg.privateKey, cfg.accountName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	api, err := compute.NewClient(&triton.ClientConfig{
		TritonURL:   cfg.tritonURL,
		AccountName: cfg.accountName,
		Signers:     []authentication.Signer{signer},
	})
	if err != nil {
		log.Fatalf("can't init client: %s", err)
	}

	// search for all instances by a given name in Triton

	insts, err := api.Instances().List(context.Background(), &compute.ListInstancesInput{
		Alias: nameFlag,
		Limit: 1,
	})
	if err != nil {
		log.Fatalf("can't find instances: %s")
	}

	for _, inst := range insts {
		fmt.Println("Instance: ", inst.Name)
	}

	// for G groups of instances, perform T goroutines for updating metadata

	// mdata, err := api.Instances().UpdateMetadata(context.Background(), &compute.UpdateMetadataInput{
	// 	ID: "2e021f5b-1aff-6f9e-c68b-fc33808f8355",
	// 	Metadata: map[string]string{
	// 		"something-two": "true",
	// 	},
	// })

	// fmt.Println("mdata: %v+", mdata)
}
