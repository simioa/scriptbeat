package beater

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"

	"github.com/simioa/scriptbeat/config"
)

// scriptbeat configuration.
type scriptbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of scriptbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	var c config.Config
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &scriptbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts scriptbeat.
func (bt *scriptbeat) Run(b *beat.Beat) error {
	logp.Info("scriptbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, scriptConfig := range bt.config.Scripts {
		wg.Add(1)
		go bt.runScriptExecutor(scriptConfig, &wg)
	}
	wg.Wait()
	logp.Info("All Script Executors closed - Shutting down")
	return nil
}

func (bt *scriptbeat) runScriptExecutor(scriptConfig config.ScriptConfig, wg *sync.WaitGroup) {
	ticker := time.NewTicker(scriptConfig.Period)
	stdOutBuffer := new(bytes.Buffer)
	stdErrBuffer := new(bytes.Buffer)
	for {
		select {
		case <-bt.done:
			wg.Done()
			return
		case <-ticker.C:
		}
		cmd := &exec.Cmd{
			Path:   scriptConfig.Script,
			Args:   scriptConfig.Args,
			Env:    scriptConfig.Env,
			Stdout: stdOutBuffer,
			Stderr: stdErrBuffer,
		}
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"script": scriptConfig.Script,
				"alias":  scriptConfig.Alias,
			},
		}
		if err := cmd.Run(); err != nil {
			event.Fields.Put("error", err)
		} else {
			event.Fields.Put("stderr", stdErrBuffer.String())
			event.Fields.Put("stdout", stdOutBuffer.String())
		}
		bt.client.Publish(event)
	}
}

// Stop stops scriptbeat.
func (bt *scriptbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
