[![Go](https://github.com/project-illium/ilxd/actions/workflows/go.yml/badge.svg)](https://github.com/project-illium/ilxd/actions/workflows/go.yml)
[![golangci-lint](https://github.com/project-illium/ilxd/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/project-illium/ilxd/actions/workflows/golangci-lint.yml)

<h1 align="center">
<img src="https://raw.githubusercontent.com/project-illium/faucet/master/static/logo-white.png" alt="Illium logo" title="Illium logo">
</h1>

# ilxd
illium full node implementation written in Go

This is an alpha version of the illium full node software. This software does *not* have the proving system built in and 
is using mock proofs. 

The purpose is to validate all the rest of the code, networking, blockchain maintenance, transaction processing, etc, 
before we turn our attention to the proofs. 

If you want to test this alpha version you can download the binaries from the github releases page and run the node with
the `--alpha` flag.

```go
$ ilxd --alpha
```

### Install
Head over to the [releases](https://github.com/project-illium/ilxd/releases) page and download the lastest release for
your operating system. 

The release contains two binaries: `ilxd` and `ilxcli`. `ilxd` is the illium full node application and `ilxcli` is a 
command line application that is used to control and interact with a running node.

### Build From Source
Please note that the master branch is considered under active development and may contain bugs. If you are running in
a production environment please checkout a release tag. 

Currently `ilxd` will only build on go 1.20. This is a limmitation of some of the internal libp2p packages. When we update
libp2p we can update the version of go we use. 

```
$ git clone https://github.com/project-illium/ilxd.git
$ cd ilxd
$ make install
```
This command builds both `ilxd` and `ilxcli`. The binaries will be put in `$GOPATH/bin`.

To put this directory in your path add these lines to your `/etc/profile` (for a system-wide installation) or `$HOME/.profile`:

```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
```

### Usage
Vist [docs.illium.org](https://docs.illium.org/docs/node/running_a_node) for a comprehensive guide to running a node.

### Contributing
We'd love your help! See the [contributing guidlines](https://github.com/project-illium/ilxd/blob/master/CONTRIBUTING.md) before submitting your first PR.
