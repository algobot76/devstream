package pluginengine

import (
	"github.com/merico-dev/stream/pkg/util/log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/merico-dev/stream/internal/pkg/configloader"
	"github.com/merico-dev/stream/internal/pkg/statemanager"
)

// TestOutputsInState normal test, if the ref value is correct
func TestOutputsInState(t *testing.T) {
	keyReferrer := statemanager.StateKey("default-trello-github_trello-github-integ")
	stateReferrer := statemanager.State{
		Name:     "default-trello-github",
		Plugin:   configloader.Plugin{Kind: "trello-github-integ", Version: "0.2.0"},
		Options:  map[string]interface{}{"a": "value"},
		Resource: map[string]interface{}{"a": "value"},
	}

	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keyReferrer, stateReferrer)
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keyReferrer); err != nil {
			log.Errorf("failed to delete state %s.", keyReferrer)
		}
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	options := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "boardId": "${{ default.trello.outputs.boardId }}", "todoListId": "${{ default.trello.outputs.todoListId }}"}
	expectResult := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "boardId": "123", "todoListId": "456"}
	err = fillRefValueWithOutputs(smgr, options)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectResult, options)
}

// TestRefInDeeperLayerState normal test, if when the ref key is in deeper layer
func TestRefInDeeperLayerState(t *testing.T) {

	keyReferrer := statemanager.StateKey("default-trello-github_trello-github-integ")
	stateReferrer := statemanager.State{
		Name:     "default-trello-github",
		Plugin:   configloader.Plugin{Kind: "trello-github-integ", Version: "0.2.0"},
		Options:  map[string]interface{}{"a": "value"},
		Resource: map[string]interface{}{"a": "value"},
	}

	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keyReferrer, stateReferrer)
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keyReferrer); err != nil {
			log.Errorf("failed to delete state %s.", keyReferrer)
		}
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	options := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "app": map[string]interface{}{"boardId": "${{ default.trello.outputs.boardId }}", "todoListId": "${{ default.trello.outputs.todoListId }}"}}

	expectResult := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "app": map[string]interface{}{"boardId": "123", "todoListId": "456"}}
	err = fillRefValueWithOutputs(smgr, options)

	assert.Equal(t, nil, err)
	assert.Equal(t, expectResult, options)
}

// TestOutputsEmpty exception test, if there is no outputs in state.
func TestOutputsEmpty(t *testing.T) {

	keyReferrer := statemanager.StateKey("default-trello-github_trello-github-integ")
	stateReferrer := statemanager.State{
		Name:     "default-trello-github",
		Plugin:   configloader.Plugin{Kind: "trello-github-integ", Version: "0.2.0"},
		Options:  map[string]interface{}{"a": "value"},
		Resource: map[string]interface{}{"a": "value"},
	}

	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456"},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keyReferrer, stateReferrer)
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keyReferrer); err != nil {
			log.Errorf("failed to delete state %s.", keyReferrer)
		}
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	options := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "boardId": "${{ default.trello.outputs.boardId }}", "todoListId": "${{ default.trello.outputs.todoListId }}"}

	err = fillRefValueWithOutputs(smgr, options)

	assert.Equal(t, "cannot find outputs from state: default", err.Error())
}

// TestWrongRefFormat exception test, if the outputs do not contain ref
func TestWrongRefFormat(t *testing.T) {

	keyReferrer := statemanager.StateKey("default-trello-github_trello-github-integ")
	stateReferrer := statemanager.State{
		Name:     "default-trello-github",
		Plugin:   configloader.Plugin{Kind: "trello-github-integ", Version: "0.2.0"},
		Options:  map[string]interface{}{"a": "value"},
		Resource: map[string]interface{}{"a": "value"},
	}

	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keyReferrer, stateReferrer)
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keyReferrer); err != nil {
			log.Errorf("failed to delete state %s.", keyReferrer)
		}
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	options := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "boardId": "${{ default.trello.outputs.aaa }}"}

	err = fillRefValueWithOutputs(smgr, options)

	assert.Equal(t, "can not find aaa in dependency outputs", err.Error())
}

// TestWrongRefFormat exception test, if the ref length is correct.
func TestRefLength(t *testing.T) {

	keyReferrer := statemanager.StateKey("default-trello-github_trello-github-integ")
	stateReferrer := statemanager.State{
		Name:     "default-trello-github",
		Plugin:   configloader.Plugin{Kind: "trello-github-integ", Version: "0.2.0"},
		Options:  map[string]interface{}{"a": "value"},
		Resource: map[string]interface{}{"a": "value"},
	}

	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keyReferrer, stateReferrer)
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keyReferrer); err != nil {
			log.Errorf("failed to delete state %s.", keyReferrer)
		}
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	options := map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "branch": "main", "boardId": "${{ default.trello.outputs }}"}

	err = fillRefValueWithOutputs(smgr, options)

	assert.Equal(t, "ref input format is not correct: default.trello.outputs", err.Error())
}
