/*
记事本
2025.4.22 by dralee
*/
package note

type Note struct {
	Id                   uint32 // id
	Name                 string // 名称
	Content              string // 内容
	CreatorId            uint32 // 创建人
	CreationTime         int64  // 创建时间
	LastModificationTime int64  // 最后修改时间
}
