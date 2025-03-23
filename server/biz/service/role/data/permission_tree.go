package data

import "github.com/moumou/server/biz/model"

// PermissionTree 权限树的内存结构，支持通过id查询子节点
type PermissionTree struct {
	permissions         []*model.Permission
	topLevelPermissions []*model.Permission
	permissionMapById   map[int64]*model.Permission
}

func BuildPermissionTree(permissions []*model.Permission) *PermissionTree {
	tree := &PermissionTree{
		permissions:         permissions,
		topLevelPermissions: make([]*model.Permission, 0, len(permissions)/3),
		permissionMapById:   make(map[int64]*model.Permission, len(permissions)),
	}
	for _, permission := range permissions {
		tree.permissionMapById[permission.Id] = permission
	}
	for _, permission := range permissions {
		if permission.Pid == 0 {
			// 根节点
			tree.topLevelPermissions = append(tree.topLevelPermissions, permission)
			continue
		}
		pid := permission.Pid
		tree.permissionMapById[pid].Children = append(tree.permissionMapById[pid].Children, permission)
		permission.Parent = tree.permissionMapById[pid]
	}

	return tree
}

func (p *PermissionTree) GetTopLevelPermissions() []*model.Permission {
	return p.topLevelPermissions
}

func (p *PermissionTree) GetPermissionById(permissionId int64) *model.Permission {
	return p.permissionMapById[permissionId]
}

// GetPermissionsFullPathByIds 根据permissionIds获取对应的节点以及其父亲到top层级的所有节点
func (p *PermissionTree) GetPermissionsFullPathByIds(permissionIds []int64) ([]*model.Permission, bool) {
	parentPathPermissionMap := make(map[int64]bool, len(permissionIds)*2)
	var isOk = true
	for _, permissionId := range permissionIds {
		if !p.markFullParentPathByPermissionId(parentPathPermissionMap, permissionId) {
			isOk = false
		}
	}
	childrenPathPermissionMap := make(map[int64]bool, len(permissionIds))
	for _, permissionId := range permissionIds {
		p.markFullChildrenPathByPermissionId(childrenPathPermissionMap, permissionId)
	}
	result := make([]*model.Permission, 0, len(permissionIds)*2)
	for _, permission := range p.permissions {
		if parentPathPermissionMap[permission.Id] || childrenPathPermissionMap[permission.Id] {
			result = append(result, permission)
		}
	}

	return result, isOk
}

func (p *PermissionTree) markFullParentPathByPermissionId(fullPathPermissionMap map[int64]bool, permissionId int64) bool {
	if fullPathPermissionMap[permissionId] {
		return true
	}
	permission := p.GetPermissionById(permissionId)
	if permission == nil {
		return false
	}
	fullPathPermissionMap[permissionId] = true
	if permission.Pid != 0 {
		return p.markFullParentPathByPermissionId(fullPathPermissionMap, permission.Pid)
	}
	return true
}

func (p *PermissionTree) markFullChildrenPathByPermissionId(childrenPathPermissionMap map[int64]bool, permissionId int64) {
	if childrenPathPermissionMap[permissionId] {
		return
	}
	permission := p.GetPermissionById(permissionId)
	if permission == nil {
		return
	}
	childrenPathPermissionMap[permissionId] = true
	if len(permission.Children) > 0 {
		for _, childPermission := range permission.Children {
			p.markFullChildrenPathByPermissionId(childrenPathPermissionMap, childPermission.Id)
		}
	}

}
