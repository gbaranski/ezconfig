package ezconfig

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// Parse parses struct passed by argument by reference
//
// Example
/*
type AppConfig struct {
  DatabaseURL           string  `env:"DATABASE_URL"`
  MaxReconnectAttempts  int     `env:"MAX_RECONNECT_ATTEMPTS"`
}

func main() {
  var config AppConfig

  err := ezconfig.Parse(&config)
  if err != nil {
    panic(fmt.Errorf("fail read config: %s", err.Error()))
  }

  fmt.Printf("DatabaseURL: %s\n", config.DatabaseURL)
  fmt.Printf("MaxReconnectAttempts: %d\n", config.MaxReconnectAttempts)
}
*/
func Parse(v interface{}) error {
	ptrv := reflect.ValueOf(v)
	if ptrv.Kind() != reflect.Ptr || ptrv.IsNil() {
		return fmt.Errorf("invalid type, expected struct, received %s", ptrv.Kind().String())
	}
	ev := ptrv.Elem()
	if ev.Kind() != reflect.Struct {
		return fmt.Errorf("invalid type, expected *struct, received %s", ev.Type().Name())

	}
	for i := 0; i < ev.NumField(); i++ {
		f := ev.Field(i)
		tf := ev.Type().Field(i)
		if !f.IsValid() {
			return fmt.Errorf("field %d is not valid", i)
		}
		if !f.CanSet() {
			return fmt.Errorf("field %d is not changeable", i)
		}
		tag, texists := tf.Tag.Lookup("env")
		if !texists {
			continue
		}
		env, eexists := os.LookupEnv(tag)
		if !eexists {
			return fmt.Errorf("ENV %s not found", tag)
		}

		switch f.Kind() {
		case reflect.String:
			f.SetString(env)
		case reflect.Int:
			intenv, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("%s: %s", tf.Name, err.Error())
			}
			f.SetInt(int64(intenv))
		default:
			return fmt.Errorf("kind %s not supported", f.Kind().String())

		}
	}
	return nil
}
