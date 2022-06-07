# COMWORK CLOUD CLI (CWC cli)

## Installation
### Using Curl
#### For Linux
##### Linux x86 (64 bit)
    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.1.4/downloads/cwc_1.1.4_linux_amd64.tar.gz" -o "cwc_cli.tar.gz"
    
    mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
    
    sudo ./cwc_cli/install.sh

##### Linux arm (64 bit)
    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.1.4/downloads/cwc_1.1.4_linux_arm64.tar.gz" -o "cwc_cli.tar.gz"
    
    mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
    
    sudo ./cwc_cli/install.sh

#### For MacOS
##### MacOS x86/arm (64 bit)

    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.1.4/downloads/cwc_1.1.4_darwin_all.tar.gz" -o "cwc_cli.tar.gz"
    
    mkdir cwc_cli && tar -xf cwc_cli.tar.gz -C cwc_cli 
    
    sudo ./cwc_cli/install.sh


#### For Windows
##### Windows x86 (64 bit)

    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.1.4/downloads/cwc_1.1.4_windows_amd64.zip" -o "cwc_cli.zip"

    unzip cwc_cli.zip 
    cd 
    cwc.exe
##### Windows arm (64 bit)

    curl -L "https://gitlab.comwork.io/cwc/cwc/-/releases/v1.1.4/downloads/cwc_1.1.4_windows_arm64.zip" -o "cwc_cli.zip"

    unzip cwc_cli.zip 
    cd cwc_cli
    cwc.exe

### Using homebrew

    brew tap cwc/cwc https://gitlab.comwork.io/cwc/homebrew-cwc.git 

    brew install cwc

## Usage
### Authentification Command
    cwc login -u <email> -p <password>

### Configure default region Command
    cwc configure -region <default_region>

### Get instances Command
    cwc get instance --all

### Get instance by Id Command
    cwc get instance -id <instanceId>

### Create instance Command

    cwc create instance -name <project_name> -env <environement> -instance_type <size> -project_id <project-id>
    
### Attach instance Command

    cwc attach instance -name <playbook-name> -instance_type <size> -project_id <project-id>

### Update instance status Command
    cwc update instance -id <instanceId> -status <action>

### Delete instance Command
    cwc delete instance -id <instanceId>
    

### Get projects Command
    cwc get project --all

### Get project by Id Command
    cwc get project -id <instanceId>

### Create project Command

    cwc create project -name <project_name>

### Delete project Command
    cwc delete -id <projectId>

### Get environments Command
    cwc get environment --all

### Get environment by Id Command
    cwc get environment -id <environmentId>