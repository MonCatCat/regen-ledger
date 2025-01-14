# Prerequisites

## Hardware

We recommend the following hardware specifications:

- 16G RAM
- 4vCPUs
- 500GB Disk space

## Image

We recommend using Ubuntu 18.04 or 20.04.

## Software

### Git

For more information, see [Git](https://git-scm.com).

### Make

For more information, see [GNU Make](https://www.gnu.org/software/make/).

### Go

For more information, see [Go](https://golang.org/).

## Initial Setup

:::tip NOTE
These commands are included in the [quickstart script](../get-started/run-a-full-node.md#quickstart).
:::

Update packages:

```bash
sudo apt update
```

Install tools:

```bash
sudo apt install git build-essential wget jq -y
```

Download Go:

```bash
wget https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
```

Verify data integrity:

```bash
sha256sum go1.17.2.linux-amd64.tar.gz
```

Verify SHA-256 hash:

```bash
f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676
```

Unpack Go download:

```bash
sudo tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz
```

Set up environment:

```bash
echo '
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.profile
```

Source profile file:

```bash
. ~/.profile
```

You're now ready to [run a full node](run-a-full-node.md).