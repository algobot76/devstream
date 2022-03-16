package pluginengine

import (
	"github.com/merico-dev/stream/internal/pkg/configloader"
	"github.com/merico-dev/stream/internal/pkg/statemanager"
	"github.com/merico-dev/stream/pkg/util/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetOutputs normal test, get Outputs by Statekey
func TestGetOutputs(t *testing.T) {
	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	outputs, err := smgr.GetOutputs(keystateReferee)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"boardId": "123", "todoListId": "456"}, outputs)
}

// TestGetOutputsIfEmpty exeption test, get Outputs by Statekey, but Outputs is empty
func TestGetOutputsIfEmpty(t *testing.T) {
	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456"},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	outputs, err := smgr.GetOutputs(keystateReferee)
	assert.NoError(t, err)
	assert.Equal(t, nil, outputs)
}

// TestGetOutputs exeption test, get Outputs by Statekey, but Statekey is wrong
func TestGetOutputsIfWrongKey(t *testing.T) {
	keystateReferee := statemanager.StateKey("default_trello")
	stateReferee := statemanager.State{
		Name:     "default",
		Plugin:   configloader.Plugin{Kind: "trello", Version: "0.2.0"},
		Options:  map[string]interface{}{"owner": "lfbdev", "repo": "golang-demo", "kanbanBoardName": "kanban-name"},
		Resource: map[string]interface{}{"boardId": "123", "todoListId": "456", "outputs": map[string]interface{}{"boardId": "123", "todoListId": "456"}},
	}

	smgr, err := statemanager.NewManager()
	assert.NoError(t, err)
	err = smgr.AddState(keystateReferee, stateReferee)
	assert.NoError(t, err)

	defer func() {
		if err := smgr.DeleteState(keystateReferee); err != nil {
			log.Errorf("failed to delete state %s.", keystateReferee)
		}
	}()

	outputs, err := smgr.GetOutputs("wrong_key")
	assert.NoError(t, err)
	assert.Equal(t, nil, outputs)
}
