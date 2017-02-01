package datatype

import (
	"encoding/json"
)

// FrameContext is a line number and string of this line
type FrameContext struct {
	LineNumber float64
	Line       string
}

// UnmarshalJSON implements the interface needed to convert to a frame context
func (f *FrameContext) UnmarshalJSON(data []byte) error {
	var frameline []interface{}
	err := json.Unmarshal(data, &frameline)
	if err != nil {
		return err
	}

	f.LineNumber = frameline[0].(float64)
	f.Line = frameline[1].(string)

	return nil
}

// Frame represents a frame in a stacktrace/exception
type Frame struct {
	Filename           *string                 `json:"filename,omitempty"`
	AbsolutePath       *string                 `json:"absPath,omitempty"`
	Module             *string                 `json:"module,omitempty"`
	Package            *string                 `json:"package,omitempty"`
	InstructionAddress *string                 `json:"instructionAddr,omitempty"`
	InstructionOffSet  *string                 `json:"instructionOffSet,omitempty"`
	Function           *string                 `json:"function,omitempty"`
	Errors             *[]string               `json:"errors,omitempty"`
	ColumnNo           *float64                `json:"columnNo,omitempty"`
	InApp              *bool                   `json:"inApp,omitempty"`
	Platform           *string                 `json:"platform,omitempty"`
	Context            *[]FrameContext         `json:"context,omitempty"`
	Vars               *map[string]interface{} `json:"vars,omitempty"`
	LineNo             *float64                `json:"lineNo,omitempty"`
	SymboleAddr        *string                 `json:"symboleAddr,omitempty"`
	Symbol             *string                 `json:"symbol,omitempty"`
}

//Exception implements a exception which olds stacktraces
type Exception struct {
	Values          *[]SingleException `json:"values,omitempty"`
	HasSystemFrames *bool              `json:"hasSystemFrames,omitempty"`
	ExcOmitted      *bool              `json:"excOmitted,omitempty"`
}

// SingleException represents a single exception in a exception
type SingleException struct {
	Type       *string            `json:"type,omitempty"`
	Value      *string            `json:"value,omitempty"`
	Mechanism  *map[string]string `json:"mechanism,omitempty"`
	Stacktrace *Stacktrace        `json:"stacktrace,omitempty"`
}

//Stacktrace implements the sentry interface for a stack trace in a event request
type Stacktrace struct {
	Frames          []Frame   `json:"frames,omitempty"`
	FramesOmitted   []float64 `json:"framesOmitted,omitempty"`
	HasSystemFrames bool      `json:"hasSystemFrames,omitempty"`
}
