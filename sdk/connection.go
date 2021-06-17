package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	llclient "github.com/RelationalAI-oss/relationalai-sdk-go/client"
	"go.uber.org/zap"
)

type Connection struct {
	database   string
	apiService *llclient.DefaultApiService
	logger     *zap.Logger
	debug      bool
	version    int32
}

type ConnectionOptions struct {
	RAIServerAddress string
	Debug            bool
	Logger           *zap.Logger
}

type option func(*ConnectionOptions)

type Mode string

const (
	Create          = Mode("CREATE")
	CreateOverwrite = Mode("CREATE_OVERWRITE")
	Clone           = Mode("CLONE")
	CloneOverwrite  = Mode("CLONE_OVERWRITE")
)

func WithRAIServerAddress(address string) option {
	return func(c *ConnectionOptions) {
		c.RAIServerAddress = strings.TrimSuffix(address, "/")
	}
}

func WithDebug(debug bool) option {
	return func(c *ConnectionOptions) {
		c.Debug = debug
	}
}

func WithLogger(logger *zap.Logger) option {
	return func(c *ConnectionOptions) {
		c.Logger = logger
	}
}

func defaultConnectionOptions() *ConnectionOptions {
	return &ConnectionOptions{
		RAIServerAddress: "http://127.0.0.1:8010",
		Debug:            false,
		Logger:           zap.NewNop(),
	}
}

func NewConnection(database string, opts ...option) *Connection {
	// Start with default options and apply user given options.
	dopts := defaultConnectionOptions()
	for _, optApply := range opts {
		optApply(dopts)
	}

	cfg := &llclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "rai-go-sdk",
		Debug:         false,
		Servers: llclient.ServerConfigurations{
			{
				URL: dopts.RAIServerAddress,
			},
		},
		OperationServers: map[string]llclient.ServerConfigurations{},
	}

	return &Connection{
		database:   database,
		apiService: llclient.NewAPIClient(cfg).DefaultApi,
		logger:     dopts.Logger,
		debug:      dopts.Debug,
	}
}

type CreateDatabaseResult struct {
	Success  bool
	Problems []interface{}
}

func (conn *Connection) CreateDatabase(ctx context.Context, mode Mode) (*CreateDatabaseResult, error) {
	txnAction := llclient.NewTransactionWithDefaults()
	apiTransactionPostRequest := conn.apiService.TransactionPost(ctx)
	txnAction.SetMode(string(mode))
	txnAction.SetDbname(conn.database)
	txnAction.SetActions([]llclient.LabeledAction{})
	txnAction.SetReadonly(false)

	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"payload", txnAction,
		)
	}

	result, response, err := apiTransactionPostRequest.Transaction(*txnAction).Execute()
	if err.Error() != "" {
		return nil, fmt.Errorf("transaction failed with %v", err.Model())
	}
	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"status_code", response.StatusCode,
			"problems", result.Problems,
		)
	}
	success := !result.Aborted

	return &CreateDatabaseResult{
		Success:  success,
		Problems: result.Problems,
	}, nil
}

type CloneDatabaseResult struct {
	Success  bool
	Problems []interface{}
}

func (conn *Connection) CloneDatabase(ctx context.Context, sourceName string, mode Mode) (*CloneDatabaseResult, error) {
	txnAction := llclient.NewTransactionWithDefaults()
	apiTransactionPostRequest := conn.apiService.TransactionPost(ctx)
	txnAction.SetMode(string(mode))
	txnAction.SetDbname(conn.database)
	txnAction.SetActions([]llclient.LabeledAction{})
	txnAction.SetSourceDbname(sourceName)
	txnAction.SetReadonly(false)

	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"payload", txnAction,
		)
	}

	result, response, err := apiTransactionPostRequest.Transaction(*txnAction).Execute()
	if err.Error() != "" {
		return nil, fmt.Errorf("transaction failed with %v", err.Model())
	}
	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"status_code", response.StatusCode,
			"problems", result.Problems,
		)
	}
	success := !result.Aborted

	return &CloneDatabaseResult{
		Success:  success,
		Problems: result.Problems,
	}, nil
}

type CardinalityResult struct {
	Relations []llclient.Relation
	Problems  []interface{}
}

func (conn *Connection) Cardinality(ctx context.Context, relNames ...string) (*CardinalityResult, error) {
	if len(relNames) > 1 {
		return nil, fmt.Errorf("only provide one relation name there are %v relation names", len(relNames))
	}
	var relName *string = nil
	if len(relNames) > 0 {
		relName = &relNames[0]
	}

	action := llclient.NewCardinalityActionWithDefaults()
	if relName != nil {
		action.SetRelname(*relName)
	}
	action.Action.SetType("CardinalityAction")
	var res llclient.CardinalityActionResult
	problems, err := conn.runAction(ctx, conn.database, "single", action, true, "OPEN", &res)
	if err != nil {
		return nil, err
	}
	return &CardinalityResult{
		Relations: res.GetResult(),
		Problems:  problems,
	}, nil
}

type DeleteEdbResult struct {
	Relkeys  []llclient.RelKey
	Problems []interface{}
}

func (conn *Connection) DeleteEdb(ctx context.Context, relName string) (*DeleteEdbResult, error) {
	if relName == "" {
		return nil, errors.New("empty relName")
	}
	action := llclient.NewModifyWorkspaceActionWithDefaults()
	action.SetDeleteEdb(relName)
	action.Action.SetType("ModifyWorkspaceAction")
	var res llclient.ModifyWorkspaceActionResult

	problems, err := conn.runAction(ctx, conn.database, "single", action, false, "OPEN", &res)
	if err != nil {
		return nil, err
	}
	return &DeleteEdbResult{
		Relkeys:  res.DeleteEdbResult,
		Problems: problems,
	}, nil
}

type EnableLibraryResult struct {
	Success  bool
	Problems []interface{}
}

func (conn *Connection) EnableLibrary(ctx context.Context, srcName string) (*EnableLibraryResult, error) {
	action := llclient.NewModifyWorkspaceActionWithDefaults()
	action.SetEnableLibrary(srcName)
	action.Action.SetType("ModifyWorkspaceAction")

	problems, err := conn.runAction(ctx, conn.database, "single", action, false, "OPEN", nil)
	if err != nil {
		return nil, err
	}
	return &EnableLibraryResult{
		Success:  true,
		Problems: problems,
	}, nil
}

type ListEdbResult struct {
	Relkeys  []llclient.RelKey
	Problems []interface{}
}

func (conn *Connection) ListEdb(ctx context.Context, relNames ...string) (*ListEdbResult, error) {
	if len(relNames) > 1 {
		return nil, fmt.Errorf("only provide one relation name there are %v relation names", len(relNames))
	}
	var relName *string = nil
	if len(relNames) > 0 {
		relName = &relNames[0]
	}

	action := llclient.NewListEdbAction()
	if relName != nil {
		action.SetRelname(*relName)
	}
	action.Action.SetType("ListEdbAction")
	var res llclient.ListEdbActionResult
	problems, err := conn.runAction(ctx, conn.database, "single", action, false, "OPEN", &res)
	if err != nil {
		return nil, err
	}
	return &ListEdbResult{
		Relkeys:  res.Rels,
		Problems: problems,
	}, nil
}

type InstallSourceResult struct {
	Success  bool
	Problems []interface{}
}

func (conn *Connection) InstallSource(ctx context.Context, sources []llclient.Source) (*InstallSourceResult, error) {
	action := llclient.NewInstallAction()
	sourcesNew := []llclient.Source{}
	for _, src := range sources {
		if len(strings.TrimSpace(src.Path)) == 0 {
			sourcesNew = append(sourcesNew, src)
		} else {
			newSrc, err := conn.readFileFromPath(src)
			if err != nil {
				return nil, err
			}
			sourcesNew = append(sourcesNew, *newSrc)
		}
	}
	action.SetSources(sourcesNew)
	action.Action.SetType("InstallAction")
	problems, err := conn.runAction(ctx, conn.database, "single", action, false, "OPEN", nil)
	if err != nil {
		return nil, err
	}
	return &InstallSourceResult{
		Success:  true,
		Problems: problems,
	}, nil
}

type ListSourceResult struct {
	Sources  map[string]*llclient.Source
	Problems []interface{}
}

func (conn *Connection) ListSource(ctx context.Context) (*ListSourceResult, error) {
	action := llclient.NewListSourceActionWithDefaults()
	action.Action.SetType("ListSourceAction")
	var res llclient.ListSourceActionResult
	problems, err := conn.runAction(ctx, conn.database, "single", action, true, "OPEN", &res)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]*llclient.Source)
	for _, src := range res.GetSources() {
		resultMap[src.GetName()] = &src
	}
	return &ListSourceResult{
		Sources:  resultMap,
		Problems: problems,
	}, nil
}

type DeleteSourcesResult struct {
	Success  bool
	Problems []interface{}
}

func (conn *Connection) DeleteSources(ctx context.Context, scrNameList []string) (*DeleteSourcesResult, error) {
	action := llclient.NewModifyWorkspaceActionWithDefaults()
	action.SetDeleteSource(scrNameList)
	action.Action.SetType("ModifyWorkspaceAction")
	problems, err := conn.runAction(ctx, conn.database, "single", action, false, "OPEN", nil)
	if err != nil {
		return nil, err
	}
	return &DeleteSourcesResult{
		Success:  true,
		Problems: problems,
	}, nil
}

type QueryResult struct {
	Relations []llclient.Relation
	Problems  []interface{}
}

func (conn *Connection) Query(ctx context.Context, queryArgs QueryArgs) (*QueryResult, error) {
	action := llclient.NewQueryActionWithDefaults()
	action.SetInputs(queryArgs.Inputs)
	action.SetOutputs(queryArgs.Outputs)
	action.SetPersist(queryArgs.Persist)
	if queryArgs.SrcString != "" {
		src := llclient.NewSourceWithDefaults()
		src.Name = "query"
		src.Value = queryArgs.SrcString
		action.SetSource(*src)
	} else {
		if len(strings.TrimSpace(queryArgs.Source.Path)) == 0 {
			action.SetSource(queryArgs.Source)
		} else {
			newSource, err := conn.readFileFromPath(queryArgs.Source)
			if err != nil {
				return nil, err
			}
			action.SetSource(*newSource)
		}
	}
	action.Action.SetType("QueryAction")
	var res llclient.QueryActionResult
	problems, err := conn.runAction(ctx, conn.database, "single", action, queryArgs.ReadOnly, "OPEN", &res)
	if err != nil {
		return nil, err
	}
	return &QueryResult{
		Relations: res.Output,
		Problems:  problems,
	}, nil
}

type action interface {
	GetType() string
}

func (conn *Connection) readFileFromPath(source llclient.Source) (*llclient.Source, error) {
	file, err := os.Open(source.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if len(strings.TrimSpace(source.Name)) == 0 {
		source.SetName(filepath.Base(source.Path))
	}
	if len(strings.TrimSpace(source.Value)) == 0 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(file)
		contents := buf.String()
		source.SetValue(contents)
	}

	return &source, nil
}

func (conn *Connection) runAction(ctx context.Context, dbname string, name string, action action, isReadOnly bool, mode string, res interface{}) ([]interface{}, error) {
	apiTransactionPostRequest := conn.apiService.TransactionPost(ctx)
	txnAction := llclient.NewTransactionWithDefaults()
	txnAction.SetMode(mode)
	txnAction.SetDbname(dbname)
	txnAction.SetReadonly(isReadOnly)
	txnAction.SetVersion(conn.version)

	labeledAction := llclient.NewLabeledActionWithDefaults()
	labeledAction.SetName(name)
	labeledAction.SetAction(action)
	txnAction.SetActions([]llclient.LabeledAction{*labeledAction})

	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"payload", txnAction,
		)
	}

	result, response, err := apiTransactionPostRequest.Transaction(*txnAction).Execute()

	if err.Error() != "" {
		return nil, fmt.Errorf("transaction failed with %v", err.Model())
	}

	if conn.debug {
		conn.logger.Sugar().Debugw("transaction executed",
			"status_code", response.StatusCode,
			"problems", result.Problems,
		)
	}

	if result.Aborted {
		return nil, fmt.Errorf("transaction is aborted. problems: %+v output: %+v", result.Problems, result.Output)
	}

	currentVersion := conn.version
	responseVersion := result.GetVersion()
	// Sync the reported database version to our local
	// connection version. Important, as we want to ensure
	// that in subsequent transactions this will be the
	// minimum required version of the database. Note that
	// only write transactions bump the version.
	if responseVersion > currentVersion {
		conn.version = responseVersion
	}
	var ret json.RawMessage
	for _, act := range result.Actions {
		if name == act.GetName() {
			ret = act.GetResult()
			break
		}
	}
	// If caller wants to get a result.
	if res != nil {
		if err := json.Unmarshal(ret, &res); err != nil {
			return nil, fmt.Errorf("decoding results failed. problems: %+v error: %v", result.Problems, err)
		}
	}

	return result.Problems, nil
}
