package entityManager

import (
	"fmt"
)

type EntityType int

const (
	Created EntityType = 0
	Founded EntityType = 1
)

type Entity struct {
	entityType            EntityType
	tableName             string
	isWaitSaveBasic       bool
	isWaitSaveAssociation bool

	basicCandidate       []fieldInfo
	associationCandidate []fieldInfo
	joinTableInfos       []joinTableInfo

	GoenId int `db:"goen_id"`
}

// 填充除了数据库字段外的值
func (p *Entity) initEntity(entityType EntityType, tableName string, goenId int) {
	if goenId != 0 {
		p.GoenId = goenId
	}
	p.entityType = entityType
	p.tableName = tableName
	//p.basicCandidate = map[string]any{}
	if entityType == Created {
		p.tryWaitSave(false)
	}
}

func (p *Entity) tryWaitSave(isAssociation bool) {
	if isAssociation {
		if !p.isWaitSaveAssociation {
			p.isWaitSaveAssociation = true
			Manager.addWaitSave(isAssociation, func() error {
				return p.saveAssociation()
			})
		}
	} else {
		if !p.isWaitSaveBasic {
			p.isWaitSaveBasic = true
			Manager.addWaitSave(isAssociation, func() error {
				return p.saveBasic()
			})
		}
	}

}

func (p *Entity) addBasicField(field string, value any) {
	p.basicCandidate = append(p.basicCandidate, fieldInfo{field: field, value: value})
	p.tryWaitSave(false)
}
func (p *Entity) addAssociationField(field string, value any) {
	p.associationCandidate = append(p.associationCandidate, fieldInfo{field: field, value: value})
	p.tryWaitSave(true)
}

type fieldInfo struct {
	field string
	value any
}
type joinTableInfo struct {
	isInsert        bool
	associationName string
	//otherTableName  string
	otherGoenId int
}

func (p *Entity) addJoinTableInsert(isInsert bool, associationName string, otherTableName string, otherGoenId int) {
	p.joinTableInfos = append(p.joinTableInfos, joinTableInfo{
		isInsert:        isInsert,
		associationName: associationName,
		//otherTableName:  otherTableName,
		otherGoenId: otherGoenId,
	})
	p.tryWaitSave(true)
}

func (p *Entity) getInsertQuery(isAssociation bool) string {
	var setCandidate []fieldInfo
	if isAssociation {
		setCandidate = p.associationCandidate
	} else {
		setCandidate = p.basicCandidate
	}

	query := fmt.Sprintf("insert into %s(goen_id, ", p.tableName)
	candidate := setCandidate[0 : len(setCandidate)-1]
	for _, field := range candidate {
		query += fmt.Sprintf("%s, ", field.field)
	}
	query += fmt.Sprintf("%s) values(:goen_id, ", setCandidate[len(setCandidate)-1].field)
	candidate = setCandidate[0 : len(setCandidate)-1]
	for _, field := range candidate {
		query += fmt.Sprintf(":%s ,", field.field)
	}
	query += fmt.Sprintf(":%s )", setCandidate[len(setCandidate)-1].field)
	print(query)
	return query
}

func (p *Entity) getUpdateQuery(isAssociation bool) string {
	var setCandidate []fieldInfo
	if isAssociation {
		setCandidate = p.associationCandidate
	} else {
		setCandidate = p.basicCandidate
	}

	query := fmt.Sprintf("update %s set ", p.tableName)
	if len(setCandidate) > 1 {
		candidate := setCandidate[0 : len(setCandidate)-1]
		for _, info := range candidate {
			query += fmt.Sprintf("%s= :%s,", info.field, info.field)
		}
	}
	query += fmt.Sprintf("%s = :%s", setCandidate[len(setCandidate)-1].field, setCandidate[len(setCandidate)-1].field)
	query += fmt.Sprintf(" where goen_id = :goen_id")
	print(query)
	return query
}

func (p *Entity) getSelectJoinTableQuery(associationName string, otherTableName string) string {
	return fmt.Sprintf("select tmp.* from %s_%s as ass, %s as tmp where ass.goen_id_1 = ? and ass.goen_id_2 = tmp.goen_id ",
		otherTableName, associationName, otherTableName)
}

// 基本类型变化的保存
func (p *Entity) saveBasic() error {
	var query string
	mp := map[string]any{}
	mp["goen_id"] = p.GoenId
	for _, fieldInto := range p.basicCandidate {
		mp[fieldInto.field] = fieldInto.value
	}
	if p.entityType == Founded {
		query = p.getUpdateQuery(false)
	} else {
		query = p.getInsertQuery(false)
		p.entityType = Founded
	}

	_, err := Db.NamedExec(query, mp)
	if err != nil {
		return err
	}
	p.basicCandidate = nil
	return nil
}

func (p *Entity) saveAssociation() error {
	var query string
	if len(p.associationCandidate) > 0 {
		mp := map[string]any{}
		mp["goen_id"] = p.GoenId
		for _, fieldInto := range p.associationCandidate {
			mp[fieldInto.field] = fieldInto.value
		}

		if p.entityType == Founded {
			query = p.getUpdateQuery(true)
		} else {
			query = p.getInsertQuery(true)
			p.entityType = Founded
		}

		_, err := Db.NamedExec(query, mp)
		if err != nil {
			return err
		}
		p.associationCandidate = nil
	}

	for _, info := range p.joinTableInfos {
		if info.isInsert {
			query = fmt.Sprintf("insert into %s_%s (goen_id_1, goen_id_2) values (?, ?)",
				p.tableName, info.associationName)
		} else {
			query = fmt.Sprintf("delete from %s_%s where goen_id_1 = ? and goen_id_2 = ?",
				p.tableName, info.associationName)
		}
		_, err := Db.Exec(query, p.GoenId, info.otherGoenId)
		if err != nil {
			return err
		}
	}
	p.joinTableInfos = nil

	p.isWaitSaveAssociation = false
	return nil
}
