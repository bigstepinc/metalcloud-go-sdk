package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	metalcloud "github.com/metalsoft-io/metal-cloud-sdk-go/v2"
)

func main() {
	user := os.Getenv("METALCLOUD_USER_EMAIL")
	apiKey := os.Getenv("METALCLOUD_API_KEY")
	endpoint := os.Getenv("METALCLOUD_ENDPOINT")

	if user == "" || apiKey == "" || endpoint == "" {
		log.Fatal("METALCLOUD_USER, METALCLOUD_API_KEY, METALCLOUD_ENDPOINT environment variables must be set")
	}

	client, err := metalcloud.GetMetalcloudClient(user, apiKey, endpoint, false)
	if err != nil {
		log.Fatal("Error initiating client: ", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf(`syntax: %s  cmd [args]
Commands:
	list
	deploy <infrastructureID>
	delete <infrastructureID>
		`, os.Args[0])
	}

	switch op := os.Args[1]; op {
	//list infrastructures

	case "list":

		infras, err := client.Infrastructures()
		if err != nil {
			log.Fatal("Error retrieving a list of infrastructures: ", err)
		}

		for _, infra := range *infras {
			log.Printf("\tInfrastructure: %s (%d) [%s]", infra.InfrastructureLabel, infra.InfrastructureID, infra.UserEmailOwner)
			_, err = client.InfrastructureGetByLabel(infra.InfrastructureLabel)

			instanceArrays, err := client.InstanceArrays(infra.InfrastructureID)
			if err != nil {
				log.Fatal("Error retrieving a list of instance arrays: ", err)
			}
			for _, ia := range *instanceArrays {

				log.Printf("\t\tInstanceArray: %s(%d)", ia.InstanceArrayLabel, ia.InstanceArrayID)

			}

		}

		break

	case "delete":

		if len(os.Args[2]) < 3 {
			log.Fatalf("syntax: %s delete <infrastructureID>\n", os.Args[0])
		}

		infrastructureID, err := strconv.Atoi(os.Args[2])
		//infrastructureID := os.Args[2]

		if err != nil {
			log.Fatalf("infrastructureID must be a number, it is %s %s\n", os.Args[2], err)
		}

		ret, err2 := client.InfrastructureGet(infrastructureID)
		if err2 != nil {
			log.Fatalf("Error retrieving infrastructure %d %s", infrastructureID, err2)
		}

		fmt.Printf("Deleting infrastructure %s (%d). Are you sure? Type \"yes\" to continue:", ret.InfrastructureLabel, infrastructureID)
		reader := bufio.NewReader(os.Stdin)
		yes, _ := reader.ReadString('\n')

		if yes == "yes\n" {
			err3 := client.InfrastructureDelete(infrastructureID)
			if err3 != nil {
				log.Fatal("Error deleting infrastructure ", err3)
			}
		} else {
			log.Printf("aborting\n")
		}

	case "deploy":

		if len(os.Args[2]) < 3 {
			log.Fatalf("syntax: %s deploy <infrastructureID>\n", os.Args[0])
		}

		infrastructureID, err := strconv.Atoi(os.Args[2])

		if err != nil {
			log.Fatalf("infrastructureID must be a number, it is %s %s\n", os.Args[2], err)
		}

		ret, err2 := client.InfrastructureGet(infrastructureID)
		if err2 != nil {
			log.Fatalf("Error retrieving infrastructure %d %s", infrastructureID, err2)
		}

		shutDownOptions := metalcloud.ShutdownOptions{
			HardShutdownAfterTimeout:   true,
			AttemptSoftShutdown:        true,
			SoftShutdownTimeoutSeconds: 180,
		}

		fmt.Printf("Deploying infrastructure %s (%d). Are you sure? Type \"yes\" to continue:", ret.InfrastructureLabel, infrastructureID)
		reader := bufio.NewReader(os.Stdin)
		yes, _ := reader.ReadString('\n')

		if yes == "yes\n" {
			err3 := client.InfrastructureDeploy(ret.InfrastructureID, shutDownOptions, false, false)
			if err3 != nil {
				log.Fatal(err3)
			}
		} else {
			log.Printf("aborting\n")
		}

	}

}
