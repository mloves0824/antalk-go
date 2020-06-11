package dao

import "antalk-go/internal/msg/model"

func (d *Dao) SaveMsg(msg *model.Msg) error  {
	sql := SqlSaveMsg
	_,err := d.db.Exec(sql, msg.MsgId, msg.SenderId, msg.RecvId, msg.Content)
	if err != nil {
		return err
	}
	return nil
}
