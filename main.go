package main

import (
	"flag"
	"fmt"
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
	Profile              string
	PrintVersionAndExit  bool
}

var (
	app     = "aws-secretsmanager-files"
	version = "SNAPSHOT"
)

type secretJSONKey struct {
	SecretID string
	JSONKey  string
}

func init() {
	config.SecretJSONKeyStrings = make(map[string]secretJSONKey)
	config.SecretJSONKeys = make(map[string]secretJSONKey)
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Ldate)
	log.SetPrefix(fmt.Sprintf("[%s] ", app))
	flag.Var(&config.SecretAssignments, "secret", "a key/value pair `FILE_PATH=SECRET_ARN` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyStringAssignments, "secret-json-key-string", "a key/value pair `FILE_PATH=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.Var(&config.SecretJSONKeyAssignments, "secret-json-key", "a key/value pair `FILE_PATH=SECRET_ARN#JSON_KEY` (may be specified repeatedly)")
	flag.StringVar(&config.Profile, "profile", "", "override the current AWS_PROFILE setting")
	flag.UintVar(&config.FileMode, "file-mode", 0400, "file mode for secret files")
	flag.BoolVar(&config.PrintVersionAndExit, "version", false, "print version and exit")
	flag.Parse()

	if config.PrintVersionAndExit {
		fmt.Printf("%s %s", app, version)
		fmt.Println()
		os.Exit(0)
	}

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
