package congo

import (
	"fmt"
	"math/big"
	"time"
)

type ValueLocation fmt.Stringer

type Watcher interface {
	Cancel() error
}

type Change struct {
	Config   Config
	Key      string
	OldValue interface{}
	NewValue interface{}
}

type ConfigSetter interface {
	Config

	// Setters
	SetString(key, value string) error
	SetBool(key string, value bool) error
	SetInt(key string, value int) error
	SetBigInt(key string, value big.Int) error
	SetFloat(key string, value float64) error
	SetTime(key string, value time.Time) error
	SetDuration(key string, value time.Duration) error
	SetBytes(key string, value []byte) error
	SetBigFloat(key string, value big.Float) error

	// Remover
	Remove(key string) (interface{}, error)

	ChildSetter(prefix string) (ConfigSetter, error)
	RootSetter() ConfigSetter
}

type Config interface {
	Setter() (ConfigSetter, error)

	GetValueLocation(key string) (ValueLocation, error)

	WatchValueChanges(key string, c chan Change, err chan error) (Watcher, error)

	// HasKey retorna true si `key` existe en el Config actual
	Has(key string) bool

	// Get popula `a` utilizando reflection para obtener los valores de un árbol de configuración
	Get(key string, a interface{}) error
	// Watch monitorea cambios en el valor de `key` populando instancias de `a` y enviándo error o nil a `c` cuando detecta cambios
	Watch(key string, a interface{}, c chan error) (Watcher, error)

	GetString(key string) (string, error)
	WatchString(key string, a *string, c chan error) (Watcher, error)

	GetBool(key string) (bool, error)
	WatchBool(key string, a *bool, c chan error) (Watcher, error)

	GetInt(key string) (int, error)
	WatchInt(key string, a *int, c chan error) (Watcher, error)

	GetBigInt(key string) (big.Int, error)
	WatchBigInt(key string, a *big.Int, c chan error) (Watcher, error)

	GetFloat(key string) (float64, error)
	WatchFloat(key string, a *float64, c chan error) (Watcher, error)

	GetTime(key string) (time.Time, error)
	WatchTime(key string, a *time.Time, c chan error) (Watcher, error)

	GetDuration(key string) (time.Duration, error)
	WatchDuration(key string, a *time.Duration, c chan error) (Watcher, error)

	GetBytes(key string) ([]byte, error)
	WatchBytes(key string, a []byte, c chan error) (Watcher, error)

	GetBigFloat(key string) (big.Float, error)
	WatchBigFloat(key string, a *big.Float, c chan error) (Watcher, error)

	// Child retorna un rama completa de configuración
	Child(prefix string) (Config, error)

	// Root retorna la configuración raíz
	Root() Config
}
