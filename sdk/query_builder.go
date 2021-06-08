package sdk

import (
	"github.com/RelationalAI/go-client-sdk/client"
)

type QueryArgs struct {
	Inputs    []client.Relation
	Outputs   []string
	Persist   []string
	Source    client.Source
	ReadOnly  bool
	SrcString string
}

func (queryArgs *QueryArgs) fill_defaults() {
	// setting default values
	queryArgs.Inputs = []client.Relation{*client.NewRelationWithDefaults()}
	queryArgs.Outputs = []string{}
	queryArgs.Persist = []string{}
	queryArgs.Source = *client.NewSourceWithDefaults()
	queryArgs.ReadOnly = false
	queryArgs.SrcString = ""
}
func NewQueryArgsWithDefaults() QueryArgs {
	queryArgs := new(QueryArgs)
	queryArgs.fill_defaults()

	return *queryArgs
}

func NewQueryArgs(inputs []client.Relation, outputs []string, persist []string, source client.Source, readonly bool, srcString string) *QueryArgs {
	this := QueryArgs{}
	this.Inputs = inputs
	this.Outputs = outputs
	this.Persist = persist
	this.Source = source
	this.ReadOnly = readonly
	this.SrcString = srcString
	return &this
}
