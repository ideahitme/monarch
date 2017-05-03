package storage

// ConfigMap responsible for reading/writing data to configmap - therefore ReadWriter interface
type ConfigMap struct {
	//kubernetes client
}

// NewConfigMap creates a configmap in kubernetes cluster with all necessary information
func NewConfigMap(name, namespace string) *ConfigMap {
	return nil
}

// SetValue updates a value in the configmap
func (cm *ConfigMap) SetValue(key, value string) error {
	return nil
}

func (cm *ConfigMap) ReadValue(key string) (string, error) {
	return "", nil
}

func (cm *ConfigMap) Keys() ([]string, error) {
	return nil, nil
}
