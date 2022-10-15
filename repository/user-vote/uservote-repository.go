package uservoterepository

import (
	"errors"
	"fmt"
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

	userVote := entity.UserVote{}
	dbRresult := db.Where("user_id = ? AND report_id = ?", userId, reportId).First(&userVote)

	// when user never vote the report that we addresed
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
		// nil to isVoting
		if userVote.IsUpVote == nil {
			fmt.Println("Hello")
			if !isVoting {
				db.Raw(` 
				UPDATE user_votes
				SET is_up_vote = FALSE
				WHERE user_id = ?;`, userId).Scan(&userVote)
			} else {
				db.Raw(` 
				UPDATE user_votes
				SET is_up_vote = TRUE
				WHERE user_id = ?;`, userId).Scan(&userVote)
			}
		} else {
			if *userVote.IsUpVote && isVoting {
				// update uservotes ke nil
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}

				userVote.IsUpVote = nil

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if *userVote.IsUpVote && !isVoting {
				// update uservotes ke false
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}

				*userVote.IsUpVote = false

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if !*userVote.IsUpVote && !isVoting {
				// update uservotes jadi nill
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote false to nil")
				}

				userVote.IsUpVote = nil

				err = db.Save(&userVote).Error
				if err != nil {
					return errors.New("failed save from vote true to nil")
				}

			} else if !*userVote.IsUpVote && isVoting {
				// update uservotes ke true
				err := db.First(&userVote).Error
				if err != nil {
					return errors.New("failed convert from vote false to true")
				}

				*userVote.IsUpVote = true

				db.Save(&userVote)
				if err != nil {
					return errors.New("failed convert from vote true to nil")
				}
			}
		}
	}

	return nil
}