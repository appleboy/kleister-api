package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Minecrafts is simply a collection of minecraft structs.
type Minecrafts []*Minecraft

// Filter searches for a name substring and returns a new collection.
func (u *Minecrafts) Filter(term string) *Minecrafts {
	res := Minecrafts{}

	for _, record := range *u {
		if strings.Contains(record.Name, term) {
			res = append(res, record)
		}
	}

	return &res
}

// Minecraft represents a minecraft model definition.
type Minecraft struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Slug      string    `json:"slug" sql:"unique_index"`
	Name      string    `json:"version" sql:"unique_index"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Builds    Builds    `json:"-"`
}

// BeforeSave invokes required actions before persisting.
func (u *Minecraft) BeforeSave(db *gorm.DB) (err error) {
	if u.Slug == "" {
		for i := 0; true; i++ {
			hash := md5.Sum([]byte(fmt.Sprintf("%s-%d", u.Name, i)))
			u.Slug = hex.EncodeToString(hash[:])

			notFound := db.Where(
				"slug = ?",
				u.Slug,
			).Not(
				"id",
				u.ID,
			).First(
				&Build{},
			).RecordNotFound()

			if notFound {
				break
			}
		}
	}

	return nil
}

// Validate does some validation to be able to store the record.
func (u *Minecraft) Validate(db *gorm.DB) {

}
