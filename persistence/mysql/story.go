package mysql

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dstory "github.com/rice9547/hakka_story/domain/story"
	"github.com/rice9547/hakka_story/lib/errors"
)

type StoryRepository struct {
	DB *gorm.DB
}

func NewStory(client *Client) dstory.Repository {
	return &StoryRepository{DB: client.DB()}
}

func (r *StoryRepository) Save(s *dstory.Story) error {
	return r.DB.Save(s).Error
}

func (r *StoryRepository) List() ([]dstory.Story, error) {
	stories := make([]dstory.Story, 0)
	err := r.DB.Preload(clause.Associations).Find(&stories).Error
	return stories, err
}

func (r *StoryRepository) FilterByCategories(categoryNames []string) ([]dstory.Story, error) {
	stories := make([]dstory.Story, 0)
	err := r.DB.Preload(clause.Associations).
		Joins("JOIN story_to_category ON story_to_category.story_id = stories.id").
		Joins("JOIN categories ON story_to_category.category_id = categories.id").
		Where("categories.name IN (?)", categoryNames).
		Find(&stories).Error
	return stories, err
}

func (r *StoryRepository) GetByID(id uint64) (*dstory.Story, error) {
	story := &dstory.Story{}
	err := r.DB.
		Model(&story).
		Preload(clause.Associations).
		Preload("Pages.AudioFiles").
		Preload("Categories", func(db *gorm.DB) *gorm.DB {
			return db.Joins("JOIN story_to_category ON story_to_category.category_id = categories.id").
				Order("story_to_category.id ASC")
		}).
		First(&story, id).Error
	return story, err
}

func (r *StoryRepository) UpdateByID(id uint64, story *dstory.Story) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&dstory.StoryPage{}, "story_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&dstory.StoryToCategory{}, "story_id = ?", id).Error; err != nil {
			return err
		}

		story.ID = id

		return tx.Save(story).Error
	})
}

func (r *StoryRepository) DeleteByID(id uint64) error {
	err := r.DB.Delete(&dstory.Story{}, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.ErrStoryNotFound
	}
	return err
}
