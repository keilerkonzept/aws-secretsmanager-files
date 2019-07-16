package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var config struct {
	SecretAssignments              AssignmentsMap
	SecretJSONKeyStringAssignments AssignmentsMap
	SecretJSONKeyAssignments       AssignmentsMap

	SecretJSONKeyStrings map[string]secretJSONKey
	SecretJSONKeys       map[string]secretJSONKey
	FileMode             uint
}

type secretJSONKey struct {
	SecretID string
	JSONKey  string
}

func init() {
	config.SecretJSONKeyStrings = make(map[string]secretJSONKey)
	config.SecretJSONKeys = make(map[string]secretJSONKey)
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Ldate)
	log.SetPrefix("[aws-secretsmanager-files] ")
	flag.Var(&config.SecretAssignments, "secret-file", "a key/value pair `FILE_PATH=SECRET_ARN` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyStringAssignments, "secret-json-key-string-file", "a key/value pair `FILE_PATH=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyAssignments, "secret-json-key-file", "a key/value pair `FILE_PATH=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.UintVar(&config.FileMode, "file-mode", 0400, "file mode for secret files")
	flag.Parse()

	for key, value := range config.SecretJSONKeyStringAssignments.Values {
		i := strings.IndexRune(value, '#')
		if i < 0 {
			log.Fatalf(`"%s" must have the form SECRET_ID#JSON_KEY`, value)
		}
		secretID, jsonKey := value[:i], value[i+1:]
		config.SecretJSONKeyStrings[key] = secretJSONKey{
			SecretID: secretID,
			JSONKey:  jsonKey,
		}
	}

	for key, value := range config.SecretJSONKeyAssignments.Values {
		i := strings.IndexRune(value, '#')
		if i < 0 {
			log.Fatalf(`"%s" must have the form SECRET_ID#JSON_KEY`, value)
		}
		secretID, jsonKey := value[:i], value[i+1:]
		config.SecretJSONKeys[key] = secretJSONKey{
			SecretID: secretID,
			JSONKey:  jsonKey,
		}
	}
}

func main() {
	awsSession, err := awsSession()
	if err != nil {
		log.Fatalf("aws: %v", err)
	}

	if err := awsSecretsFiles(secretsmanager.New(awsSession)); err != nil {
		log.Fatalf("error(s) while generating secret files: %v", err)
	}
}
