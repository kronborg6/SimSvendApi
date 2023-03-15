package models

import "time"

type Match struct {
	ID           int
	PlayTime     time.Time  `json:"play_time" gorm:"not null"`
	PlaceId      int        `json:"place_id"`
	Place        Club       `gorm:"foreignKey:PlaceId"`
	CourtId      int        `json:"court_id"`
	Court        ClubCourts `gorm:"foreignKey:CourtId"`
	Don          bool       `json:"don" gorm:"NOT NULL; default:false"`
	Comp         bool       `json:"comp" gorm:"NOT NULL; default:false"`
	TeamAPlayerA int        `json:"team_a_player_a" gorm:"NOT NULL"`
	TeamAPlayerB int        `json:"team_a_player_b"`
	TeamBPlayerA int        `json:"team_b_player_a"`
	TeamBPlayerB int        `json:"team_b_player_b"`
	User1        UserInfo   `gorm:"foreignKey:TeamAPlayerA"`
	User2        UserInfo   `gorm:"foreignKey:TeamAPlayerB"`
	User3        UserInfo   `gorm:"foreignKey:TeamBPlayerA"`
	User4        UserInfo   `gorm:"foreignKey:TeamBPlayerB"`
	Result       *Results   `gorm:"foreignKey:MatchID" `
}

type Results struct {
	ID int
	// MatchId int    `json:"match_id"`
	// Game    Matchs `gorm:"foreignKey:MatchId"`
	AOne    int32 `gorm:"not null;default:0"`
	BOne    int32 `gorm:"not null;default:0"`
	ATwo    int32 `gorm:"not null;default:0"`
	BTwo    int32 `gorm:"not null;default:0"`
	AThree  int32 `gorm:"default:null"`
	BThree  int32 `gorm:"default:null"`
	MatchID int
}

type EloCaulator struct {
	ID       int
	Player   int // mabey forekey to user or just put user uder (Player User)
	EloBefor int
	EloAfter *int
	Game     int // mabey forekey to user or just put user uder (Game Match)
}

/*
	 type MatchHistoryTeams struct {
	}
*/
