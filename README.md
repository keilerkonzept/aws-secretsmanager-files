# aws-secretsmanager-files

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
  - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/sgreben/aws-secretsmanager-files
```

Or [download the binary](https://github.com/sgreben/aws-secretsmanager-files/releases/latest) from the releases page.

```bash
# Linux
curl -LO https://github.com/sgreben/aws-secretsmanager-files/releases/download/1.0.0/aws-secretsmanager-files_1.0.0_linux_x86_64.zip
unzip aws-secretsmanager-files_1.0.0_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/aws-secretsmanager-files/releases/download/1.0.0/aws-secretsmanager-files_1.0.0_osx_x86_64.zip
unzip aws-secretsmanager-files_1.0.0_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/aws-secretsmanager-files/releases/download/1.0.0/aws-secretsmanager-files_1.0.0_windows_x86_64.zip
unzip aws-secretsmanager-files_1.0.0_windows_x86_64.zip
```

## Use it

```text

aws-secretsmanager-files [OPTIONS]

Usage of aws-secretsmanager-files:
  -file-mode uint
    	file mode for secret files (default 256)
  -secret-file FILE_PATH=SECRET_ARN
    	a key/value pair FILE_PATH=SECRET_ARN (may be specified repeatedly)
  -secret-json-key-file FILE_PATH=SECRET_ARN#JSON_KEY
    	a key/value pair FILE_PATH=SECRET_ARN#JSON_KEY (may be specified repeatedly)
  -secret-json-key-string-file FILE_PATH=SECRET_ARN#JSON_KEY
    	a key/value pair FILE_PATH=SECRET_ARN#JSON_KEY (may be specified repeatedly)
```

### Examples

```shell
$ aws-secretsmanager-file -secret-file ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1
$ cat ./secret.json
{"hello":"world"}

$ aws-secretsmanager-file -secret-json-key-file ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
"world"

$ aws-secretsmanager-file secret-json-key-string-file ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
world
```
