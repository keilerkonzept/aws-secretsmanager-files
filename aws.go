package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func awsSession() (*session.Session, error) {
	ec2MetadataConfig := aws.NewConfig()
	ec2MetadataSession, err := session.NewSession(ec2MetadataConfig)
	if err != nil {
		return nil, err
	}

	ec2Metadata := ec2metadata.New(ec2MetadataSession)
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&credentials.SharedCredentialsProvider{},
			&ec2rolecreds.EC2RoleProvider{Client: ec2Metadata},
		},
	)
	return session.NewSession(aws.NewConfig().WithCredentials(creds))
}

func awsSecretsFiles(s *secretsmanager.SecretsManager) error {
	var errors []error

	for path, secretID := range config.SecretAssignments.Values {
		var value []byte
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		if result.SecretString != nil {
			value = []byte(*result.SecretString)
		}
		if result.SecretBinary != nil {
			value = result.SecretBinary
		}
		ioutil.WriteFile(path, value, os.FileMode(config.FileMode))
	}

	for path, secret := range config.SecretJSONKeyStrings {
		var value []byte
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secret.SecretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		var jsonObject map[string]interface{}
		switch {
		case result.SecretString != nil:
			if err := json.Unmarshal([]byte(*result.SecretString), &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = []byte(fmt.Sprint(jsonObject[secret.JSONKey]))
		case result.SecretString != nil:
			if err := json.Unmarshal(result.SecretBinary, &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value = []byte(fmt.Sprint(jsonObject[secret.JSONKey]))
		}
		ioutil.WriteFile(path, value, os.FileMode(config.FileMode))
	}

	for path, secret := range config.SecretJSONKeys {
		var value []byte
		result, err := s.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secret.SecretID),
		})
		if err != nil {
			errors = append(errors, err)
			continue
		}
		var jsonObject map[string]interface{}
		switch {
		case result.SecretString != nil:
			if err := json.Unmarshal([]byte(*result.SecretString), &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value, _ = json.Marshal(jsonObject[secret.JSONKey])
		case result.SecretString != nil:
			if err := json.Unmarshal(result.SecretBinary, &jsonObject); err != nil {
				errors = append(errors, err)
				continue
			}
			value, _ = json.Marshal(jsonObject[secret.JSONKey])
		}
		ioutil.WriteFile(path, value, os.FileMode(config.FileMode))
	}
	if len(errors) == 1 {
		return errors[0]
	}
	if len(errors) > 0 {
		return fmt.Errorf("%d error(s): [%q, ...]", len(errors), errors[0])
	}
	return nil
}
