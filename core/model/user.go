package core_model

import (
	"html"
	"strings"
	"time"

	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ID        string    `gorm:"primaryKey;size:30"`
	Username  string    `gorm:"unique;size:20;not null"`
	Password  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	FirstName string    `gorm:"size:50;not null"`
	LastName  *string   `gorm:"size:50"`
	Status    uint      `gorm:"default:1;precision:1;size:1;not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
}

func (u *User) SaveUser(DB_SYS *gorm.DB) (*User, error) {
	tx := DB_SYS.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table("m_user").Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	uuid := shortuuid.New()
	u.ID = uuid
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	*u.LastName = html.EscapeString(strings.TrimSpace(*u.LastName))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Status = 1

	return nil
}

func (u *User) FindByUsername(DB_SYS *gorm.DB, username *string) error {
	tx := DB_SYS.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where("username = ?", *username).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ChangePassword(DB_SYS *gorm.DB) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}

	tx := DB_SYS.Session(&gorm.Session{SkipDefaultTransaction: true})
	err = tx.Table("m_user").Model(&u).Update("password", string(hashedPassword)).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func SearchActiveUsers(DB_SYS *gorm.DB, like_cond *[]string, like_value *map[string]interface{}, users *[]map[string]interface{}) error {
	tx := DB_SYS.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where(strings.Join(*like_cond, " AND "), *like_value).Find(users).Error

	if err != nil {
		return err
	}
	return nil
}
