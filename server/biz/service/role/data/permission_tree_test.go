package data

import (
	"testing"

	"github.com/moumou/server/biz/model"
	"github.com/stretchr/testify/assert"
)

func TestBuildPermissionTree(t *testing.T) {
	permissions := []*model.Permission{
		{BaseModel: model.BaseModel{Id: 1}, Name: "Topa", Pid: 0},
		{BaseModel: model.BaseModel{Id: 2}, Name: "Topb", Pid: 0},
		{BaseModel: model.BaseModel{Id: 3}, Name: "Topa1", Pid: 1},
		{BaseModel: model.BaseModel{Id: 4}, Name: "Topa2", Pid: 1},
		{BaseModel: model.BaseModel{Id: 5}, Name: "Topb1", Pid: 2},
	}

	tree := BuildPermissionTree(permissions)
	// get by id
	assert.Equal(t, int64(2), tree.GetPermissionById(2).Id)
	assert.Equal(t, int64(3), tree.GetPermissionById(3).Id)

	// tree
	topLevelPermissions := tree.GetTopLevelPermissions()
	assert.Equal(t, int64(1), topLevelPermissions[0].Id)
	assert.Equal(t, int64(2), topLevelPermissions[1].Id)
	assert.Equal(t, int64(3), topLevelPermissions[0].Children[0].Id)
	assert.Equal(t, int64(4), topLevelPermissions[0].Children[1].Id)
	assert.Equal(t, int64(5), topLevelPermissions[1].Children[0].Id)

	// GetPermissionsFullPathByIds
	permissions1, isOk1 := tree.GetPermissionsFullPathByIds([]int64{3, 5})
	assert.Len(t, permissions1, 4)
	assert.True(t, isOk1)
	assert.Equal(t, int64(1), permissions1[0].Id)
	assert.Equal(t, int64(2), permissions1[1].Id)
	assert.Equal(t, int64(3), permissions1[2].Id)
	assert.Equal(t, int64(5), permissions1[3].Id)

	permissions2, isOk2 := tree.GetPermissionsFullPathByIds([]int64{3, 5, 6})
	assert.Len(t, permissions2, 4)
	assert.False(t, isOk2)
}
