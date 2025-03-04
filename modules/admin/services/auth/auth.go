package authService

import (
	"errors"
	"image/color"
	"m-server-api/config"
	"m-server-api/initializers"
	authDto "m-server-api/modules/admin/dtos/auth"
	"m-server-api/modules/admin/models"
	authVo "m-server-api/modules/admin/vos/auth"
	menuVo "m-server-api/modules/admin/vos/sys-menu"
	md5Encrypt "m-server-api/utils/encrypt/md5"
	"m-server-api/utils/jwt"
	"sort"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

// 生成验证码
func Captcha() (id string, b64s string, err error) {
	driverStringConfig := &base64Captcha.DriverString{
		// 验证码图片的高度
		Height: 50,
		// 验证码图片的宽度
		Width: 120,
		// 验证码图片中随机噪点的数量
		NoiseCount: 2,
		// 控制显示在验证码图片中的线条的选项
		ShowLineOptions: 2 | 4,
		// 验证码的长度，即验证码中字符的数量
		Length: 4,
		// 验证码的字符源，用于生成验证码的字符。在这个例子中，使用数字和小写字母作为字符源
		Source: "1234567890abcdefghijklmnopqrstuvwxyz",
		BgColor: &color.RGBA{
			//验证码图片的背景颜色
			R: 100,
			G: 200,
			B: 100,
			A: 125,
		},
		//用于绘制验证码文本的字体文件。使用"wqy-microhei.ttc"的字体文件
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 将driverString中指定的字体文件转换为驱动程序所需的字体格式,这个步骤是为了将字体文件转换为正确的格式，以便在生成验证码时使用正确的字体
	driver := driverStringConfig.ConvertFonts()
	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}
	b64s = item.EncodeB64string()
	err = initializers.RDB.Set(initializers.Ctx, "captcha:"+id, answer, 2*time.Minute).Err()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil
}

func Login(dto authDto.LoginDto) (*authVo.LoginVo, error) {
	// 判断验证码
	captchaVal := initializers.RDB.Get(initializers.Ctx, "captcha:"+dto.CaptchaId)
	if captchaVal.Val() == "" || captchaVal.Val() != dto.Captcha {
		return nil, errors.New("验证码错误")
	}
	// 加密密码比对
	password := md5Encrypt.Encode(dto.Password)

	var user models.SysUser
	err := initializers.DB.Where("account = ? AND password = ?", dto.Account, password).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	var tenant models.SysTenant
	err = initializers.DB.First(&tenant, user.TenantId).Error
	if err != nil {
		return nil, err
	}
	// 判断租户状态
	if *tenant.BaseModelNoTenant.Status != 1 {
		return nil, errors.New("租户已停用")
	}

	// 判断用户状态
	if *user.BaseModel.Status != 1 {
		return nil, errors.New("账号已停用")
	}

	var token string
	token, err = jwt.GenerateToken(jwt.SessionUserInfo{
		Id:       user.BaseModel.ID,
		TenantId: *user.BaseModel.TenantId,
		Platform: config.ADMIN,
	}, time.Duration(cast.ToInt(config.Get().Jwt.Expire))*time.Hour)
	if err != nil {
		return nil, err
	}

	// 销毁验证码
	initializers.RDB.Del(initializers.Ctx, "captcha:"+dto.CaptchaId)

	return &authVo.LoginVo{
		Token: token,
	}, nil
}

// 用户信息
func UserInfo(userId int64) (*authVo.UserInfoVo, error) {
	var res authVo.UserInfoVo
	err := initializers.DB.Model(&models.SysUser{}).Select("sys_user.*,sys_tenant.tenant_name").Joins("LEFT JOIN sys_tenant ON sys_tenant.id = sys_user.tenant_id").Where("sys_user.id = ?", userId).First(&res).Error
	return &res, err
}

// 排序树形菜单
func sortMenuTree(nodes []*menuVo.MenuTree) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Sort < nodes[j].Sort
	})
	for _, node := range nodes {
		if len(node.Children) > 0 {
			sortMenuTree(node.Children)
		}
	}
}

// 用户菜单树
func MenuTree(userId int64) ([]*menuVo.MenuTree, error) {
	var res []models.SysMenu
	err := initializers.DB.Raw(`
		SELECT
			sys_menu.*
		FROM
			sys_menu
		LEFT JOIN sys_role_menu ON sys_menu.id = sys_role_menu.menu_id
		LEFT JOIN sys_role ON sys_role.id = sys_role_menu.role_id
		LEFT JOIN sys_user_role ON sys_user_role.role_id = sys_role.id
		LEFT JOIN sys_user ON sys_user.id = sys_user_role.user_id 
		WHERE
			sys_user.id = ? 
		AND sys_menu.type IN ( 1, 2 ) 
		AND sys_menu.status = 1
		AND sys_role.status = 1
	`, userId).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	menuMap := make(map[int64]*menuVo.MenuTree)
	var roots []*menuVo.MenuTree
	for _, m := range res {
		node := &menuVo.MenuTree{
			ID:       m.ID,
			Name:     m.Name,
			Sort:     m.Sort,
			Type:     m.Type,
			Icon:     m.Icon,
			Path:     m.Path,
			Alias:    m.Alias,
			Status:   m.Status,
			ParentId: m.ParentId,
			Keep:     m.Keep,
			Children: []*menuVo.MenuTree{},
		}
		menuMap[m.ID] = node
	}
	// 构建树形结构
	for _, node := range menuMap {
		if node.ParentId == nil {
			// 一级菜单
			roots = append(roots, node)
		} else {
			// 找到父级，将当前节点挂载到父级的 Children 中
			if parent, ok := menuMap[*node.ParentId]; ok {
				parent.Children = append(parent.Children, node)
			}
		}
	}
	sortMenuTree(roots)
	return roots, nil
}

func ChangePassword(dto authDto.ChangePasswordDto, userId int64) error {
	var user models.SysUser
	err := initializers.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return err
	}
	if user.Password != md5Encrypt.Encode(dto.Password) {
		return errors.New("原密码不正确")
	}
	err = initializers.DB.Model(&models.SysUser{}).Where("id = ?", userId).Update("password", md5Encrypt.Encode(dto.NewPassword)).Error
	return err
}
