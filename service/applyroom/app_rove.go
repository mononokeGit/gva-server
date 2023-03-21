package applyroom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/applyroom"
	applyroomReq "github.com/flipped-aurora/gin-vue-admin/server/model/applyroom/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type ApproveService struct {
}

// CreateApprove 创建Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) CreateApprove(approve applyroom.Approve) (err error) {
	err = global.GVA_DB.Create(&approve).Error
	return err
}

// DeleteApprove 删除Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) DeleteApprove(approve applyroom.Approve) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&applyroom.Approve{}).Where("id = ?", approve.ID).Update("deleted_by", approve.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&approve).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteApproveByIds 批量删除Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) DeleteApproveByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&applyroom.Approve{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&applyroom.Approve{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateApprove 更新Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) UpdateApprove(approve applyroom.Approve) (err error) {
	err = global.GVA_DB.Save(&approve).Error
	return err
}

// GetApprove 根据id获取Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) GetApprove(id uint) (approve applyroom.Approve, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&approve).Error
	return
}

// GetApproveInfoList 分页获取Approve记录
// Author [piexlmax](https://github.com/piexlmax)
func (approveService *ApproveService) GetApproveInfoList(info applyroomReq.ApproveSearch) (list []applyroom.Approve, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&applyroom.Approve{})
	var approves []applyroom.Approve
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Appunit != "" {
		db = db.Where("appunit LIKE ?", "%"+info.Appunit+"%")
	}
	if info.Appleader != "" {
		db = db.Where("appleader LIKE ?", "%"+info.Appleader+"%")
	}
	if info.Approom != nil {
		db = db.Where("approom = ?", info.Approom)
	}
	if info.Appamount != nil {
		db = db.Where("appamount = ?", info.Appamount)
	}
	if info.Apptype != "" {
		db = db.Where("apptype LIKE ?", "%"+info.Apptype+"%")
	}
	if info.StartApptime != nil && info.EndApptime != nil {
		db = db.Where("apptime BETWEEN ? AND ? ", info.StartApptime, info.EndApptime)
	}
	if info.Appdevice != "" {
		db = db.Where("appdevice LIKE ?", "%"+info.Appdevice+"%")
	}
	if info.Appremarks != "" {
		db = db.Where("appremarks LIKE ?", "%"+info.Appremarks+"%")
	}
	if info.Appstatus != "" {
		db = db.Where("appstatus LIKE ?", "%"+info.Appstatus+"%")
	}
	if info.Appopinion != "" {
		db = db.Where("appopinion LIKE ?", "%"+info.Appopinion+"%")
	}
	if info.Appother != "" {
		db = db.Where("appother LIKE ?", "%"+info.Appother+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&approves).Error
	return approves, total, err
}
