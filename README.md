# ezconfig

## Installation
```bash
go get github.com/gbaranski/ezconfig
```

## Examples

```go
import (
  "github.com/gbaranski/ezconfig"
)

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
```
