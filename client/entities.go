package client

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateInstanceRequest struct {
	Status string `json:"status"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ReponseLogin struct {
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}
type Client struct {
	region     string
	provider   string
	httpClient *http.Client
}

type Environment struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

type AttachInstanceRequest struct {
	ProjectId     int    `json:"project_id"`
	Name          string `json:"name"`
	Instance_type string `json:"type"`
}

type Dns_zones struct {
	Zones []string `json:"zones"`
}

type InstancesTypes struct {
	Types []string `json:"types"`
}
type Instance struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Zone          string `json:"zone"`
	Root_dns_zone string `json:"root_dns_zone"`
	Instance_type string `json:"type"`
	Environment   string `json:"environment"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	Project       int    `json:"project_id"`
	Region        string `json:"region"`
	Ip_address    string `json:"ip_address"`
	Project_name  string `json:"project_name"`
	Project_url   string `json:"project_url"`
}

type Bucket struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	AccessKey string `json:"access_key"`
	Endpoint  string `json:"endpoint"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
	Type      string `json:"type"`
}
type Registry struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	AccessKey string `json:"access_key"`
	Endpoint  string `json:"endpoint"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
	Type      string `json:"type"`
}

type ApiKey struct {
	Accesskey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

type Project struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	CreatedAt string     `json:"created_at"`
	Instances []Instance `json:"instances"`
}

type AddProjectBody struct {
	Name        string `json:"name"`
	Host        string `json:"host"`
	Token       string `json:"token"`
	Namespace   string `json:"namespace"`
	GitUsername string `json:"git_username"`
}
type ProviderRegion struct {
	Name  string   `json:"name"`
	Zones []string `json:"zones"`
}
type ProviderRegions struct {
	Regions []ProviderRegion `json:"regions"`
}
type Provider struct {
	Name string `json:"name"`
}
type AvailableProviders struct {
	Providers []Provider `json:"providers"`
}
