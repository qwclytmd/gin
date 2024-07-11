package account_service

import (
	"bcw/app/common/message"
	"bcw/app/common/public"
	"bcw/models"
	"bcw/server"
	"log"
)

type AccountService struct{}

type AccountReq struct {
	Username string `json:"username"`
	models.PageSize
}

type OperatorList struct {
	Username string `json:"username"`
}

// 账号列表
func (s AccountService) GetAccountList(input AccountReq) (*models.PaginateReq, *message.StatusCode) {

	var record []OperatorList

	db := server.HttpServers.DB.Table("operators").Joins("LEFT JOIN roles ON operators.role_id = roles.id").Where("status = ?", 1)
	if input.Username != "" {
		db.Where("username = ?", input.Username)
	}

	pageInfo := public.Paginate(db, input.Page, input.Size)

	if err := db.Scopes(pageInfo.Data).Select("username, login_ip, role_name, created_at").Order("operators.id DESC").Find(&record).Error; err != nil {
		log.Println(err)
		return nil, &message.StatusInternalError
	}

	result := &models.PaginateReq{
		Total: pageInfo.TotalCount,
		Data:  record,
	}

	return result, nil
}

//// 账号添加
//func InsertManager(c *gin.Context, input dto.AddManagerAccountRequest) error {
//	currentManager := bgmodel.Managers{}
//	err := config.DB_WEB.DB.Select("username").Where("username = ?", input.Username).Limit(1).Find(&currentManager).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		return err
//	}
//
//	if currentManager.Username != "" {
//		return errors.New("用户已存在")
//	}
//
//	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return errors.New("密码生成错误")
//	}
//
//	input.Password = string(password)
//
//	now := customtime.Time{Time: time.Now()}
//	input.RegisterAt = &now
//	input.UpdatePasswordAt = &now
//	input.OperatedId = int(c.GetInt64("operated_id"))
//
//	err = config.DB_WEB.DB.Table("managers").Create(&input).Error
//
//	if err != nil {
//		return errors.New("系统错误")
//	}
//
//	return nil
//}
//
//func UpdateManager(input dto.UpdateManagerAccountRequest) error {
//
//	now := customtime.Time{Time: time.Now()}
//	field := []string{"nickname", "role_id", "status", "google_check", "white_ip"}
//	if input.Password != "" {
//		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
//		if err != nil {
//			return errors.New("密码生成错误")
//		}
//
//		input.Password = string(password)
//		input.UpdatePasswordAt = &now
//		field = append(field, "password", "update_password_at")
//	}
//
//	if input.GoogleCheck == 1 {
//		input.UpdateGoogleAt = &now
//		field = append(field, "google_secret", "google_url", "update_google_at")
//	}
//
//	err := config.DB_WEB.DB.Table("managers").Select(field).Save(&input).Error
//
//	return err
//}
//
//func DeleteManager(id int) error {
//
//	err := config.DB_WEB.DB.Delete(&bgmodel.Managers{}, id).Error
//
//	return err
//}
//
//func UpdatePassword(input dto.UpdatePasswordRequest) error {
//
//	var currentPassword string
//	if err := config.DB_WEB.DB.Model(&bgmodel.Managers{}).Where("id = ?", input.Id).Select("password").Scan(&currentPassword).Error; err != nil {
//		return err
//	}
//
//	if err := bcrypt.CompareHashAndPassword([]byte(currentPassword), []byte(input.OldPassword)); err != nil {
//		return errors.New("旧密码错误")
//	}
//
//	//新密码
//	newPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
//	if err != nil {
//		return err
//	}
//
//	//更新密码
//	update := map[string]any{
//		"password":           newPassword,
//		"update_password_at": &customtime.Time{Time: time.Now()},
//	}
//	if err := config.DB_WEB.DB.Model(&bgmodel.Managers{}).Where("id = ?", input.Id).Updates(update).Error; err != nil {
//		return err
//	}
//
//	return err
//}
