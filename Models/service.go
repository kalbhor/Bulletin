package Models

import (
    "../Config"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    
)

func CreateService(c *Service) error {
    if err := Config.DB.Create(&c).Error; err != nil {
        return err
    }
    return nil
}

func DeleteService(c *Service, id string) error {
 var service Service
 if err := Config.DB.Where("id = ?", id).Delete(&service).Error; err != nil {
    return err
 }
 return nil
}

func UpdateService(c *Service, id string) error {
 var service Service
 if err := Config.DB.Where("id = ?", id).First(&service).Error; err != nil {
    return err
 }
 Config.DB.Save(&c)
 return nil
}

func GetService(c *Service, id string) error {
 if err := Config.DB.Where("id = ?", id).First(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}

func GetServices(c *[]Service) error {
 if err := Config.DB.Find(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}