# COMWORK CLOUD CLI

👋 Welcome to cwc CLI!

<p align="center">
    <img src="./img/command-line.png"/> <br/>
</p>

This is a CLI written in go that will help you work with [comwork cloud API](./README.md).  
It's pretty easy to ship into your IaC pipelines.  
You'll find everything you need in [our documentation](https://doc.cloud.comwork.io/docs/tutorials/api/cli).  

## Installation

Choose the installation mode that suits your operating system and follow the instructions.

### Homebrew

First installation:

```shell
brew tap cwc/cwc https://gitlab.comwork.io/oss/cwc/homebrew-cwc.git 
brew install cwc
```

Upgrade:

```shell
brew update
brew upgrade cwc
```

### Curl

#### Linux

##### Linux x86 (64 bit)

```shell
version="1.9.0"
curl -L "https://gitlab.comwork.io/oss/cwc/cwc/-/releases/v${version}/downloads/cwc_${version}_linux_amd64.tar.gz" -o "cwc_cli.tar.gz"
mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
sudo ./cwc_cli/install.sh
```

Beware of checking if the version is available in the [releases](https://gitlab.comwork.io/oss/cwc/cwc/-/releases) because we only keep the 5 last builds.

##### Linux arm (64 bit)

```shell
version="1.9.0"
curl -L "https://gitlab.comwork.io/oss/cwc/cwc/-/releases/v${version}/downloads/cwc_${version}_linux_arm64.tar.gz" -o "cwc_cli.tar.gz" 
mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
sudo ./cwc_cli/install.sh
```

Beware of checking if the version is available in the [releases](https://gitlab.comwork.io/oss/cwc/cwc/-/releases) because we only keep the 5 last builds.

#### For MacOS

##### MacOS x86/arm (64 bit)

```shell
version="1.9.0"
curl -L "https://gitlab.comwork.io/oss/cwc/cwc/-/releases/v${version}/downloads/cwc_${version}_darwin_all.tar.gz" -o "cwc_cli.tar.gz"
mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli     
sudo ./cwc_cli/install.sh
```

Beware of checking if the version is available in the [releases](https://gitlab.comwork.io/oss/cwc/cwc/-/releases) because we only keep the 5 last builds.

#### For Windows

##### Windows x86 (64 bit)

```shell
curl -L "https://gitlab.comwork.io/oss/cwc/cwc/-/releases/v1.9.0/downloads/cwc_1.9.0_windows_amd64.zip" -o "cwc_cli.zip"
unzip cwc_cli.zip 
cd 
cwc.exe
```

Beware of checking if the version is available in the [releases](https://gitlab.comwork.io/oss/cwc/cwc/-/releases) because we only keep the 5 last builds.

##### Windows arm (64 bit)

```shell
curl -L "https://gitlab.comwork.io/oss/cwc/cwc/-/releases/v1.9.0/downloads/cwc_1.9.0_windows_arm64.zip" -o "cwc_cli.zip"
unzip cwc_cli.zip 
cd cwc_cli
cwc.exe
```

## Where you can find this repository ?

* Main repo: https://gitlab.comwork.io/oss/cwc/cwc.git
* Github mirror: https://github.com/comworkio/cwc.git
* Gitlab mirror: https://gitlab.com/ineumann/cwc.git

## Contributions

If you're interested in contributing to our project please check out [contributing section](./CONTRIBUTING.md).
