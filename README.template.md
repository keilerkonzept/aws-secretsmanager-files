# ${APP}

<!-- TOC -->

- [Get it](#get-it)
- [Use it](#use-it)
  - [Examples](#examples)
- [Comments](#comments)

<!-- /TOC -->

## Get it

Using go get:

```bash
go get -u github.com/sgreben/${APP}
```

Or [download the binary](https://github.com/sgreben/${APP}/releases/latest) from the releases page.

```bash
# Linux
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_linux_x86_64.zip
unzip ${APP}_${VERSION}_linux_x86_64.zip

# OS X
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_osx_x86_64.zip
unzip ${APP}_${VERSION}_osx_x86_64.zip

# Windows
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_windows_x86_64.zip
unzip ${APP}_${VERSION}_windows_x86_64.zip
```

## Use it

```text

${APP} [OPTIONS]

${USAGE}
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
