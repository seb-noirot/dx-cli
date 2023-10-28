package k8s

type KubeConfig struct {
	APIVersion     string      `yaml:"apiVersion"`
	Clusters       []Cluster   `yaml:"clusters"`
	Contexts       []Context   `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
	Kind           string      `yaml:"kind"`
	Preferences    interface{} `yaml:"preferences"`
	Users          []UserEntry `yaml:"users"`
}

type Cluster struct {
	Name    string            `yaml:"name"`
	Cluster ClusterAttributes `yaml:"cluster"`
}

type ClusterAttributes struct {
	CertificateAuthority string `yaml:"certificate-authority"`
	Server               string `yaml:"server"`
}

type Context struct {
	Name    string            `yaml:"name"`
	Context ContextAttributes `yaml:"context"`
}

type ContextAttributes struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
}

type UserEntry struct {
	Name string   `yaml:"name"`
	User UserSpec `yaml:"user"`
}

type UserSpec struct {
	Exec              UserSpecExec `yaml:"exec"`
	ClientCertificate string       `yaml:"client-certificate,omitempty"`
	ClientKey         string       `yaml:"client-key,omitempty"`
}

type UserSpecExec struct {
	ApiVersion string   `yaml:"apiVersion"`
	Args       []string `yaml:"args"`
	Command    string   `yaml:"command"`
	Env        string   `yaml:"env"`
}
