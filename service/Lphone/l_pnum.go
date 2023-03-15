package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Lphone"
	LphoneReq "github.com/flipped-aurora/gin-vue-admin/server/model/Lphone/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type LPhoneService struct {
}

// CreateLPhone 创建LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) CreateLPhone(lPhone Lphone.LPhone) (err error) {
	err = global.GVA_DB.Create(&lPhone).Error
	return err
}

// DeleteLPhone 删除LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) DeleteLPhone(lPhone Lphone.LPhone) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Lphone.LPhone{}).Where("id = ?", lPhone.ID).Update("deleted_by", lPhone.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&lPhone).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteLPhoneByIds 批量删除LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) DeleteLPhoneByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Lphone.LPhone{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&Lphone.LPhone{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateLPhone 更新LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) UpdateLPhone(lPhone Lphone.LPhone) (err error) {
	err = global.GVA_DB.Save(&lPhone).Error
	return err
}

// GetLPhone 根据id获取LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) GetLPhone(id uint) (lPhone Lphone.LPhone, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lPhone).Error
	return
}

// GetLPhoneInfoList 分页获取LPhone记录
// Author [piexlmax](https://github.com/piexlmax)
func (lPhoneService *LPhoneService) GetLPhoneInfoList(info LphoneReq.LPhoneSearch) (list []Lphone.LPhone, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Lphone.LPhone{})
	var lPhones []Lphone.LPhone
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.IPhone != nil {
		db = db.Where("I_phone = ?", info.IPhone)
	}
	if info.LName != "" {
		db = db.Where("l_name LIKE ?", "%"+info.LName+"%")
	}
	if info.LNum != nil {
		db = db.Where("l_num = ?", info.LNum)
	}
	if info.LOther != "" {
		db = db.Where("l_other LIKE ?", "%"+info.LOther+"%")
	}
	if info.LTitle != "" {
		db = db.Where("l_title LIKE ?", "%"+info.LTitle+"%")
	}
	if info.LUnit != "" {
		db = db.Where("l_unit LIKE ?", "%"+info.LUnit+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["lNum"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&lPhones).Error
	return lPhones, total, err
}
