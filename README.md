# COMWORK CLOUD CLI (CWC cli)

## Install
### Using Curl

    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.0.9/downloads/cwc_1.0.9_darwin_all.tar.gz" -o "cwc_cli.tar.gz"
    
    mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
    
    sudo ./cwc_cli/install.sh


### Using homebrew

    brew tap cwc/cwc https://gitlab.comwork.io/cwc/homebrew-cwc.git 

    brew install cwc

## Usage
### Authentification Command
    cwc login -u <email> -p <password>

### Configure default region Command
    cwc configure -region <default_region>

### Get instances Command
    cwc get --all

### Get instance by Id Command
    cwc get -id <instanceId>

### Create instance Command

    cwc create -name <project_name> -env <environement> -instance_type <size> -email<email_address>
    
## Update instance status Command
    cwc update -id <instanceId> -status <action>

## Delete instance Command
    cwc delete -id <instanceId>
    