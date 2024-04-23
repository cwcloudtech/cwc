package admin

import "net/http"

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
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Path                string `json:"path"`
	Roles               string `json:"roles"`
	IsPrivate           bool   `json:"is_private"`
	Description         string `json:"description"`
	EnvironmentTemplate string `json:"environment_template"`
	DocTemplate         string `json:"doc_template"`
	SubDomains          string `json:"subdomains"`
	LogUrl              string `json:"logo_url"`
}

type ResponseUsers struct {
	Result []User `json:"result"`
}

type ResponseUser struct {
	Result User `json:"result"`
}

type User struct {
	Id                 int    `json:"id"`
	Email              string `json:"email"`
	RegistrationNumber string `json:"registration_number"`
	Address            string `json:"address"`
	CompanyName        string `json:"company_name"`
	ContactInfo        string `json:"contact_info"`
	IsAdmin            bool   `json:"is_admin"`
	Confirmed          bool   `json:"confirmed"`
	Billable           bool   `json:"billable"`
}

type AddEnvironmentType struct {
	Name                string `json:"name"`
	Path                string `json:"path"`
	Roles               string `json:"roles"`
	IsPrivate           bool   `json:"is_private"`
	Description         string `json:"description"`
	EnvironmentTemplate string `json:"environment_template"`
	DocTemplate         string `json:"doc_template"`
	SubDomains          string `json:"subdomains"`
	LogoUrl             string `json:"logo_url"`
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
	Email         string `json:"email"`
	Project       int    `json:"project_id"`
	Region        string `json:"region"`
	Ip_address    string `json:"ip_address"`
	Project_name  string `json:"project_name"`
	Project_url   string `json:"project_url"`
}

type Email struct {
	From      *string `json:"from,omitempty"`
	To        string `json:"to"`
	Bcc       *string `json:"bcc,omitempty"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	Templated bool   `json:"templated"`
}

type EmailResponse struct {
	Status string `json:"status"`
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
	Email     string `json:"email"`
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
	Email     string `json:"email"`
}

type RenewCredentials struct {
	Email       string `json:"email"`
	UpdateCreds bool   `json:"update_creds"`
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
	Email       string `json:"email"`
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

type CallbacksContent struct {
	Type                      string           `json:"type"`
	Endpoint                  string           `json:"endpoint"`
	Token                     string           `json:"token"`
	Client_id                 string           `json:"client_id"`
	User_data                 string           `json:"user_data"`
	Username                  string           `json:"username"`
	Password                  string           `json:"password"`
	Port                      string           `json:"port"`
	Subscription              string           `json:"subscription"`
	Qos                       string           `json:"qos"`
	Topic                     string           `json:"topic"`
}

type FunctionContent struct {
	Code      string             `json:"code"`
	Language  string             `json:"language"`
	Name      string             `json:"name"`
	Args      []string           `json:"args"`
	Regexp    string             `json:"regexp"`
	Callbacks []CallbacksContent `json:"callbacks"`
	Env       map[string]string  `json:"env"`
}

type Function struct {
	Id         string          `json:"id"`
	Owner_id   int             `json:"owner_id"`
	Content    FunctionContent `json:"content"`
	Is_public  bool            `json:"is_public"`
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

type InvocationDisplay struct {
	Id          string `json:"id"`
	Invoker_id  int    `json:"invoker_id"`
	Function_id string `json:"function_id"`
	State       string `json:"state"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type InvocationsResponse struct {
	Status     string       `json:"status"`
	Code       int          `json:"code"`
	StartIndex int          `json:"start_index"`
	MaxIndex   int          `json:"max_index"`
	Results    []Invocation `json:"results"`
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

type FunctionOwner struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type InvocationInvoker struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type TriggerOwner struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type ObjectTypeContent struct {
	Name string `json:"name"`
	Public bool `json:"public"`
	DecodingFunction string `json:"decoding_function"`
	Triggers []string `json:"triggers"`
}

type ObjectType struct {
	Id          string `json:"id"`
	User_id    int    `json:"user_id"`
	Content ObjectTypeContent `json:"content"`
}

type ObjectTypesDisplay struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Public bool `json:"public"`
	DecodingFunction string `json:"decoding_function"`
}

type UpdateObjectTypeBody struct {
	Content ObjectTypeContent `json:"content"`
}

type Device struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Typeobject_id string `json:"typeobject_id"`
	Active bool `json:"active"`
}

type DeviceDisplay struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Typeobject_id string `json:"typeobject_id"`
	Active bool `json:"active"`
}
