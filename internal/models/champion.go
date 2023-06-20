package models

type ChampionSelectionEvent struct {
	EventType string `json:"eventType"`
	Data      struct {
		Actions []struct {
			ActorCellID   int    `json:"actorCellId"`
			ActionID      int    `json:"actionId"`
			ChampionID    int    `json:"championId"`
			Completed     bool   `json:"completed"`
			ID            int    `json:"id"`
			IsAllyAction  bool   `json:"isAllyAction"`
			PickTurn      int    `json:"pickTurn"`
			RemainingTime int    `json:"remainingTime"`
			Type          string `json:"type"`
		} `json:"actions"`
		MyTeam []struct {
			CellID     int    `json:"cellId"`
			ChampionID int    `json:"championId"`
			SummonerID string `json:"summonerId"`
		} `json:"myTeam"`
		TheirTeam []struct {
			CellID     int    `json:"cellId"`
			ChampionID int    `json:"championId"`
			SummonerID string `json:"summonerId"`
		} `json:"theirTeam"`
	} `json:"data"`
}

type RiotChampionInfo struct {
}

type ChampionInfo struct {
}

type ChampionVersionData struct {
	Version string `json:"version"`
	Hero    []Hero
}

type Hero struct {
	Alias      string `json:"alias,omitempty"`
	Id         string `json:"heroId"`
	IsWeekFree string `json:"isWeekFree"`
	Name       string `json:"name"`
	Title      string `json:"title"`
}
