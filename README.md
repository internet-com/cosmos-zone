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

Example:
cosmos-zone init democoin
This will ask for a remote path for the project - usually github.com/your_user_name/democoin
Will create a new democoin application under $GOPATH/src/github.com/your_user_name/democoin

# TODO
1. Provide release build for cosmos-zone tool so new users need not run make.
