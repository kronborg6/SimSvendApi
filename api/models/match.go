package models

import "time"

type Matchs struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	PlayTime     time.Time `json:"play_time" gorm:"not null"`
	PlaceId      int64
	Place        Clubs `gorm:"foreignKey:PlaceId"`
	CourtId      int64
	Court        ClubCourts `gorm:"foreignKey:CourtId"`
	TeamAPlayerA int64      `json:"team_a_player_a"`
	TeamAPlayerB int64      `json:"team_a_player_b"`
	TeamBPlayerA int64      `json:"team_b_player_a"`
	TeamBPlayerB int64      `json:"team_b_player_b"`
	User1        UserInfo   `gorm:"foreignKey:TeamAPlayerA"`
	User2        UserInfo   `gorm:"foreignKey:TeamAPlayerB"`
	User3        UserInfo   `gorm:"foreignKey:TeamBPlayerA"`
	User4        UserInfo   `gorm:"foreignKey:TeamBPlayerB"`
	ResultsId    int64
	Result       Results `gorm:"foreignKey:ResultsId"`
}

type Results struct {
	Id int64 `json:"id" gorm:"primaryKey"`
}

/*
	 type MatchHistoryTeams struct {
	}
*/
