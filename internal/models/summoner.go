package models

import "time"

type UserInfo struct {
	AccountId                   int64  `json:"accountId"`
	DisplayName                 string `json:"displayName"`
	InternalName                string `json:"internalName"`
	NameChangeFlag              bool   `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int    `json:"percentCompleteForNextLevel"`
	ProfileIconId               int    `json:"profileIconId"`
	Puuid                       string `json:"puuid"`
	RerollPoints                struct {
		CurrentPoints    int `json:"currentPoints"`
		MaxRolls         int `json:"maxRolls"`
		NumberOfRolls    int `json:"numberOfRolls"`
		PointsCostToRoll int `json:"pointsCostToRoll"`
		PointsToReroll   int `json:"pointsToReroll"`
	} `json:"rerollPoints"`
	SummonerId       int64 `json:"summonerId"`
	SummonerLevel    int   `json:"summonerLevel"`
	Unnamed          bool  `json:"unnamed"`
	XpSinceLastLevel int   `json:"xpSinceLastLevel"`
	XpUntilNextLevel int   `json:"xpUntilNextLevel"`
}

// MatchHistoryTft
// TFT查询战绩
type MatchHistoryTft struct {
	ActivePuuid string `json:"active_puuid"`
	Games       []struct {
		Json struct {
			GameDatetime string  `json:"game_datetime"`
			GameId       int64   `json:"game_id"`
			GameLength   float64 `json:"game_length"`
			GameVersion  string  `json:"game_version"`
			Participants []struct {
				Augments  []string `json:"augments"`
				Companion struct {
					ContentID string `json:"content_ID"`
					ItemID    int    `json:"item_ID"`
					SkinID    int    `json:"skin_ID"`
					Species   string `json:"species"`
				} `json:"companion"`
				GoldLeft             int     `json:"gold_left"`
				LastRound            int     `json:"last_round"`
				Level                int     `json:"level"`
				Placement            int     `json:"placement"`
				PlayersEliminated    int     `json:"players_eliminated"`
				Puuid                string  `json:"puuid"`
				TimeEliminated       float64 `json:"time_eliminated"`
				TotalDamageToPlayers int     `json:"total_damage_to_players"`
				Traits               []struct {
					Name        string `json:"name"`
					NumUnits    int    `json:"num_units"`
					Style       int    `json:"style"`
					TierCurrent int    `json:"tier_current"`
					TierTotal   int    `json:"tier_total"`
				} `json:"traits"`
				Units []struct {
					CharacterId string   `json:"character_id"`
					ItemNames   []string `json:"itemNames"`
					Name        string   `json:"name"`
					Rarity      int      `json:"rarity"`
					Tier        int      `json:"tier"`
				} `json:"units"`
			} `json:"participants"`
			QueueId        int    `json:"queue_id"`
			TftGameType    string `json:"tft_game_type"`
			TftSetCoreName string `json:"tft_set_core_name"`
			TftSetNumber   int    `json:"tft_set_number"`
		} `json:"json"`
		Metadata struct {
			DataVersion  int      `json:"data_version"`
			InfoType     string   `json:"info_type"`
			MatchId      string   `json:"match_id"`
			Participants []string `json:"participants"`
			Private      bool     `json:"private"`
			Product      string   `json:"product"`
			Tags         []string `json:"tags"`
			Timestamp    int64    `json:"timestamp"`
		} `json:"metadata"`
	} `json:"games"`
}

// MatchHistoryLol
// LOL查询战绩
type MatchHistoryLol struct {
	AccountId int64 `json:"accountId"`
	Games     struct {
		GameBeginDate  string `json:"gameBeginDate"`
		GameCount      int    `json:"gameCount"`
		GameEndDate    string `json:"gameEndDate"`
		GameIndexBegin int    `json:"gameIndexBegin"`
		GameIndexEnd   int    `json:"gameIndexEnd"`
		Games          []struct {
			GameCreation          int64     `json:"gameCreation"`
			GameCreationDate      time.Time `json:"gameCreationDate"`
			GameDuration          int       `json:"gameDuration"`
			GameId                int64     `json:"gameId"`
			GameMode              string    `json:"gameMode"`
			GameType              string    `json:"gameType"`
			GameVersion           string    `json:"gameVersion"`
			MapId                 int       `json:"mapId"`
			ParticipantIdentities []struct {
				ParticipantId int `json:"participantId"`
				Player        struct {
					AccountId         int    `json:"accountId"`
					CurrentAccountId  int    `json:"currentAccountId"`
					CurrentPlatformId string `json:"currentPlatformId"`
					MatchHistoryUri   string `json:"matchHistoryUri"`
					PlatformId        string `json:"platformId"`
					ProfileIcon       int    `json:"profileIcon"`
					Puuid             string `json:"puuid"`
					SummonerId        int64  `json:"summonerId"`
					SummonerName      string `json:"summonerName"`
				} `json:"player"`
			} `json:"participantIdentities"`
			Participants []struct {
				ChampionId                int    `json:"championId"`
				HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
				ParticipantId             int    `json:"participantId"`
				Spell1Id                  int    `json:"spell1Id"`
				Spell2Id                  int    `json:"spell2Id"`
				Stats                     struct {
					Assists                         int  `json:"assists"`
					CausedEarlySurrender            bool `json:"causedEarlySurrender"`
					ChampLevel                      int  `json:"champLevel"`
					CombatPlayerScore               int  `json:"combatPlayerScore"`
					DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
					DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
					DamageSelfMitigated             int  `json:"damageSelfMitigated"`
					Deaths                          int  `json:"deaths"`
					DoubleKills                     int  `json:"doubleKills"`
					EarlySurrenderAccomplice        bool `json:"earlySurrenderAccomplice"`
					FirstBloodAssist                bool `json:"firstBloodAssist"`
					FirstBloodKill                  bool `json:"firstBloodKill"`
					FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
					FirstInhibitorKill              bool `json:"firstInhibitorKill"`
					FirstTowerAssist                bool `json:"firstTowerAssist"`
					FirstTowerKill                  bool `json:"firstTowerKill"`
					GameEndedInEarlySurrender       bool `json:"gameEndedInEarlySurrender"`
					GameEndedInSurrender            bool `json:"gameEndedInSurrender"`
					GoldEarned                      int  `json:"goldEarned"`
					GoldSpent                       int  `json:"goldSpent"`
					InhibitorKills                  int  `json:"inhibitorKills"`
					Item0                           int  `json:"item0"`
					Item1                           int  `json:"item1"`
					Item2                           int  `json:"item2"`
					Item3                           int  `json:"item3"`
					Item4                           int  `json:"item4"`
					Item5                           int  `json:"item5"`
					Item6                           int  `json:"item6"`
					KillingSprees                   int  `json:"killingSprees"`
					Kills                           int  `json:"kills"`
					LargestCriticalStrike           int  `json:"largestCriticalStrike"`
					LargestKillingSpree             int  `json:"largestKillingSpree"`
					LargestMultiKill                int  `json:"largestMultiKill"`
					LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
					MagicDamageDealt                int  `json:"magicDamageDealt"`
					MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
					MagicalDamageTaken              int  `json:"magicalDamageTaken"`
					NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
					NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
					NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
					ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
					ParticipantId                   int  `json:"participantId"`
					PentaKills                      int  `json:"pentaKills"`
					Perk0                           int  `json:"perk0"`
					Perk0Var1                       int  `json:"perk0Var1"`
					Perk0Var2                       int  `json:"perk0Var2"`
					Perk0Var3                       int  `json:"perk0Var3"`
					Perk1                           int  `json:"perk1"`
					Perk1Var1                       int  `json:"perk1Var1"`
					Perk1Var2                       int  `json:"perk1Var2"`
					Perk1Var3                       int  `json:"perk1Var3"`
					Perk2                           int  `json:"perk2"`
					Perk2Var1                       int  `json:"perk2Var1"`
					Perk2Var2                       int  `json:"perk2Var2"`
					Perk2Var3                       int  `json:"perk2Var3"`
					Perk3                           int  `json:"perk3"`
					Perk3Var1                       int  `json:"perk3Var1"`
					Perk3Var2                       int  `json:"perk3Var2"`
					Perk3Var3                       int  `json:"perk3Var3"`
					Perk4                           int  `json:"perk4"`
					Perk4Var1                       int  `json:"perk4Var1"`
					Perk4Var2                       int  `json:"perk4Var2"`
					Perk4Var3                       int  `json:"perk4Var3"`
					Perk5                           int  `json:"perk5"`
					Perk5Var1                       int  `json:"perk5Var1"`
					Perk5Var2                       int  `json:"perk5Var2"`
					Perk5Var3                       int  `json:"perk5Var3"`
					PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
					PerkSubStyle                    int  `json:"perkSubStyle"`
					PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
					PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
					PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
					PlayerScore0                    int  `json:"playerScore0"`
					PlayerScore1                    int  `json:"playerScore1"`
					PlayerScore2                    int  `json:"playerScore2"`
					PlayerScore3                    int  `json:"playerScore3"`
					PlayerScore4                    int  `json:"playerScore4"`
					PlayerScore5                    int  `json:"playerScore5"`
					PlayerScore6                    int  `json:"playerScore6"`
					PlayerScore7                    int  `json:"playerScore7"`
					PlayerScore8                    int  `json:"playerScore8"`
					PlayerScore9                    int  `json:"playerScore9"`
					QuadraKills                     int  `json:"quadraKills"`
					SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
					TeamEarlySurrendered            bool `json:"teamEarlySurrendered"`
					TimeCCingOthers                 int  `json:"timeCCingOthers"`
					TotalDamageDealt                int  `json:"totalDamageDealt"`
					TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
					TotalDamageTaken                int  `json:"totalDamageTaken"`
					TotalHeal                       int  `json:"totalHeal"`
					TotalMinionsKilled              int  `json:"totalMinionsKilled"`
					TotalPlayerScore                int  `json:"totalPlayerScore"`
					TotalScoreRank                  int  `json:"totalScoreRank"`
					TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
					TotalUnitsHealed                int  `json:"totalUnitsHealed"`
					TripleKills                     int  `json:"tripleKills"`
					TrueDamageDealt                 int  `json:"trueDamageDealt"`
					TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
					TrueDamageTaken                 int  `json:"trueDamageTaken"`
					TurretKills                     int  `json:"turretKills"`
					UnrealKills                     int  `json:"unrealKills"`
					VisionScore                     int  `json:"visionScore"`
					VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
					WardsKilled                     int  `json:"wardsKilled"`
					WardsPlaced                     int  `json:"wardsPlaced"`
					Win                             bool `json:"win"`
				} `json:"stats"`
				TeamId   int `json:"teamId"`
				Timeline struct {
					CreepsPerMinDeltas struct {
					} `json:"creepsPerMinDeltas"`
					CsDiffPerMinDeltas struct {
					} `json:"csDiffPerMinDeltas"`
					DamageTakenDiffPerMinDeltas struct {
					} `json:"damageTakenDiffPerMinDeltas"`
					DamageTakenPerMinDeltas struct {
					} `json:"damageTakenPerMinDeltas"`
					GoldPerMinDeltas struct {
					} `json:"goldPerMinDeltas"`
					Lane               string `json:"lane"`
					ParticipantId      int    `json:"participantId"`
					Role               string `json:"role"`
					XpDiffPerMinDeltas struct {
					} `json:"xpDiffPerMinDeltas"`
					XpPerMinDeltas struct {
					} `json:"xpPerMinDeltas"`
				} `json:"timeline"`
			} `json:"participants"`
			PlatformId string `json:"platformId"`
			QueueId    int    `json:"queueId"`
			SeasonId   int    `json:"seasonId"`
			Teams      []struct {
				Bans                 []interface{} `json:"bans"`
				BaronKills           int           `json:"baronKills"`
				DominionVictoryScore int           `json:"dominionVictoryScore"`
				DragonKills          int           `json:"dragonKills"`
				FirstBaron           bool          `json:"firstBaron"`
				FirstBlood           bool          `json:"firstBlood"`
				FirstDargon          bool          `json:"firstDargon"`
				FirstInhibitor       bool          `json:"firstInhibitor"`
				FirstTower           bool          `json:"firstTower"`
				InhibitorKills       int           `json:"inhibitorKills"`
				RiftHeraldKills      int           `json:"riftHeraldKills"`
				TeamId               int           `json:"teamId"`
				TowerKills           int           `json:"towerKills"`
				VilemawKills         int           `json:"vilemawKills"`
				Win                  string        `json:"win"`
			} `json:"teams"`
		} `json:"games"`
	} `json:"games"`
	PlatformId string `json:"platformId"`
}

// SummonerInProcess
// 游戏内获取到的召唤师信息
type SummonerInProcess struct {
	GameData struct {
		GameId                   int64  `json:"gameId"`
		GameName                 string `json:"gameName"`
		IsCustomGame             bool   `json:"isCustomGame"`
		Password                 string `json:"password"`
		PlayerChampionSelections []struct {
			ChampionId           float64 `json:"championId"`
			SelectedSkinIndex    float64 `json:"selectedSkinIndex"`
			Spell1Id             float64 `json:"spell1Id"`
			Spell2Id             float64 `json:"spell2Id"`
			SummonerInternalName string  `json:"summonerInternalName"`
		} `json:"playerChampionSelections"`
		Queue struct {
			AllowablePremadeSizes   []interface{} `json:"allowablePremadeSizes"`
			AreFreeChampionsAllowed bool          `json:"areFreeChampionsAllowed"`
			AssetMutator            string        `json:"assetMutator"`
			Category                string        `json:"category"`
			ChampionsRequiredToPlay int           `json:"championsRequiredToPlay"`
			Description             string        `json:"description"`
			DetailedDescription     string        `json:"detailedDescription"`
			GameMode                string        `json:"gameMode"`
			GameTypeConfig          struct {
				AdvancedLearningQuests bool   `json:"advancedLearningQuests"`
				AllowTrades            bool   `json:"allowTrades"`
				BanMode                string `json:"banMode"`
				BanTimerDuration       int    `json:"banTimerDuration"`
				BattleBoost            bool   `json:"battleBoost"`
				CrossTeamChampionPool  bool   `json:"crossTeamChampionPool"`
				DeathMatch             bool   `json:"deathMatch"`
				DoNotRemove            bool   `json:"doNotRemove"`
				DuplicatePick          bool   `json:"duplicatePick"`
				ExclusivePick          bool   `json:"exclusivePick"`
				Id                     int    `json:"id"`
				LearningQuests         bool   `json:"learningQuests"`
				MainPickTimerDuration  int    `json:"mainPickTimerDuration"`
				MaxAllowableBans       int    `json:"maxAllowableBans"`
				Name                   string `json:"name"`
				OnboardCoopBeginner    bool   `json:"onboardCoopBeginner"`
				PickMode               string `json:"pickMode"`
				PostPickTimerDuration  int    `json:"postPickTimerDuration"`
				Reroll                 bool   `json:"reroll"`
				TeamChampionPool       bool   `json:"teamChampionPool"`
			} `json:"gameTypeConfig"`
			Id                         int    `json:"id"`
			IsRanked                   bool   `json:"isRanked"`
			IsTeamBuilderManaged       bool   `json:"isTeamBuilderManaged"`
			LastToggledOffTime         int    `json:"lastToggledOffTime"`
			LastToggledOnTime          int    `json:"lastToggledOnTime"`
			MapId                      int    `json:"mapId"`
			MaximumParticipantListSize int    `json:"maximumParticipantListSize"`
			MinLevel                   int    `json:"minLevel"`
			MinimumParticipantListSize int    `json:"minimumParticipantListSize"`
			Name                       string `json:"name"`
			NumPlayersPerTeam          int    `json:"numPlayersPerTeam"`
			QueueAvailability          string `json:"queueAvailability"`
			QueueRewards               struct {
				IsChampionPointsEnabled bool          `json:"isChampionPointsEnabled"`
				IsIpEnabled             bool          `json:"isIpEnabled"`
				IsXpEnabled             bool          `json:"isXpEnabled"`
				PartySizeIpRewards      []interface{} `json:"partySizeIpRewards"`
			} `json:"queueRewards"`
			RemovalFromGameAllowed      bool   `json:"removalFromGameAllowed"`
			RemovalFromGameDelayMinutes int    `json:"removalFromGameDelayMinutes"`
			ShortName                   string `json:"shortName"`
			ShowPositionSelector        bool   `json:"showPositionSelector"`
			SpectatorEnabled            bool   `json:"spectatorEnabled"`
			Type                        string `json:"type"`
		} `json:"queue"`
		SpectatorsAllowed bool `json:"spectatorsAllowed"`
		TeamOne           []struct {
			AccountId         float64 `json:"accountId,omitempty"`
			AdjustmentFlags   float64 `json:"adjustmentFlags,omitempty"`
			BotDifficulty     string  `json:"botDifficulty"`
			ClientInSynch     bool    `json:"clientInSynch,omitempty"`
			GameCustomization struct {
				Regalia        string `json:"Regalia,omitempty"`
				Perks          string `json:"perks,omitempty"`
				SummonerEmotes string `json:"summonerEmotes,omitempty"`
			} `json:"gameCustomization"`
			Index                   float64     `json:"index,omitempty"`
			LastSelectedSkinIndex   float64     `json:"lastSelectedSkinIndex"`
			Locale                  interface{} `json:"locale"`
			Minor                   bool        `json:"minor,omitempty"`
			OriginalAccountNumber   float64     `json:"originalAccountNumber,omitempty"`
			OriginalPlatformId      string      `json:"originalPlatformId,omitempty"`
			PartnerId               string      `json:"partnerId,omitempty"`
			PickMode                float64     `json:"pickMode"`
			PickTurn                float64     `json:"pickTurn"`
			ProfileIconId           float64     `json:"profileIconId,omitempty"`
			Puuid                   string      `json:"puuid,omitempty"`
			QueueRating             float64     `json:"queueRating,omitempty"`
			RankedTeamGuest         bool        `json:"rankedTeamGuest,omitempty"`
			SelectedPosition        interface{} `json:"selectedPosition"`
			SelectedRole            interface{} `json:"selectedRole"`
			SummonerId              float64     `json:"summonerId,omitempty"`
			SummonerInternalName    string      `json:"summonerInternalName"`
			SummonerName            string      `json:"summonerName"`
			TeamOwner               bool        `json:"teamOwner,omitempty"`
			TeamParticipantId       interface{} `json:"teamParticipantId"`
			TeamRating              float64     `json:"teamRating,omitempty"`
			TimeAddedToQueue        interface{} `json:"timeAddedToQueue"`
			TimeChampionSelectStart float64     `json:"timeChampionSelectStart,omitempty"`
			TimeGameCreated         float64     `json:"timeGameCreated,omitempty"`
			TimeMatchmakingStart    float64     `json:"timeMatchmakingStart,omitempty"`
			VoterRating             float64     `json:"voterRating,omitempty"`
			BotSkillLevel           float64     `json:"botSkillLevel,omitempty"`
			ChampionId              interface{} `json:"championId"`
			Role                    interface{} `json:"role"`
			Spell1Id                interface{} `json:"spell1Id"`
			Spell2Id                interface{} `json:"spell2Id"`
			TeamId                  string      `json:"teamId,omitempty"`
		} `json:"teamOne"`
		TeamTwo []struct {
			BotDifficulty     string      `json:"botDifficulty"`
			BotSkillLevel     float64     `json:"botSkillLevel"`
			ChampionId        interface{} `json:"championId"`
			GameCustomization struct {
			} `json:"gameCustomization"`
			LastSelectedSkinIndex float64     `json:"lastSelectedSkinIndex"`
			Locale                interface{} `json:"locale"`
			PickMode              float64     `json:"pickMode"`
			PickTurn              float64     `json:"pickTurn"`
			Role                  interface{} `json:"role"`
			Spell1Id              interface{} `json:"spell1Id"`
			Spell2Id              interface{} `json:"spell2Id"`
			SummonerInternalName  string      `json:"summonerInternalName"`
			SummonerName          string      `json:"summonerName"`
			TeamId                string      `json:"teamId"`
		} `json:"teamTwo"`
	} `json:"gameData"`
	GameDodge struct {
		DodgeIds []interface{} `json:"dodgeIds"`
		Phase    string        `json:"phase"`
		State    string        `json:"state"`
	} `json:"gameDodge"`
	Map struct {
		Assets struct {
			ChampSelectBackgroundSound  string `json:"champ-select-background-sound"`
			ChampSelectFlyoutBackground string `json:"champ-select-flyout-background"`
			ChampSelectPlanningIntro    string `json:"champ-select-planning-intro"`
			GameSelectIconActive        string `json:"game-select-icon-active"`
			GameSelectIconActiveVideo   string `json:"game-select-icon-active-video"`
			GameSelectIconDefault       string `json:"game-select-icon-default"`
			GameSelectIconDisabled      string `json:"game-select-icon-disabled"`
			GameSelectIconHover         string `json:"game-select-icon-hover"`
			GameSelectIconIntroVideo    string `json:"game-select-icon-intro-video"`
			GameflowBackground          string `json:"gameflow-background"`
			GameflowBackgroundDark      string `json:"gameflow-background-dark"`
			GameselectButtonHoverSound  string `json:"gameselect-button-hover-sound"`
			IconDefeat                  string `json:"icon-defeat"`
			IconDefeatV2                string `json:"icon-defeat-v2"`
			IconDefeatVideo             string `json:"icon-defeat-video"`
			IconEmpty                   string `json:"icon-empty"`
			IconHover                   string `json:"icon-hover"`
			IconLeaver                  string `json:"icon-leaver"`
			IconLeaverV2                string `json:"icon-leaver-v2"`
			IconLossForgivenV2          string `json:"icon-loss-forgiven-v2"`
			IconV2                      string `json:"icon-v2"`
			IconVictory                 string `json:"icon-victory"`
			IconVictoryVideo            string `json:"icon-victory-video"`
			MapNorth                    string `json:"map-north"`
			MapSouth                    string `json:"map-south"`
			MusicInqueueLoopSound       string `json:"music-inqueue-loop-sound"`
			PartiesBackground           string `json:"parties-background"`
			PostgameAmbienceLoopSound   string `json:"postgame-ambience-loop-sound"`
			ReadyCheckBackground        string `json:"ready-check-background"`
			ReadyCheckBackgroundSound   string `json:"ready-check-background-sound"`
			SfxAmbiencePregameLoopSound string `json:"sfx-ambience-pregame-loop-sound"`
			SocialIconLeaver            string `json:"social-icon-leaver"`
			SocialIconVictory           string `json:"social-icon-victory"`
		} `json:"assets"`
		CategorizedContentBundles struct {
		} `json:"categorizedContentBundles"`
		Description                         string `json:"description"`
		GameMode                            string `json:"gameMode"`
		GameModeName                        string `json:"gameModeName"`
		GameModeShortName                   string `json:"gameModeShortName"`
		GameMutator                         string `json:"gameMutator"`
		Id                                  int    `json:"id"`
		IsRGM                               bool   `json:"isRGM"`
		MapStringId                         string `json:"mapStringId"`
		Name                                string `json:"name"`
		PerPositionDisallowedSummonerSpells struct {
		} `json:"perPositionDisallowedSummonerSpells"`
		PerPositionRequiredSummonerSpells struct {
		} `json:"perPositionRequiredSummonerSpells"`
		PlatformId   string `json:"platformId"`
		PlatformName string `json:"platformName"`
		Properties   struct {
			SuppressRunesMasteriesPerks bool `json:"suppressRunesMasteriesPerks"`
		} `json:"properties"`
	} `json:"map"`
	Phase string `json:"phase"`
}
