package BasicRoom

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/BasicRoom"
	BasicRoomReq "github.com/flipped-aurora/gin-vue-admin/server/model/BasicRoom/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type BasicRoomSService struct {
}

// CreateBasicRoomS 创建BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) CreateBasicRoomS(basicroom BasicRoom.BasicRoomS) (err error) {
	err = global.GVA_DB.Create(&basicroom).Error
	return err
}

// DeleteBasicRoomS 删除BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) DeleteBasicRoomS(basicroom BasicRoom.BasicRoomS) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&BasicRoom.BasicRoomS{}).Where("id = ?", basicroom.ID).Update("deleted_by", basicroom.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&basicroom).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBasicRoomSByIds 批量删除BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) DeleteBasicRoomSByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&BasicRoom.BasicRoomS{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&BasicRoom.BasicRoomS{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBasicRoomS 更新BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) UpdateBasicRoomS(basicroom BasicRoom.BasicRoomS) (err error) {
	err = global.GVA_DB.Save(&basicroom).Error
	return err
}

// GetBasicRoomS 根据id获取BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) GetBasicRoomS(id uint) (basicroom BasicRoom.BasicRoomS, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&basicroom).Error
	return
}

// GetBasicRoomSInfoList 分页获取BasicRoomS记录
// Author [piexlmax](https://github.com/piexlmax)
func (basicroomService *BasicRoomSService) GetBasicRoomSInfoList(info BasicRoomReq.BasicRoomSSearch) (list []BasicRoom.BasicRoomS, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&BasicRoom.BasicRoomS{})
	var basicrooms []BasicRoom.BasicRoomS
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Broom != nil {
		db = db.Where("broom = ?", info.Broom)
	}
	if info.StartBamounts != nil && info.EndBamounts != nil {
		db = db.Where("bamounts BETWEEN ? AND ? ", info.StartBamounts, info.EndBamounts)
	}
	if info.Bstatus != nil {
		db = db.Where("bstatus = ?", info.Bstatus)
	}
	if info.Bmarks != "" {
		db = db.Where("bmarks LIKE ?", "%"+info.Bmarks+"%")
	}
	if info.Bother != "" {
		db = db.Where("bother LIKE ?", "%"+info.Bother+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&basicrooms).Error
	return basicrooms, total, err
}
