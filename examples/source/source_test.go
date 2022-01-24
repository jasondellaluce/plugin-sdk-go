package main

import (
	"testing"

	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins"
	"github.com/falcosecurity/plugin-sdk-go/pkg/sdk/plugins/source"
)

type WorkflowInput struct {
	APIVersion string
	InitConfig string
	OpenParams string
}

type WorkflowState struct {
	WorkflowInput
	Test     testing.TB
	Plugin   source.Plugin
	Continue bool // todo: maybe context.Context?
}

type Workflow struct {
	LastError  func(*WorkflowState)
	Info       func(*WorkflowState) *plugins.Info
	InitSchema func(*WorkflowState) string
	Init       func(*WorkflowState) error
	Open       func(*WorkflowState) (source.Instance, error)
	Extract    func(*WorkflowState, sdk.ExtractRequest, sdk.EventReader) error
	NextBatch  func(*WorkflowState, source.Instance) (int, error)
	Close      func(*WorkflowState)
	Destroy    func(*WorkflowState)
}

// var defaultWorkflow = &Workflow{
// 	Info: func(t *testing.T, input *WorkflowInput, i *plugins.Info) bool {
// 		if i.RequiredAPIVersion != input.APIVersion {
// 			//todo: semver check
// 			t.Fail()
// 			return false
// 		}
// 		return true
// 	},
// 	NextBatch: func(t *testing.T, in *WorkflowInput, i source.Instance) (int, bool) {
// 		n, err := i.NextBatch(in.Plugin, i.Events())
// 		if err == sdk.ErrEOF {
// 			return n, true
// 		}
// 	},
// }

func (w *Workflow) Run(t *testing.T, p source.Plugin, input *WorkflowInput) {
	// check infos
	w.Info(t, p, input)
	if !input.Continue {
		return
	}

	// fields

	// check init schema
	w.InitSchema(t, p, input)
	if !input.Continue {
		return
	}

	// perform init
	w.Init(t, p, input)
	if !input.Continue {
		return
	}

	// perform open
	i, _ := w.Open(t, p, input)
	if !input.Continue {
		return
	}

	for {
		n, err := w.NextBatch(t, p, input, i)
		if !input.Continue {
			return
		}

		if err == sdk.ErrEOF {
			break
		} else if err == sdk.ErrTimeout {
			continue
		}

		// w.Extract(.., ...)
		//

		// w.Progress(.., ...)
		//

		// w.String(.., ...)
		//
	}

	// close the plugin
	w.Close(t, p, input)
	if !input.Continue {
		return
	}

	// destroy the plugin
	w.Destroy(t, p, input)
	if !input.Continue {
		return
	}
}

// func TestInit(t *testing.T) {
// 	runner := NewRunner(&MyPlugin{})
// 	runner.OnInit("kudq", func(t *testing.T, err error) {

// 	})

// 	runner.Run(t)
// }
