package congo

import (
	"fmt"
	"math/big"
	"time"
)

func ExampleGetString() {
	e := exampleConfig{}
	str, _ := e.GetString("test")
	fmt.Print(str)
	// Output: test string
}

func ExampleWatchString() {
	e := exampleConfig{}
	ch := make(chan error, 1)
	str := ""
	e.WatchString("watchable-string", &str, ch)

	for {
		select {
		case x := <-ch:
			if x != nil {
				fmt.Printf("%s", x)
			} else {
				fmt.Print(str)
			}
			return
		}
	}
	// Output: watchable test string
}

func ExampleCancelWatchString() {
	e := exampleConfig{}
	ch := make(chan error, 1)
	str := ""
	w, _ := e.WatchString("cancel-watchable-string", &str, ch)

	time.Sleep(300 * time.Millisecond)

	select {
	case x := <-ch:
		if x != nil {
			fmt.Printf("%s", x)
		} else {
			fmt.Print(str)
		}
	}

	w.Cancel()
	// Output: watchable test string
}

/// Example Congo Config

type exampleConfig struct {
}

func (e *exampleConfig) Setter() (ConfigSetter, error) {
	return nil, nil
}

func (e *exampleConfig) GetValueLocation(key string) (ValueLocation, error) {
	return nil, nil
}

func (e *exampleConfig) WatchValueChanges(key string, c chan Change, err chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) Get(key string, a interface{}) error {
	return nil
}

func (e *exampleConfig) Watch(key string, a interface{}, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetString(key string) (string, error) {
	if key == "test" {
		return "test string", nil
	}
	return "", nil
}

type stringWatcher struct {
	Dest      *string
	Channel   chan error
	Value     string
	Repeat    bool
	Cancelled bool
}

func (s stringWatcher) Cancel() error {
	s.Repeat = false
	return nil
}

func (s stringWatcher) Run() {
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			*s.Dest = s.Value
			s.Channel <- nil
			if !s.Repeat {
				break
			}
		}
	}()
}

func (e *exampleConfig) WatchString(key string, a *string, c chan error) (Watcher, error) {
	w := stringWatcher{
		Dest:    a,
		Channel: c,
		Value:   "watchable test string",
	}
	if key == "watchable-string" {
		w.Repeat = false
		defer w.Run()
		return w, nil
	} else if key == "cancel-watchable-string" {
		w.Repeat = true
		defer w.Run()
		return w, nil
	}
	return nil, nil
}

func (e *exampleConfig) GetBool(key string) (bool, error) {
	return false, nil
}
func (e *exampleConfig) WatchBool(key string, a *bool, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetInt(key string) (int, error) {
	return 0, nil
}
func (e *exampleConfig) WatchInt(key string, a *int, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetBigInt(key string) (big.Int, error) {
	return big.Int{}, nil
}
func (e *exampleConfig) WatchBigInt(key string, a *big.Int, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetFloat(key string) (float64, error) {
	return 0.0, nil
}
func (e *exampleConfig) WatchFloat(key string, a *float64, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetTime(key string) (time.Time, error) {
	return time.Time{}, nil
}
func (e *exampleConfig) WatchTime(key string, a *time.Time, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetDuration(key string) (time.Duration, error) {
	return time.Duration(0), nil
}
func (e *exampleConfig) WatchDuration(key string, a *time.Duration, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetBytes(key string) ([]byte, error) {
	return []byte{}, nil
}
func (e *exampleConfig) WatchBytes(key string, a []byte, c chan error) (Watcher, error) {
	return nil, nil
}

func (e *exampleConfig) GetBigFloat(key string) (big.Float, error) {
	return big.Float{}, nil
}
func (e *exampleConfig) WatchBigFloat(key string, a *big.Float, c chan error) (Watcher, error) {
	return nil, nil
}

// Child retorna un rama completa de configuración
func (e *exampleConfig) Child(prefix string) (Config, error) {
	return nil, nil
}

// Root retorna la configuración raíz
func (e *exampleConfig) Root() Config {
	return e
}
