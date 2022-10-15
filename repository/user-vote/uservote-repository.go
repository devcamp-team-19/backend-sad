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
			IsUpVote: nil,
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
			db.Save(&user)

		} else if *user.IsUpVote && !isVoting {
			// update uservotes ke false
			err := db.First(&user).Error
			*user.IsUpVote = false
			err = db.Save(&user).Error

		} else if !*user.IsUpVote && !isVoting {
			// update uservotes jadi nill
			db.First(&user)
			user.IsUpVote = nil
			db.Save(&user)

		} else if !*user.IsUpVote && isVoting {
			// update uservotes ke true
			db.First(&user)
			*user.IsUpVote = true
			db.Save(&user)
		}
	}

	return nil
}
