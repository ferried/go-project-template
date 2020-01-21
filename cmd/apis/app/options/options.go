package options

import (
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"os"
	"os/signal"
	"syscall"
)
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})

type CustomOptions struct {
	name  string
	value string
	age   int
}
type ServerRunOptions struct {
	CustomOptions CustomOptions
}

func NewServerRunOptions() *ServerRunOptions {
	options := ServerRunOptions{
		CustomOptions{"啊哈哈", "哦吼吼", 12},
	}
	return &options
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {
	s.CustomOptions.AddFlags(fss.FlagSet("sailor"))
	return fss
}

func (c *CustomOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.name, "name", "嘻嘻嘻", "随便起的名字")
	fs.StringVar(&c.value, "value", "呵呵呵", "随便起的名字")
	fs.IntVar(&c.age, "age", 10, "随便一个数字")
}


func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

