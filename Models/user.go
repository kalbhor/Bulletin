package Models

import (
    "../Config"
    _ "github.com/jinzhu/gorm/dialects/sqlite"

)

func CreateUser(c *User) error {
    if err := Config.DB.Create(&c).Error; err != nil {
        return err
    }
    return nil
}

func DeleteUser(c *User, id string) error {
 var user User
 if err := Config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
    return err
 }
 return nil
}

func UpdateUser(c *User, id string) error {
 var user User
 if err := Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
    return err
 }
 Config.DB.Save(&c)
 return nil
}

func GetUser(c *User, id string) error {
 if err := Config.DB.Where("id = ?", id).First(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}

func GetUsers(c *[]User) error {
 if err := Config.DB.Find(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}

func AddOfferRequest(o OfferRequest, id string) error {
 var user User
 if err := Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
    return err
 }
 user.OfferRequests = append(user.OfferRequests, o)
 Config.DB.Save(&user)
 return nil
}

func AddOffer(o Offer, id string) error {
 var user User
 if err := Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
    return err
 }
 user.Offers = append(user.Offers, o)
 Config.DB.Save(&user)
 return nil
}