package main

import (
	"errors"
	"fmt"
	"net/smtp"

	//"net/smtp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func sendEmail(body string, from string) error {
	fmt.Println("Attempting to retrieve secret...")
	password, err := getSecret()
	if err != nil {
		return err
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", "trouys16@gmail.com", password, "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"trouys16@gmail.com"}
	msg := []byte("To: trouys16@gmail.com\r\n" +
		"Subject: Message from" + from + "\r\n" +
		"\r\n" +
		body + "\r\n")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "trouys16@gmail.com", to, msg)
	if err != nil {
		return err
	}
	return nil
}

// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html

func getSecret() (string, error) {
	fmt.Println("1")
	secretName := "smtppassword"
	region := "us-east-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				return "", errors.New("ecryption failure")

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				return "", errors.New("internal service error")

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				return "", errors.New("invalid parameter exception")

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				return "", errors.New("invalid request exception")

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				return "", errors.New("resource not found exception")
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return "", errors.New(err.Error())
		}
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
		return string(secretString[17:(len(secretString) - 2)]), nil
	} else {
		return "", errors.New("error decrypting secret")
	}
}
