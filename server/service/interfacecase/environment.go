package interfacecase

import (
	"github.com/gin-gonic/gin"

	"github.com/test-instructor/yangfan/server/global"
	"github.com/test-instructor/yangfan/server/model/interfacecase"
	interfacecaseReq "github.com/test-instructor/yangfan/server/model/interfacecase/request"
)

type EnvironmentService struct {
}

func (env *EnvironmentService) CreateEnv(environment interfacecase.ApiEnv) (envs []interfacecase.ApiEnv, err error) {
	err = global.GVA_DB.Save(&environment).Error
	if err != nil {
		return nil, err
	}
	db := global.GVA_DB.Model(&interfacecase.ApiEnv{})
	err = db.Find(&envs, projectDB(db, environment.ProjectID)).Error
	return
}
func (env *EnvironmentService) DeleteEnv(environment interfacecase.ApiEnv) (err error) {
	err = global.GVA_DB.Delete(&environment).Error
	return err
}
func (env *EnvironmentService) DeleteEnvByIds(c *gin.Context) {

}
func (env *EnvironmentService) UpdateEnv(environment interfacecase.ApiEnv) (id uint, err error) {
	var oId interfacecase.Operator
	global.GVA_DB.Model(interfacecase.ApiEnv{}).Where("id = ?", environment.ID).First(&oId)
	environment.CreatedBy = oId.CreatedBy
	err = global.GVA_DB.Where(&interfacecase.ApiEnv{GVA_MODEL: global.GVA_MODEL{ID: environment.ID}}).
		Save(&environment).Error
	return environment.ID, err
}
func (env *EnvironmentService) FindEnv(id uint) (err error, environment interfacecase.ApiEnv) {
	err = global.GVA_DB.
		Where("id = ?", id).First(&environment).Error
	return
}

func (env *EnvironmentService) GetEnvList(info interfacecaseReq.EnvSearch) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiEnv{})
	var envs []interfacecase.ApiEnv
	err = db.Model(&interfacecase.ApiEnv{}).
		Find(&envs, projectDB(db, info.ProjectID)).
		Error
	return err, envs
}

//
func (env *EnvironmentService) CreateEnvVariable(envVar interfacecase.ApiEnvDetail) (err error) {
	// 创建db
	err = global.GVA_DB.Save(&envVar).Error

	return err
}

func (env *EnvironmentService) DeleteEnvVariable(environment interfacecase.ApiEnvDetail) (err error) {
	err = global.GVA_DB.Delete(&environment).Error
	return err
}

func (env *EnvironmentService) FindEnvVariable(id uint) (err error, environment interfacecase.ApiEnvDetail) {
	err = global.GVA_DB.
		Where("id = ?", id).First(&environment).Error
	return
}

func (env *EnvironmentService) GetEnvVariableList(info interfacecaseReq.EnvVariableSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&interfacecase.ApiEnvDetail{})
	var envs []interfacecase.ApiEnvDetail
	db.Where("project_id = ?", info.ProjectID)
	if info.Name != "" {
		db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Key != "" {
		db.Where("`key` LIKE ?", "%"+info.Key+"%")
	}
	err = db.Debug().Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("ID desc").Find(&envs).Error
	if err != nil {
		return
	}
	if info.ShowKey {
		var keys []string
		for i, _ := range envs {
			keys = append(keys, envs[i].Key)
		}
		return nil, keys, total
	}
	return err, envs, total
}
