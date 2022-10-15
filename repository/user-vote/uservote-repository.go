package uservoterepository

import (
	"errors"
	"log"
	"strconv"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repositoryUserVote struct {
}

func New() repository_intf.UserVoteRepository {
	return &repositoryUserVote{}
}

func (r *repositoryUserVote) ChooseVotes(c *gin.Context) error {
	var userId uint = 2 // dummy dulu, nanti isi ini pake jwt

	query := c.Query("isVoting")
	isVoting, err := strconv.ParseBool(query)

	if err != nil {
		log.Fatal(err)
	}

	paramsId, err := strconv.ParseInt(c.Params.ByName("reportId"), 32, 32)
	if err != nil {
		return errors.New("failed to convert params")
	}
	reportId := uint(paramsId)

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	user := entity.UserVote{}
	dbRresult := db.Where("user_id = ? AND report_id = ?", userId, reportId).First(&user)
	if errors.Is(dbRresult.Error, gorm.ErrRecordNotFound) {
		vote := entity.UserVote{
			UserID:   userId,
			ReportID: reportId,
			IsUpVote: &isVoting,
		}
		if err := db.Create(&vote).Error; err != nil {
			return errors.New("failed to create vote")
		}
	} else {
		if *user.IsUpVote && isVoting {
			// update uservotes ke nil
			err := db.First(&user).Error
			if err != nil {
				return errors.New("failed convert from vote true to nil")
			}
			user.IsUpVote = nil

			err = db.Save(&user).Error
			if err != nil {
				return errors.New("failed save from vote true to nil")
			}

		} else if *user.IsUpVote && !isVoting {
			// update uservotes ke false
			err := db.First(&user).Error
			if err != nil {
				return errors.New("failed convert from vote true to nil")
			}
			*user.IsUpVote = false

			err = db.Save(&user).Error
			if err != nil {
				return errors.New("failed save from vote true to nil")
			}

		} else if !*user.IsUpVote && !isVoting {
			// update uservotes jadi nill
			err := db.First(&user).Error
			if err != nil {
				return errors.New("failed convert from vote false to nil")
			}

			user.IsUpVote = nil
			err = db.Save(&user).Error
			if err != nil {
				return errors.New("failed save from vote true to nil")
			}

		} else if !*user.IsUpVote && isVoting {
			// update uservotes ke true
			err := db.First(&user).Error
			if err != nil {
				return errors.New("failed convert from vote false to true")
			}
			*user.IsUpVote = true
			db.Save(&user)
			if err != nil {
				return errors.New("failed convert from vote true to nil")
			}
		}
	}

	return nil
}
