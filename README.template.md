# ${APP}

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/sgreben/${APP}.svg)](https://hub.docker.com/r/sgreben/${APP}/tags)

Writes AWS Secrets Manager secrets to files on disk.

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
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_osx_x86_64.tar.gz | tar xz

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
$ ${APP} -secret ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1
$ cat ./secret.json
{"hello":"world"}

$ ${APP} -secret-json-key ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
"world"

$ ${APP} secret-json-key-string ./secret.json=arn:aws:secretsmanager:eu-west-1:28381901202:secret:example-secret-1#hello
$ cat ./secret.json
world
```
