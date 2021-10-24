// Handige links:
// https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/
// https://www.youtube.com/watch?v=GfKQEN65Vtg&t

package auth

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	AppClientId string
	UserPoolId  string
	*cip.Client // pointer naar cip.Client
}

// deze functie geeft een output van type CognitoClient
func Init() *CognitoClient {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	return &CognitoClient{
		AppClientId: os.Getenv("AWS_COGNITO_CLIENT_ID"),
		UserPoolId:  os.Getenv("AWS_USER_POOL_ID"),
		Client:      cip.NewFromConfig(cfg),
	}

}
