package repos

import (
	"errors"
	"fmt"

	"github.com/kronborg6/SimSvendApi/api/models"
	"gorm.io/gorm"
)

type MatchRepo struct {
	db *gorm.DB
}

func (repo *MatchRepo) FindMyGame(gameID int) ([]models.Match, error) {
	var game []models.Match

	err := repo.db.Where("id = ?", gameID).Find(&game)
	if err.Error != nil {
		return nil, errors.New("can't find game")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find game")
	}

	return game, nil
}
func (repo *MatchRepo) FindMyGames(userID int) ([]models.Match, error) {
	var games []models.Match
	var game []models.Match

	err := repo.db.Where("team_a_player_a = ?", userID).Or("team_a_player_b = ?", userID).Or("team_b_player_a = ?", userID).Or("team_b_player_b = ?", userID).Find(&games)
	if err.Error != nil {
		return nil, errors.New("can't find games")
	}
	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find games")
	}
	for i := range games {
		if !games[i].Don {
			game = append(games, game...)
			fmt.Println("gg")
		}
	}
	return game, nil
}
func (repo *MatchRepo) GameHistory(userID int) ([]models.Match, error) {
	var games []models.Match

	err := repo.db.Where("team_a_player_a = ? AND don = true", userID).Or("team_a_player_b = ?", userID).Or("team_b_player_a = ?", userID).Or("team_b_player_b = ?", userID).Preload("Result").Preload("User1").Preload("User2").Preload("User3").Preload("User4").Preload("Place").Preload("Court").Find(&games)
	if err.Error != nil {
		return nil, errors.New("can't find games")
	}

	if err.RowsAffected <= 0 {
		return nil, errors.New("can't find games")
	}
	for i := range games {
		games[i].User1.Password = ""
		games[i].User2.Password = ""
		games[i].User3.Password = ""
		games[i].User4.Password = ""
	}
	return games, nil
}
func (repo *MatchRepo) FindGame(gameID int) ([]models.Match, error) {
	var game []models.Match

	if err := repo.db.Debug().Where("id = ?", gameID).Preload("Result").Preload("User1").Preload("User2").Preload("User3").Preload("User4").Preload("Place").Preload("Court").Find(&game).Error; err != nil {
		return nil, err
	}
	if game == nil {
		return nil, errors.New("can't find the game")
	}
	game[0].User1.Password = ""
	game[0].User2.Password = ""
	game[0].User3.Password = ""
	game[0].User4.Password = ""
	// for i := range game {
	// 	game[i].User1.Password = ""
	// 	game[i].User2.Password = ""
	// 	game[i].User3.Password = ""
	// 	game[i].User4.Password = ""
	// }

	return game, nil
}

/*
func (repo *MatchRepo) SetGameScore(Score models.Results) (models.Results, error) {
	var game models.Results
	var match models.Match
	if err := repo.db.Debug().Where("id = ?", Score.MatchID).Find(&match).Error; err != nil {
		return game, err
	}
	if match.Don {
		return game, errors.New("game is allready don")
	}
	if err := repo.db.Debug().Model(&Score).Where("match_id = ?", Score.MatchID).Updates(&Score).Error; err != nil {
		return game, err
	}
	match.Don = true
	if err := repo.db.Debug().Model(&match).Where("id = ?", Score.MatchID).Updates(&match).Error; err != nil {
		return game, err
	}
	return game, nil
}
*/
/*
	 func (repo *MatchRepo) NewCasualGame(game models.Matchs) (models.Matchs, error) {
		// var game models.Matchs
		if err := repo.db.Create(&game).Error; err != nil {
			return game, err
		}
		return game, nil
	}

	func (repo *MatchRepo) AddResualtCasualGame(game models.Matchs) (models.Matchs, error) {
		if err := repo.db.Model(&game).Where("id = ?", game.ID).Updates(&game).Error; err != nil {
			return game, nil
		}
		return game, nil

}
*/
func NewMatchRepo(db *gorm.DB) *MatchRepo {
	return &MatchRepo{db}
}
