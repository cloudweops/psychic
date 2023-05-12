package mysql

import (
	"github.com/CloudWeOps/psychic/cmd/user/pkg/md5"
	"github.com/CloudWeOps/psychic/share/consts"
	"github.com/CloudWeOps/psychic/share/errno"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"primarykey"`
	PhoneNumber  string
	AvatarBlobId string
	Username     string `gorm:"type:varchar(40)"`
	OpenID       string `gorm:"column:openid;type:varchar(100);uniqueIndex"`
	Balance      int32  `gorm:"column:balance"`
	Deleted      gorm.DeletedAt
}

// BeforeCreate uses snowflake to generate an ID.
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if u.ID != "" {
		return nil
	}
	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Fatalf("generate id failed: %s", err.Error())
	}
	u.ID = sf.Generate().String()
	return nil
}

type UserManager struct {
	salt string
	db   *gorm.DB
}

// NewUserManager creates a mysql manager.
func NewUserManager(db *gorm.DB, salt string) *UserManager {
	m := db.Migrator()
	if !m.HasTable(&User{}) {
		if err := m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	return &UserManager{
		db:   db,
		salt: salt,
	}
}

func (m *UserManager) CreateUser(user *User) (*User, error) {
	if user.OpenID == "" {
		return nil, errno.UserSrvErr.WithMessage("openId not set")
	}
	if _, err := m.GetUserByOpenId(user.OpenID); err == nil {
		return nil, errno.RecordAlreadyExist
	} else if err != errno.RecordNotFound {
		return nil, err
	}
	user.OpenID = md5.Md5Crypt(user.OpenID, m.salt)
	err := m.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (m *UserManager) GetUserByOpenId(openID string) (*User, error) {
	var user User
	cryOpenID := md5.Md5Crypt(openID, m.salt)
	err := m.db.Where(&User{OpenID: cryOpenID}).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errno.RecordNotFound
		} else {
			return nil, err
		}
	}
	return &user, nil
}
