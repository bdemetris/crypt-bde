# Crypt-BDE

**WARNING:** As this has the potential for stopping users from logging in, extensive testing should take place before deploying into production.

Managing bitlocker on windows using golang, and then submit it to an instance of  [crypt-server](https://github.com/grahamgilbert/crypt-server)

## Features

* Rotate keys and escrow to your crypt server.

# Getting Started

## Building cryptbde

### Downloading the source.

```golang
go get github.com/bdemetris/crypt-bde
```

### Install dependent programs and libraries.

```shell
make deps
```

### Build the code.

```shell
make build
```

## Configuration

The configuration is defined in the config.json example is contained in the root of the Repo please update this with the URL of your crypt server.

## Running

Rotate keys - By default the rotate key command will rotate and then escrow the keys.

```shell
crypt-bde.exe --config=config.json rotatekey
```

Checkin - The checkin command will  escrow the current key without rotation, this will not enable bitlocker.

```shell
crypt-bde.exe --config=config.json checkin
```

### Dont read too much into this.  Just having fun.
