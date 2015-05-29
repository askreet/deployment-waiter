package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/awslabs/aws-sdk-go/service/codedeploy"
)

func main() {
	var deploymentId string
	flag.StringVar(&deploymentId, "d", "", "Deployment ID")
	flag.Parse()

	// Get the configuration from the environment.
	cd := codedeploy.New(nil)

	for {
		output, err := cd.GetDeployment(&codedeploy.GetDeploymentInput{
			DeploymentID: &deploymentId,
		})

		if err != nil {
			fmt.Println("Could not GetDeployment:", err.Error())
			os.Exit(1)
		}

		if output.DeploymentInfo == nil {
			fmt.Println("Response did not contain DeploymentInfo.")
			os.Exit(1)
		} else if output.DeploymentInfo.Status == nil {
			fmt.Println("Response did not contain a Status")
			os.Exit(1)
		}

		switch *output.DeploymentInfo.Status {
		case "Succeeded":
			fmt.Println("Deployment succeeded!")
			os.Exit(0)
		case "Failed":
			fmt.Println("Deployment failed!")
			os.Exit(1)
		default:
			fmt.Println("Deployment in state:", *output.DeploymentInfo.Status, "waiting...")
		}

		time.Sleep(2 * time.Second)
	}
}
