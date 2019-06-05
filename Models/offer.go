package Models

import (
    "../Config"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    
)

func CreateOffer(c *Offer) error {
    if err := Config.DB.Create(&c).Error; err != nil {
        return err
    }
    return nil
}

func DeleteOffer(c *Offer, id string) error {
 var offer Offer
 if err := Config.DB.Where("id = ?", id).Delete(&offer).Error; err != nil {
    return err
 } // check
 return nil
}

func UpdateOffer(c *Offer, id string) error {
 var offer Offer
 if err := Config.DB.Where("id = ?", id).First(&offer).Error; err != nil {
    return err
 }
 Config.DB.Save(&c)
 return nil
}

func GetOffer(c *Offer, id string) error {
 if err := Config.DB.Where("id = ?", id).First(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}

func GetOffers(c *[]Offer) error {
 if err := Config.DB.Find(&c).Error; err != nil {
    return err
 } else {
    return nil
 }
}