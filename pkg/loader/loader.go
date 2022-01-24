/*
Copyright (C) 2021 The Falco Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package loader

import "errors"

type PluginState interface{}

type PluginType uint32

type SchemaType uint32

type FieldType uint32

const (
	PluginTypeSource    PluginType = 1
	PluginTypeExtractor PluginType = 1
)

const (
	SchemaTypeNone SchemaType = 0
	SchemTypeJSON  SchemaType = 1
)

const (
	FieldTypeU64    FieldType = 8
	FieldTypeString FieldType = 9
)

var (
	ErrNoSymbol     = errors.New("no symbol")
	ErrFailure      = errors.New("failure")
	ErrTimeout      = errors.New("timeout")
	ErrEOF          = errors.New("eof")
	ErrNotSupported = errors.New("not supported")
)

type Event struct {
	EvtNum    uint64
	Data      []byte
	Timestamp uint64
}

type ExtractField struct {
	FieldID      uint32
	Field        string
	Arg          string
	FieldType    uint32 // todo
	FieldPresent bool
	Res          interface{}
}

type Plugin interface {
	Unload()
	GetType() PluginType
	GetName() string
	GetDescription() string
	GetContact() string
	GetVersion() string
	GetRequiredAPIVersion() string
	GetInitSchema() (string, SchemaType)
	Init(config string) (PluginState, error)
	Destroy(s PluginState)
	GetLastError(s PluginState) error
}

type SourcePlugin interface {
	Plugin
	GetID() uint32
	GetEventSource() string
	GetFields() string
	EventToString(s PluginState, data []byte) string
	ExtractFields(s PluginState, evt *Event, fields []*ExtractField) error
	ListOpenParams(s PluginState) (string, error)
	Open(s PluginState, params string) (PluginState, error)
	Close(s, h PluginState)
	GetProgress(s, h PluginState) (string, float64)
	NextBatch(s, h PluginState) ([]*Event, error)
}

type ExtractorPlugin interface {
	Plugin
	GetFields() string
	GetExtractEventSources() string
	ExtractFields(s PluginState, evt *Event, fields []*ExtractField) error
}

func LoadPlugin(path string) (Plugin, error) {
	return nil, nil
}
