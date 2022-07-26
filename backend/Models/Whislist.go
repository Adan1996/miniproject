package Models

import (
	"example.com/m/Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllWhislist(whislist *[]Whislist) (err error) {
	if err = Config.DB.Find(whislist).Error; err != nil {
		return err
	}
	return nil
}
func CreateWhislist(whislist *Whislist) (err error) {
	if err = Config.DB.Create(whislist).Error; err != nil {
		return err
	}
	return nil
}
