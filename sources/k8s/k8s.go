package k8s

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cheebo/go-config"
	"github.com/cheebo/go-config/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type k8s struct {
	namespace string
	client    *kubernetes.Clientset
	data      map[string]interface{}
	k8s       map[string]interface{}
	values    map[string]*go_config.Variable
}

func Source(namespace string, client *kubernetes.Clientset) go_config.Source {
	return &k8s{
		namespace: namespace,
		client:    client,
		data:      make(map[string]interface{}),
	}
}

func (self *k8s) Init(vals map[string]*go_config.Variable) error {
	self.values = vals

	var data interface{}

	for name, val := range vals {
		name = self.name(name)
		namespace := self.namespace

		tag := val.Tag.Get("k8s")
		opts := strings.Split(tag, ",")

		if len(opts[0]) > 0 {
			// override namespace
			if len(opts) == 3 && len(opts[2]) > 0 {
				namespace = opts[2]
			}

			cmap, err := self.client.CoreV1().ConfigMaps(namespace).Get(opts[0], metav1.GetOptions{})
			if err != nil {
				return err
			}

			if len(opts) == 1 || (len(opts) > 1 && len(opts[1]) == 0) {
				// from literal
				data = cmap.Data
			} else {
				// from file
				if d, ok := cmap.Data[opts[1]]; ok {
					data = []byte(d)
				} else {
					return fmt.Errorf("CofigMap %s not contained file %s.", opts[0], opts[1])
				}
			}

			switch val.Type.Kind() {
			case reflect.Struct:
				fallthrough
			case reflect.Slice:
				m, err := utils.JsonParse(data)
				if err != nil {
					return err
				}
				for n, v := range m {
					self.data[name+"."+n] = v
				}
			default: // FIXME
				//self.data[name] = data
			}
		}
	}
	return nil
}

func (self *k8s) Int(name string) (int, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	return int(val.(float64)), nil
}

func (self *k8s) Float(name string) (float64, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}

	return float64(val.(float64)), nil
}

func (self *k8s) UInt(name string) (uint, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return 0, nil
	}
	return uint(val.(float64)), nil
}

func (self *k8s) String(name string) (string, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return "", nil
	}
	return val.(string), nil
}

func (self *k8s) Bool(name string) (bool, error) {
	val, ok := self.data[self.name(name)]
	if !ok {
		return false, nil
	}
	b, ok := val.(bool)
	if !ok {
		return false, nil
	}
	return b, nil
}

func (self *k8s) Slice(name, delimiter string, kind reflect.Kind) ([]interface{}, error) {
	return []interface{}{}, nil
}

func (self *k8s) Export(opt ...go_config.SourceOpt) ([]byte, error) {
	return []byte{}, nil
}

func (self *k8s) name(name string) string {
	return strings.ToLower(name)
}
