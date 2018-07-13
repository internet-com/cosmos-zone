# cosmos-zone
Create a new blockchain project based on cosmos-sdk with a single command.

---

# Installation

```shell
go get github.com/svaishnavy/cosmos-zone
cd $GOPATH/src/github.com/svaishnavy/cosmos-zone
make
```

This will install a binary cosmos-zone

# Creating a new cosmos-zone

**cosmos-zone init** _Your-Project-Name_

This will initialize a project, the dependencies, directory structures with the specified project name.

### Example:
```shell
cosmos-zone init testzone
```

This will ask for a remote path for the project - usually github.com/your_user_name/testzone and will create a new testzone application under $GOPATH/src/github.com/your_user_name/testzone along with Makefile

```shell
cd $GOPATH/src/github.com/your_user_name/testzone
make
```
This will create two binaries(testzonecli and testzoned) under bin folder. testzoned is the full node of the application which you can run, and testzonecli is your light client.

# TODO
1. Provide release binaries for cosmos-zone tool so it's easy for new users to get started.
2. Add --bare option, so only the top level app will be created without all the extensions.
