# Crypt-BDE

**WARNING:** As this has the potential for stopping users from logging in, extensive testing should take place before deploying into production.

Managing bitlocker on windows using golang, and then submit it to an instance of  [crypt-server](https://github.com/grahamgilbert/crypt-server)

## Features

* Uses native authorization plugin so FileVault enforcement cannot be skipped.
* Escrow is delayed until there is an active user, so FileVault can be enforced when the Mac is offline.
* Administrators can specify a series of username that should not have to enable FileVault (IT admin, for example).

## Configuration

I don't know if these are all valid however good placeholders until tested.

### ServerURL

The `ServerURL` preference sets your Crypt Server. Crypt will not enforce FileVault if this preference isn't set.

``` bash
$ sudo defaults write /Library/Preferences/com.grahamgilbert.crypt ServerURL "https://crypt.example.com"
```

### SkipUsers

The `SkipUsers` preference allows you to define an array of users that will not be forced to enable FileVault.

``` bash
$ sudo defaults write /Library/Preferences/com.grahamgilbert.crypt SkipUsers -array-add adminuser
```

### RemovePlist

By default, the plist with the FileVault Key will be removed once it has been escrowed. In a future version of Crypt, there will be the possibility of verifying the escrowed key with the client. In preparation for this feature, you can now choose to leave the key on disk.

``` bash
$ sudo defaults write /Library/Preferences/com.grahamgilbert.crypt RemovePlist -bool FALSE
```

### RotateUsedKey

Crypt2 can rotate the recovery key, if the key is used to unlock the disk. There is a small caveat that this feature only works if the key is still present on the disk. This is set to `TRUE` by default.

## Building cryptbde

### Downloading the source.

```golang
go get github.com/bdemetris/crypt-bde
```

### Install dependent programs and libraries.

```shell
make deps
```

### Delete all build artifacts.

```shell
make clean
```

### Build the code.

```shell
make build
```

### Run the Go tests.

```shell
make test
```

### Run the Go linters.

```shell
make lint
```

### Dont read too much into this.  Just having fun.
