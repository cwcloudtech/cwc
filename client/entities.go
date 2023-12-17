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

type Email struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Bcc       string `json:"bcc"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	Templated bool   `json:"templated"`
}

type EmailResponse struct {
	Status string `json:"status"`
}

type ModelsResponse struct {
	Models []string `json:"models"`
}

type Prompt struct {
	Model   string `json:"model"`
	Message string `json:"message"`
}

type PromptResponse struct {
	Status   string   `json:"status"`
	Response []string `json:"response"`
	Score    float64  `json:"score"`
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

type LanguagesResponse struct {
	Languages []string `json:"languages"`
}

type TriggerKindsResponse struct {
	TriggerKinds []string `json:"kinds"`
}

type FunctionContent struct {
	Code                          string   `json:"code"`
	Name                          string   `json:"name"`
	Language                      string   `json:"language"`
	Args                          []string `json:"args"`
	Regexp                        string   `json:"regexp"`
	Callback_url                  string   `json:"callback_url"`
	Callback_authorization_header string   `json:"callback_authorization_header"`
}

type Function struct {
	Id         string          `json:"id"`
	Owner_id   int             `json:"owner_id"`
	Is_public  bool            `json:"is_public"`
	Content    FunctionContent `json:"content"`
	Created_at string          `json:"created_at"`
	Updated_at string          `json:"updated_at"`
}

type FunctionDisplay struct {
	Id         string `json:"id"`
	Owner_id   int    `json:"owner_id"`
	Is_public  bool   `json:"is_public"`
	Name       string `json:"name"`
	Language   string `json:"language"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type FunctionsResponse struct {
	Status     string     `json:"status"`
	Code       int        `json:"code"`
	StartIndex int        `json:"start_index"`
	MaxIndex   int        `json:"max_index"`
	Results    []Function `json:"results"`
}

type AddFunctionBody struct {
	Is_public bool            `json:"is_public"`
	Content   FunctionContent `json:"content"`
}

type UpdateFunctionBody struct {
	Id        string          `json:"id"`
	Is_public bool            `json:"is_public"`
	Content   FunctionContent `json:"content"`
}

type FunctionCodeTemplate struct {
	Args     []string `json:"args"`
	Language string   `json:"language"`
}

type FunctionCodeTemplateResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Template string `json:"template"`
}

type Argument struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InvocationContent struct {
	Args        []Argument `json:"args"`
	State       string     `json:"state"`
	Result      string     `json:"result"`
	User_id     string     `json:"user_id"`
	Function_id string     `json:"function_id"`
}

type Invocation struct {
	Id         string            `json:"id"`
	Invoker_id int               `json:"invoker_id"`
	Updated_at string            `json:"updated_at"`
	Content    InvocationContent `json:"content"`
	Created_at string            `json:"created_at"`
}

type SyncronousInvocation struct {
	Status     string     `json:"status"`
	Code       int        `json:"code"`
	Invocation Invocation `json:"entity"`
}

type InvocationDisplay struct {
	Id          string `json:"id"`
	Invoker_id  int    `json:"invoker_id"`
	Function_id string `json:"function_id"`
	State       string `json:"state"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
	Result      string `json:"result"`
}

type InvocationsResponse struct {
	Status     string       `json:"status"`
	Code       int          `json:"code"`
	StartIndex int          `json:"start_index"`
	MaxIndex   int          `json:"max_index"`
	Results    []Invocation `json:"results"`
}

type InvocationAddContent struct {
	Args        []Argument `json:"args"`
	Function_id string     `json:"function_id"`
}

type AddInvocationBody struct {
	Content InvocationAddContent `json:"content"`
}

type InvocationUpdate struct {
	Id               string            `json:"id"`
	Invoker_id       int               `json:"invoker_id"`
	Content          InvocationContent `json:"content"`
	Invoker_username string            `json:"invoker_username"`
}

type TriggerContent struct {
	Args        []Argument `json:"args"`
	Name        string     `json:"name"`
	Cron_expr   string     `json:"cron_expr"`
	Function_id string     `json:"function_id"`
}

type Trigger struct {
	Id         string         `json:"id"`
	Kind       string         `json:"kind"`
	Owner_id   int            `json:"owner_id"`
	Content    TriggerContent `json:"content"`
	Created_at string         `json:"created_at"`
	Updated_at string         `json:"updated_at"`
}

type TriggerDisplay struct {
	Id          string `json:"id"`
	Kind        string `json:"kind"`
	Owner_id    int    `json:"owner_id"`
	Name        string `json:"name"`
	Cron_expr   string `json:"cron_expr"`
	Function_id string `json:"function_id"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type TriggersResponse struct {
	Status     string    `json:"status"`
	Code       int       `json:"code"`
	StartIndex int       `json:"start_index"`
	MaxIndex   int       `json:"max_index"`
	Results    []Trigger `json:"results"`
}

type AddTriggerBody struct {
	Kind    string         `json:"kind"`
	Content TriggerContent `json:"content"`
}
