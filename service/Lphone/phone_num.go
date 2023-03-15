package Lphone

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Lphone"
	LphoneReq "github.com/flipped-aurora/gin-vue-admin/server/model/Lphone/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type PhoneNumService struct {
}

// CreatePhoneNum 创建PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) CreatePhoneNum(phoneNum Lphone.PhoneNum) (err error) {
	err = global.GVA_DB.Create(&phoneNum).Error
	return err
}

// DeletePhoneNum 删除PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) DeletePhoneNum(phoneNum Lphone.PhoneNum) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Lphone.PhoneNum{}).Where("id = ?", phoneNum.ID).Update("deleted_by", phoneNum.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&phoneNum).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePhoneNumByIds 批量删除PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) DeletePhoneNumByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Lphone.PhoneNum{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&Lphone.PhoneNum{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePhoneNum 更新PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) UpdatePhoneNum(phoneNum Lphone.PhoneNum) (err error) {
	err = global.GVA_DB.Save(&phoneNum).Error
	return err
}

// GetPhoneNum 根据id获取PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) GetPhoneNum(id uint) (phoneNum Lphone.PhoneNum, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&phoneNum).Error
	return
}

// GetPhoneNumInfoList 分页获取PhoneNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (phoneNumService *PhoneNumService) GetPhoneNumInfoList(info LphoneReq.PhoneNumSearch) (list []Lphone.PhoneNum, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Lphone.PhoneNum{})
	var phoneNums []Lphone.PhoneNum
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Uid != "" {
		db = db.Where("uid LIKE ?", "%"+info.Uid+"%")
	}
	if info.NOther != "" {
		db = db.Where("n_other LIKE ?", "%"+info.NOther+"%")
	}
	if info.NBackup != "" {
		db = db.Where("n_backup LIKE ?", "%"+info.NBackup+"%")
	}
	if info.StartNNumber != nil && info.EndNNumber != nil {
		db = db.Where("n_number BETWEEN ? AND ? ", info.StartNNumber, info.EndNNumber)
	}
	if info.NRes != "" {
		db = db.Where("n_res LIKE ?", "%"+info.NRes+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&phoneNums).Error
	return phoneNums, total, err
}
