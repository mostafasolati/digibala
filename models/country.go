package models

import (
  //"gorm.io/gorm"
  //"gorm.io/driver/postgres"
)

type Country struct {
	//gorm.Model
	CountryID int				//`gorm:"CountryID"`
	CountryName string	//`gorm:"CountryName"`
}
