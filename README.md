# Go Cosmos

Project building structure cosmos

## Install package

```
go mod download
```

## Setup local node

- Initialize the validator node

```
mkdir -p node-main
archwayd init node-main --chain-id my-chain --home ./node-main
```

- Create a key to hold your account.

```
archwayd keys add node-main-account
```


