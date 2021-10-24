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
	// TODO: add email to schema and confirmationflow, use password hashing

	// definieer client
	cognitoClient := auth.Init()

	// build request
	awsReq := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(cognitoClient.AppClientId),
		Password: aws.String(input.Password),
		Username: aws.String(input.Username),
	}
	// post request
	_, err := cognitoClient.SignUp(ctx, awsReq)

	// build request
	confirmInput := &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(cognitoClient.UserPoolId),
		Username:   aws.String(input.Username),
	}

	// auto confirm all users for now
	_, err = cognitoClient.AdminConfirmSignUp(ctx, confirmInput)

	return "you are registered: " + input.Username, err
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.UserInput) (*model.SignInResponse, error) {
	// definieer client
	cognitoClient := auth.Init()

	// build request
	awsReq := &cognitoidentityprovider.AdminInitiateAuthInput{
		ClientId:       aws.String(cognitoClient.AppClientId),
		AuthFlow:       "ADMIN_USER_PASSWORD_AUTH",
		UserPoolId:     aws.String(cognitoClient.UserPoolId),
		AuthParameters: map[string]string{"USERNAME": input.Username, "PASSWORD": input.Password},
	}
	fmt.Println(awsReq)
	// AdminInitiateAuthOutput komt in output, hierin zit onder andere AuthenticationResultType in. Hier zitten de tokens in: https://docs.aws.amazon.com/sdk-for-go/api/service/cognitoidentityprovider/#AuthenticationResultType
	// hier wordt := gebruik want het is een nieuwe variabele EN er komt een waarde in. (declaration en assignment)
	output, err := cognitoClient.AdminInitiateAuth(ctx, awsReq)

	response := &model.SignInResponse{
		AccessToken:  *output.AuthenticationResult.AccessToken,
		ExpiresIn:    int(output.AuthenticationResult.ExpiresIn),
		IDToken:      *output.AuthenticationResult.IdToken,
		RefreshToken: *output.AuthenticationResult.RefreshToken,
		TokenType:    *output.AuthenticationResult.TokenType,
	}

	return response, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
