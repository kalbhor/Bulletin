package Models

import (
    "../Config"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

)

func CreateCategory(c *Category) error {
    if err := Config.DB.Create(&c).Error; err != nil {
        return err
    }
    return nil
}

func DeleteCategory(c *Category, id string) error {
 var category Category
 if err := Config.DB.Where("id = ?", id).Delete(&category).Error; err != nil {
    return err
 }
 return nil
}

func UpdateCategory(c *Category, id string) error {
 var category Category
 if err := Config.DB.Where("id = ?", id).First(&category).Error; err != nil {
    return err
 }
 Config.DB.Save(&c)
 return nil
}

func GetCategory(c *Category, id string) error {
 if err := Config.DB.Where("id = ?", id).First(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}

func GetCategories(c *[]Category) error {
 if err := Config.DB.Find(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}