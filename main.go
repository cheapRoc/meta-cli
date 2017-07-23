package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	triton "github.com/joyent/triton-go"
	"github.com/joyent/triton-go/authentication"
	"github.com/joyent/triton-go/compute"
)

var (
	hlName = color.New(color.FgBlue, color.Bold).SprintFunc()
	hlKey  = color.New(color.FgYellow).SprintFunc()
)

func main() {
	args, profile := NewFlagArgs(), NewProfile()

	signer, err := authentication.NewSSHAgentSigner(profile.keyID, profile.accountName)
	if err != nil {
		log.Fatalf("can't configure NewSSHAgentSigner: %s", err)
	}

	c, err := compute.NewClient(&triton.ClientConfig{
		TritonURL:   profile.tritonURL,
		AccountName: profile.accountName,
		Signers:     []authentication.Signer{signer},
	})
	if err != nil {
		log.Fatalf("can't init client: %s", err)
	}

	if err = args.ValidateName(); err != nil {
		log.Fatal(err)
	}

	insts, err := c.Instances().List(context.Background(), &compute.ListInstancesInput{
		Name: args.name,
	})
	if err != nil {
		log.Fatalf("can't find instances: %s", err)
	}

	if args.deleteAll {
		for _, inst := range insts {
			err := c.Instances().DeleteAllMetadata(context.Background(), &compute.DeleteAllMetadataInput{
				ID: inst.ID,
			})
			if err == nil {
				fmt.Printf("%s: Cleared all metadata\n", hlName(inst.Name))
			}
		}

		os.Exit(0)
		return
	}

	if args.delete {
		for _, inst := range insts {
			err := c.Instances().DeleteMetadata(context.Background(), &compute.DeleteMetadataInput{
				ID:  inst.ID,
				Key: args.key,
			})
			if err == nil {
				fmt.Printf("%s: Removed %s\n", hlName(inst.Name), hlKey(args.key))
			}
		}

		os.Exit(0)
		return
	}

	if len(args.metadata) > 0 {
		for _, inst := range insts {
			mdata, _ := c.Instances().UpdateMetadata(context.Background(), &compute.UpdateMetadataInput{
				ID:       inst.ID,
				Metadata: args.metadata,
			})
			for key, val := range mdata {
				fmt.Printf("%s: %s = %s\n", hlName(inst.Name), hlKey(key), val)
			}
			fmt.Println("")
		}

		os.Exit(0)
		return
	}

	if args.key != "" && args.value != "" {
		for _, inst := range insts {
			mdata, _ := c.Instances().UpdateMetadata(context.Background(), &compute.UpdateMetadataInput{
				ID: inst.ID,
				Metadata: map[string]string{
					args.key: args.value,
				},
			})
			for key, val := range mdata {
				fmt.Printf("%s: %s = %s\n", hlName(inst.Name), hlKey(key), val)
			}
			fmt.Println("")
		}

		os.Exit(0)
		return
	}

	if args.key != "" {
		for _, inst := range insts {
			body, _ := c.Instances().GetMetadata(context.Background(), &compute.GetMetadataInput{
				ID:  inst.ID,
				Key: args.key,
			})
			fmt.Printf("%s: %s = %s\n", hlName(inst.Name), hlKey(args.key), body)
		}

		os.Exit(0)
		return
	}

	for _, inst := range insts {
		mdata, _ := c.Instances().ListMetadata(context.Background(), &compute.ListMetadataInput{
			ID: inst.ID,
		})
		for key, val := range mdata {
			fmt.Printf("%s: %s = %s\n", hlName(inst.Name), hlKey(key), val)
		}
	}
}
