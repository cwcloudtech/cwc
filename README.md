# COMWORK CLOUD CLI (CWC cli)

## Install
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
    