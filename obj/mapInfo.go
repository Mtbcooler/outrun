package obj

const (
	DefaultMapDistance   = 0
	DefaultNumBossAttack = 0 // Boss HP gets subtracted by this number, for some reason.
	DefaultStageDistance = 1337000
	DefaultStageMaxScore = 8008135
)

type MapInfo struct {
	MapDistance   int64 `json:"mapDistance" db:"map_distance"`      // used sparingly in game...?
	NumBossAttack int64 `json:"numBossAttack" db:"num_boss_attack"` // TODO: number of boss fights? Check how often this is used in game
	StageDistance int64 `json:"stageDistance" db:"stage_distance"`  // TODO: discover use
	StageMaxScore int64 `json:"stageMaxScore" db:"stage_max_score"` // TODO: discover use
}

func DefaultMapInfo() MapInfo {
	return MapInfo{
		DefaultMapDistance,
		DefaultNumBossAttack,
		DefaultStageDistance,
		DefaultStageMaxScore,
	}
}
