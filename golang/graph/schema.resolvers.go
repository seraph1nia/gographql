package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Seraph1nia/gographql/golang/auth"
	"github.com/Seraph1nia/gographql/golang/graph/generated"
	"github.com/Seraph1nia/gographql/golang/graph/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (r *mutationResolver) SignupUser(ctx context.Context, input model.UserInput) (string, error) {
	// definieer client
	cognitoClient := auth.Init()
	// even een printje om cognitoClient te gebruiken
	fmt.Println(cognitoClient.AppClientId)

	// build request
	awsReq := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(cognitoClient.AppClientId),
		Password: aws.String(input.Password),
		Username: aws.String(input.Username),
	}
	_, err := cognitoClient.SignUp(ctx, awsReq)

	return "missie geslaagd", err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
