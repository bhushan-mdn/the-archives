package models

import (
	"log"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}

// ExistFileByID checks if an file exists based on ID
func ExistFileByID(id int) (bool, error) {
	var file File
	err := DB.Select("id").Where("id = ?", id).First(&file).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if file.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistFileByID checks if an file exists based on ID
func ExistFileByName(name string) (bool, error) {
	var file File
	err := DB.Select("id").Where("name = ?", name).First(&file).Error
	log.Println("error: ", err)
	log.Println("file: ", file)
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if file.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetFilesTotal gets the total number of files based on the constraints
func GetFilesTotal(maps interface{}) (int64, error) {
	var count int64
	if err := DB.Model(&File{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetFiles gets a list of files based on paging constraints
// func GetFiles(pageNum int, pageSize int, maps interface{}) ([]*File, error) {
func GetFiles() ([]*File, error) {
	var files []*File
	// err := DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&files).Error
	// err := DB.Where(maps).Find(&files).Error
	err := DB.Find(&files).Error
	log.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return files, nil
}

// GetFile Get a single file based on ID
func GetFile(id int) (*File, error) {
	var file File
	// err := DB.Where("id = ? AND deleted_at = ? ", id, nil).First(&file).Error
	err := DB.Where("id = ?", id).First(&file).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = DB.Model(&file).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &file, nil
}

// EditFile modify a single file
func EditFile(id int, data interface{}) error {
	if err := DB.Model(&File{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// CreateFile adds a single file
func CreateFile(data map[string]interface{}) error {

	file := File{
		Name:     data["name"].(string),
		Location: data["location"].(string),
		Type:     data["type"].(string),
		Size:     data["size"].(int64),
	}

	if err := DB.Create(&file).Error; err != nil {
		return err
	}

	return nil
}

// DeleteFile delete a single file
func DeleteFile(id int) error {

	if err := DB.Where("id = ?", id).Delete(&File{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAllFile clears all files
func DeleteAllFiles() error {
	if err := DB.Unscoped().Where("deletedAt != ? ", 0).Delete(&File{}).Error; err != nil {
		return err
	}

	return nil
}
