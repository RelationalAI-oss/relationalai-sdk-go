// +build integration

package sdk

import (
	"context"
	"testing"

	"github.com/RelationalAI-oss/relationalai-sdk-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const dbname = "testDb1"
const dbname2 = "testDb2"

func TestWorkflow(t *testing.T) {
	logger := zap.NewExample()
	conn := NewConnection(dbname, WithDebug(true), WithLogger(logger))
	ctx := context.Background()

	// 1. create database
	createDatabaseResult, err := conn.CreateDatabase(ctx, CreateOverwrite)
	require.Nil(t, err)
	require.Equal(t, true, createDatabaseResult.Success)

	// 2. install a source
	src := client.NewSourceWithDefaults()
	src.Name = "test-1.delve"
	src.Value = "def foo = {(1,);(2,);(3,)}"
	installSourceResult, err := conn.InstallSource(ctx, []client.Source{*src})
	require.Nil(t, err)
	require.Equal(t, true, installSourceResult.Success)
	require.Equal(t, 0, len(installSourceResult.Problems))

	// 3. clone database
	conn2 := NewConnection(dbname2, WithDebug(true), WithLogger(logger))
	createDatabaseResult, err = conn2.CreateDatabase(ctx, CreateOverwrite)
	require.Nil(t, err)
	require.Equal(t, true, createDatabaseResult.Success)

	cloneDatabaseResult, err := conn2.CloneDatabase(ctx, conn.database, CloneOverwrite)
	assert.Nil(t, err)
	assert.Equal(t, true, cloneDatabaseResult.Success)

	// 4. list all the sources for the db
	listSourceResult, err := conn.ListSource(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(listSourceResult.Sources))

	// 5. delete a source
	deleteSourcesResult, err := conn.DeleteSources(ctx, []string{"test-1.delve"})
	assert.Nil(t, err)
	assert.Equal(t, true, deleteSourcesResult.Success)

	// 6. insert data in database
	queryArgs := NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def insert[:x1] = 1"
	queryArgs.ReadOnly = false

	queryResult, err := conn.Query(ctx, queryArgs)
	require.Nil(t, err)
	require.Equal(t, 0, len(queryResult.Relations))

	// 7. list Edb
	relName := "x1"
	listEdbResult, err := conn.ListEdb(ctx, relName)
	assert.Nil(t, err)
	assert.Equal(t, relName, listEdbResult.Relkeys[0].Name)

	// 8. query database
	// query 1
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def p = {1; 2; 3}"
	queryArgs.Outputs = []string{"p"}

	queryResult, err = conn.Query(ctx, queryArgs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queryResult.Relations))
	// query 2
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def bar = 2"
	queryArgs.Outputs = []string{"bar"}

	queryResult, err = conn.Query(ctx, queryArgs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queryResult.Relations))
	assert.Equal(t, queryResult.Relations[0].RelKey.Name, "bar")
	// query 3
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def p = {(1.1,); (2.2,); (3.4,)}"
	queryArgs.Outputs = []string{"p"}

	queryResult, err = conn.Query(ctx, queryArgs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queryResult.Relations))
	assert.Equal(t, 3, len(queryResult.Relations[0].GetColumns()[0]))

	// query 4
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def p = {(parse_decimal[64, 2, \"1.1\"],); (parse_decimal[64, 2, \"2.2\"],); (parse_decimal[64, 2, \"3.4\"],)}"
	queryArgs.Outputs = []string{"p"}

	queryResult, err = conn.Query(ctx, queryArgs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queryResult.Relations))
	assert.Equal(t, 1.1, queryResult.Relations[0].GetColumns()[0][0].(float64))

	// query 5
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.SrcString = "def p = {(1, 5); (2, 7); (3, 9)}"
	queryArgs.Outputs = []string{"p"}

	queryResult, err = conn.Query(ctx, queryArgs)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(queryResult.Relations))
	assert.Equal(t, float64(5), queryResult.Relations[0].GetColumns()[1][0].(float64))

	// 9. cardinality
	queryArgs = NewQueryArgsWithDefaults()
	queryArgs.Source.Value = "def p = {(1,); (2,); (3,)}"
	queryArgs.Persist = []string{"p"}
	queryArgs.ReadOnly = false
	queryResult, err = conn.Query(ctx, queryArgs)
	require.Nil(t, err)
	require.Equal(t, 0, len(queryResult.Relations))
	cardinalityResult, err := conn.Cardinality(ctx, "p")
	require.Nil(t, err)
	require.Equal(t, 1, len(cardinalityResult.Relations))
	require.Equal(t, float64(3), cardinalityResult.Relations[0].GetColumns()[0][0].(float64))

	//10. delete Edb
	deleteEdbResult, kerr := conn.DeleteEdb(ctx, "p")
	assert.Nil(t, kerr)
	assert.Equal(t, 1, len(deleteEdbResult.Relkeys))
	relKey := client.NewRelKeyWithDefaults()
	relKey.Name = "p"
	relKey.Keys = []string{"Int64"}
	relKey.Values = []string{}
	assert.Equal(t, deleteEdbResult.Relkeys[0], *relKey)

	//10. Enable Library
	enableLibraryResult, lerr := conn.EnableLibrary(ctx, "stdlib")
	assert.Nil(t, lerr)
	assert.Equal(t, true, enableLibraryResult.Success)

}
