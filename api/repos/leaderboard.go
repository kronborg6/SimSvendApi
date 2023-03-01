package repos

import (
	"errors"
	"fmt"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LeaderboardRepo struct {
	db *gorm.DB
}

func (repo *LeaderboardRepo) FindAllLeaderboards(mounthId int) ([]models.Leaderboards, error) {
	var leaderboard []models.Leaderboards
	err := repo.db.Debug().Where("month_id = ?", mounthId).Preload("Player." + clause.Associations).Preload("Month").Find(&leaderboard)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't a Leaderboard")
	}
	for i := range leaderboard {
		leaderboard[i].Player.Userinfo.Password = ""
		leaderboard[i].Player.FriendList = nil
		leaderboard[i].Player.RoleID = 0
		leaderboard[i].Player.Role.ID = 0
		leaderboard[i].Player.Role.Name = ""

	}
	// leaderboard[0].Player.Userinfo.Password = ""
	fmt.Println(len(leaderboard))
	return leaderboard, nil
}
func (repo *LeaderboardRepo) FindLeaderboards() ([]models.Leaderboards, error) {
	return nil, nil

}

func NewLeaderbaordRepo(db *gorm.DB) *LeaderboardRepo {
	return &LeaderboardRepo{db}
}
