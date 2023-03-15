package meetcalendar

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/meetcalendar"
	meetcalendarReq "github.com/flipped-aurora/gin-vue-admin/server/model/meetcalendar/request"
	"gorm.io/gorm"
)

type MeetcalendarService struct {
}

// CreateMeetcalendar 创建Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) CreateMeetcalendar(meetcalendar meetcalendar.Meetcalendar) (err error) {
	err = global.GVA_DB.Create(&meetcalendar).Error
	return err
}

// DeleteMeetcalendar 删除Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) DeleteMeetcalendar(xmeetcalendar meetcalendar.Meetcalendar) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&meetcalendar.Meetcalendar{}).Where("id = ?", xmeetcalendar.ID).Update("deleted_by", xmeetcalendar.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&xmeetcalendar).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMeetcalendarByIds 批量删除Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) DeleteMeetcalendarByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&meetcalendar.Meetcalendar{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&meetcalendar.Meetcalendar{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMeetcalendar 更新Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) UpdateMeetcalendar(meetcalendar meetcalendar.Meetcalendar) (err error) {
	err = global.GVA_DB.Save(&meetcalendar).Error
	return err
}

// GetMeetcalendar 根据id获取Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) GetMeetcalendar(id uint) (meetcalendar meetcalendar.Meetcalendar, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&meetcalendar).Error
	return
}

// GetMeetcalendarInfoList 分页获取Meetcalendar记录
// Author [piexlmax](https://github.com/piexlmax)
func (meetcalendarService *MeetcalendarService) GetMeetcalendarInfoList(info meetcalendarReq.MeetcalendarSearch) (list []meetcalendar.Meetcalendar, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&meetcalendar.Meetcalendar{})
	var meetcalendars []meetcalendar.Meetcalendar
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartMDate != nil && info.EndMDate != nil {
		db = db.Where("m_date BETWEEN ? AND ? ", info.StartMDate, info.EndMDate)
	}
	if info.MRoom != nil {
		db = db.Where("m_room = ?", info.MRoom)
	}
	if info.MType != "" {
		db = db.Where("m_type LIKE ?", "%"+info.MType+"%")
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.StartStime != nil && info.EndStime != nil {
		db = db.Where("stime BETWEEN ? AND ? ", info.StartStime, info.EndStime)
	}
	if info.StartPtime != nil && info.EndPtime != nil {
		db = db.Where("ptime BETWEEN ? AND ? ", info.StartPtime, info.EndPtime)
	}
	if info.Leader != "" {
		db = db.Where("leader LIKE ?", "%"+info.Leader+"%")
	}
	if info.Mleader != "" {
		db = db.Where("mleader LIKE ?", "%"+info.Mleader+"%")
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.Level != "" {
		db = db.Where("level LIKE ?", "%"+info.Level+"%")
	}
	if info.Resposen != "" {
		db = db.Where("resposen LIKE ?", "%"+info.Resposen+"%")
	}
	if info.Route != "" {
		db = db.Where("route LIKE ?", "%"+info.Route+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&meetcalendars).Error
	return meetcalendars, total, err
}
