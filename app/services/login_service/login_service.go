package login_service

import (
	"bcw/app/common/message"
	"bcw/config"
	"bcw/libs"
	"bcw/models"
	"bcw/server"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct{}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserRespond struct {
	Token string `json:"token"`
}

var tm libs.TokenManager

func init() {
	tm = libs.TokenManager{SecretKey: config.GetViper().GetString("server.jwt_secret_key")}
}

func (a *AuthService) LoginAuthCheck(input LoginUserRequest) (*LoginUserRespond, error) {

	return nil, nil
}

func LoginAuthCheck(input LoginUserRequest) (*LoginUserRespond, *message.StatusCode) {
	var currentUser models.Operators

	if err := server.HttpServers.DB.Where("username = ?", input.Username).Limit(1).Find(&currentUser).Error; err != nil {
		return nil, &message.StatusWrongUsernameOrPassword
	}

	if currentUser.Username == "" {
		return nil, &message.StatusWrongUsernameOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(input.Password)); err != nil {
		return nil, &message.StatusWrongUsernameOrPassword
	}

	token, err := tm.CreateToken(currentUser.Id, currentUser.Username, currentUser.RoleId)

	if err != nil {
		log.Println(err)
		return nil, &message.StatusInternalError
	}

	var result = &LoginUserRespond{Token: token}
	log.Println(token)
	//权限
	//currentRole := bgmodel.ManagerRoles{
	//	Id: currentManager.RoleId,
	//}
	//config.DB_WEB.DB.Where(currentRole).Limit(1).Find(&currentRole)
	//permissionsId := common.StringToIntSlice(currentRole.Permissions)
	//menus := managerctr.PermissionTree(permissionsId)
	//permissionSlice, _ := managerctr.GetPermissionList(permissionsId)
	return result, nil
}
