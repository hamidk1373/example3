package users

import (
	"hamid/example3/settings"
)

func getAllFromDB() ([]User, error) {
	db, err := settings.GetDB()
	if err != nil {
		return nil, err
	}
	// users := []User{}
	var users []User
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getOneFromDB(ID uint) (*User, error) {
	db, err := settings.GetDB()
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) createOneToDB() error {
	db, err := settings.GetDB()
	if err != nil {
		return err
	}
	err = db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) updateOneToDB() error {
	db, err := settings.GetDB()
	if err != nil {
		return err
	}
	return db.Save(&user).Error
}

func deleteOneFromDB(ID uint) error {
	db, err := settings.GetDB()
	if err != nil {
		return err
	}
	return db.Delete(&User{}, ID).Error
}

func getOneByEmailFromDB(Email string) (*User, error) {
	db, err := settings.GetDB()
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = db.Where("Email = ?", Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
