# aws-secretsmanager-files

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sgreben/aws-secretsmanager-files.svg)](https://hub.docker.com/r/sgreben/aws-secretsmanager-files/tags)

Writes AWS Secrets Manager secrets to files on disk. (If you need secrets as environment variables instead, you can use [aws-secretsmanager-env](https://github.com/keilerkonzept/aws-secretsmanager-env))

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
  - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/keilerkonzept/aws-secretsmanager-files
```

Or [download the binary](https://github.com/keilerkonzept/aws-secretsmanager-files/releases/latest) from the releases page.

```bash
# Linux
curl -L https://github.com/keilerkonzept/aws-secretsmanager-files/releases/download/1.2.225/aws-secretsmanager-files_1.2.225_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/keilerkonzept/aws-secretsmanager-files/releases/download/1.2.225/aws-secretsmanager-files_1.2.225_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/keilerkonzept/aws-secretsmanager-files/releases/download/1.2.225/aws-secretsmanager-files_1.2.225_windows_x86_64.zip
unzip aws-secretsmanager-files_1.2.225_windows_x86_64.zip
```

## Use it

```text

aws-secretsmanager-files [OPTIONS]

Usage of aws-secretsmanager-files:
  -file-mode uint
    	file mode for secret files (default 256)
  -profile string
    	override the current AWS_PROFILE setting
  -secret FILE_PATH=SECRET_ARN
    	a key/value pair FILE_PATH=SECRET_ARN (may be specified repeatedly)
  -secret-json-key FILE_PATH=SECRET_ARN#JSON_KEY
    	a key/value pair FILE_PATH=SECRET_ARN#JSON_KEY (may be specified repeatedly)
  -secret-json-key-string FILE_PATH=SECRET_ARN#JSON_KEY
    	a key/value pair FILE_PATH=SECRET_ARN#JSON_KEY (may be specified repeatedly)
  -version
    	print version and exit
```

### Examples

```shell
$ aws-secretsmanager-files -secret ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1
$ cat ./secret.json
{"hello":"world"}

$ aws-secretsmanager-files -secret-json-key ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
"world"

$ aws-secretsmanager-files secret-json-key-string ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
world
```
