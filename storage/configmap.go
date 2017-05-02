package storage

// ConfigMap responsible for reading/writing data to configmap - therefore ReadWriter interface
type ConfigMap struct {
	//kubernetes client
}

// NewConfigMap creates a configmap in kubernetes cluster with all necessary information
func NewConfigMap(name, namespace string) *ConfigMap {
	return nil
}

func (cm *ConfigMap) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (cm *ConfigMap) Write(p []byte) (n int, err error) {
	return 0, nil
}
