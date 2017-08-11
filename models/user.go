package models

import (
	"time"
)

type UserResponse struct {
	Success bool `json:"success"`
	Data    struct {
		UnderscoreID      string `json:"_id"`
		ABTests struct {
			OnboardingPushNotification string `json:"onboardingPushNotification"`
		} `json:"_ABTests"`
		InvitesSent     int           `json:"invitesSent"`
		LoginIncentives int           `json:"loginIncentives"`
		Webhooks        []interface{} `json:"webhooks"`
		PushDevices     []interface{} `json:"pushDevices"`
		Extra           struct {
		} `json:"extra"`
		TasksOrder struct {
			Rewards []string      `json:"rewards"`
			Todos   []string      `json:"todos"`
			Dailys  []interface{} `json:"dailys"`
			Habits  []string      `json:"habits"`
		} `json:"tasksOrder"`
		Inbox struct {
			OptOut   bool `json:"optOut"`
			Messages struct {
			} `json:"messages"`
			Blocks      []interface{} `json:"blocks"`
			NewMessages int           `json:"newMessages"`
		} `json:"inbox"`
		Tags []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"tags"`
		Notifications []interface{} `json:"notifications"`
		Stats         struct {
			Training struct {
				Con int `json:"con"`
				Str int `json:"str"`
				Per int `json:"per"`
				Int int `json:"int"`
			} `json:"training"`
			Buffs struct {
				Seafoam        bool `json:"seafoam"`
				ShinySeed      bool `json:"shinySeed"`
				SpookySparkles bool `json:"spookySparkles"`
				Snowball       bool `json:"snowball"`
				Streaks        bool `json:"streaks"`
				Stealth        int  `json:"stealth"`
				Con            int  `json:"con"`
				Per            int  `json:"per"`
				Int            int  `json:"int"`
				Str            int  `json:"str"`
			} `json:"buffs"`
			Per         int    `json:"per"`
			Int         int    `json:"int"`
			Con         int    `json:"con"`
			Str         int    `json:"str"`
			Points      int    `json:"points"`
			Class       string `json:"class"`
			Lvl         int    `json:"lvl"`
			Gp          int    `json:"gp"`
			Exp         int    `json:"exp"`
			Mp          int    `json:"mp"`
			Hp          int    `json:"hp"`
			ToNextLevel int    `json:"toNextLevel"`
			MaxHealth   int    `json:"maxHealth"`
			MaxMP       int    `json:"maxMP"`
		} `json:"stats"`
		Profile struct {
			Name string `json:"name"`
		} `json:"profile"`
		Preferences struct {
			Language              string        `json:"language"`
			Background            string        `json:"background"`
			ImprovementCategories []interface{} `json:"improvementCategories"`
			Tasks                 struct {
				ConfirmScoreNotes bool `json:"confirmScoreNotes"`
				GroupByChallenge  bool `json:"groupByChallenge"`
			} `json:"tasks"`
			SuppressModals struct {
				Streak   bool `json:"streak"`
				RaisePet bool `json:"raisePet"`
				HatchPet bool `json:"hatchPet"`
				LevelUp  bool `json:"levelUp"`
			} `json:"suppressModals"`
			PushNotifications struct {
				InvitedQuest       bool `json:"invitedQuest"`
				QuestStarted       bool `json:"questStarted"`
				InvitedGuild       bool `json:"invitedGuild"`
				InvitedParty       bool `json:"invitedParty"`
				GiftedSubscription bool `json:"giftedSubscription"`
				GiftedGems         bool `json:"giftedGems"`
				WonChallenge       bool `json:"wonChallenge"`
				NewPM              bool `json:"newPM"`
				UnsubscribeFromAll bool `json:"unsubscribeFromAll"`
			} `json:"pushNotifications"`
			EmailNotifications struct {
				Onboarding             bool `json:"onboarding"`
				WeeklyRecaps           bool `json:"weeklyRecaps"`
				ImportantAnnouncements bool `json:"importantAnnouncements"`
				InvitedQuest           bool `json:"invitedQuest"`
				QuestStarted           bool `json:"questStarted"`
				InvitedGuild           bool `json:"invitedGuild"`
				InvitedParty           bool `json:"invitedParty"`
				GiftedSubscription     bool `json:"giftedSubscription"`
				GiftedGems             bool `json:"giftedGems"`
				WonChallenge           bool `json:"wonChallenge"`
				KickedGroup            bool `json:"kickedGroup"`
				NewPM                  bool `json:"newPM"`
				UnsubscribeFromAll     bool `json:"unsubscribeFromAll"`
			} `json:"emailNotifications"`
			Webhooks struct {
			} `json:"webhooks"`
			DisplayInviteToPartyWhenPartyIs1 bool   `json:"displayInviteToPartyWhenPartyIs1"`
			ReverseChatOrder                 bool   `json:"reverseChatOrder"`
			ToolbarCollapsed                 bool   `json:"toolbarCollapsed"`
			AdvancedCollapsed                bool   `json:"advancedCollapsed"`
			TagsCollapsed                    bool   `json:"tagsCollapsed"`
			DailyDueDefaultView              bool   `json:"dailyDueDefaultView"`
			NewTaskEdit                      bool   `json:"newTaskEdit"`
			DisableClasses                   bool   `json:"disableClasses"`
			StickyHeader                     bool   `json:"stickyHeader"`
			Sleep                            bool   `json:"sleep"`
			DateFormat                       string `json:"dateFormat"`
			AutoEquip                        bool   `json:"autoEquip"`
			AllocationMode                   string `json:"allocationMode"`
			Chair                            string `json:"chair"`
			Sound                            string `json:"sound"`
			TimezoneOffset                   int    `json:"timezoneOffset"`
			Shirt                            string `json:"shirt"`
			Skin                             string `json:"skin"`
			HideHeader                       bool   `json:"hideHeader"`
			Hair                             struct {
				Flower   int    `json:"flower"`
				Mustache int    `json:"mustache"`
				Beard    int    `json:"beard"`
				Bangs    int    `json:"bangs"`
				Base     int    `json:"base"`
				Color    string `json:"color"`
			} `json:"hair"`
			Size     string `json:"size"`
			DayStart int    `json:"dayStart"`
		} `json:"preferences"`
		Party struct {
			Quest struct {
				RSVPNeeded bool `json:"RSVPNeeded"`
				Progress   struct {
					CollectedItems int `json:"collectedItems"`
					Collect        struct {
					} `json:"collect"`
					Down int `json:"down"`
					Up   int `json:"up"`
				} `json:"progress"`
			} `json:"quest"`
			OrderAscending string `json:"orderAscending"`
			Order          string `json:"order"`
		} `json:"party"`
		Guilds      []interface{} `json:"guilds"`
		Invitations struct {
			Parties []interface{} `json:"parties"`
			Party   struct {
			} `json:"party"`
			Guilds []interface{} `json:"guilds"`
		} `json:"invitations"`
		Challenges  []interface{} `json:"challenges"`
		NewMessages struct {
		} `json:"newMessages"`
		LastCron time.Time `json:"lastCron"`
		Items    struct {
			LastDrop struct {
				Count int       `json:"count"`
				Date  time.Time `json:"date"`
			} `json:"lastDrop"`
			Quests struct {
				Dustbunnies int `json:"dustbunnies"`
			} `json:"quests"`
			Mounts struct {
			} `json:"mounts"`
			Food struct {
			} `json:"food"`
			HatchingPotions struct {
			} `json:"hatchingPotions"`
			Eggs struct {
			} `json:"eggs"`
			Pets struct {
			} `json:"pets"`
			Special struct {
				GoodluckReceived  []interface{} `json:"goodluckReceived"`
				Goodluck          int           `json:"goodluck"`
				GetwellReceived   []interface{} `json:"getwellReceived"`
				Getwell           int           `json:"getwell"`
				CongratsReceived  []interface{} `json:"congratsReceived"`
				Congrats          int           `json:"congrats"`
				BirthdayReceived  []interface{} `json:"birthdayReceived"`
				Birthday          int           `json:"birthday"`
				ThankyouReceived  []interface{} `json:"thankyouReceived"`
				Thankyou          int           `json:"thankyou"`
				GreetingReceived  []interface{} `json:"greetingReceived"`
				Greeting          int           `json:"greeting"`
				NyeReceived       []interface{} `json:"nyeReceived"`
				Nye               int           `json:"nye"`
				ValentineReceived []interface{} `json:"valentineReceived"`
				Valentine         int           `json:"valentine"`
				Seafoam           int           `json:"seafoam"`
				ShinySeed         int           `json:"shinySeed"`
				SpookySparkles    int           `json:"spookySparkles"`
				Snowball          int           `json:"snowball"`
			} `json:"special"`
			Gear struct {
				Costume struct {
					Shield string `json:"shield"`
					Head   string `json:"head"`
					Armor  string `json:"armor"`
				} `json:"costume"`
				Equipped struct {
					Shield string `json:"shield"`
					Head   string `json:"head"`
					Armor  string `json:"armor"`
				} `json:"equipped"`
				Owned struct {
					EyewearSpecialYellowTopFrame bool `json:"eyewear_special_yellowTopFrame"`
					EyewearSpecialWhiteTopFrame  bool `json:"eyewear_special_whiteTopFrame"`
					EyewearSpecialRedTopFrame    bool `json:"eyewear_special_redTopFrame"`
					EyewearSpecialPinkTopFrame   bool `json:"eyewear_special_pinkTopFrame"`
					EyewearSpecialGreenTopFrame  bool `json:"eyewear_special_greenTopFrame"`
					EyewearSpecialBlueTopFrame   bool `json:"eyewear_special_blueTopFrame"`
					EyewearSpecialBlackTopFrame  bool `json:"eyewear_special_blackTopFrame"`
				} `json:"owned"`
			} `json:"gear"`
		} `json:"items"`
		History struct {
			Todos []interface{} `json:"todos"`
			Exp   []interface{} `json:"exp"`
		} `json:"history"`
		Flags struct {
			OnboardingEmailsPhase       string    `json:"onboardingEmailsPhase"`
			WarnedLowHealth             bool      `json:"warnedLowHealth"`
			CardReceived                bool      `json:"cardReceived"`
			ArmoireEmpty                bool      `json:"armoireEmpty"`
			ArmoireOpened               bool      `json:"armoireOpened"`
			ArmoireEnabled              bool      `json:"armoireEnabled"`
			Welcomed                    bool      `json:"welcomed"`
			CronCount                   int       `json:"cronCount"`
			CommunityGuidelinesAccepted bool      `json:"communityGuidelinesAccepted"`
			LastWeeklyRecap             time.Time `json:"lastWeeklyRecap"`
			WeeklyRecapEmailsPhase      int       `json:"weeklyRecapEmailsPhase"`
			RecaptureEmailsPhase        int       `json:"recaptureEmailsPhase"`
			LevelDrops                  struct {
			} `json:"levelDrops"`
			RebirthEnabled bool `json:"rebirthEnabled"`
			ClassSelected  bool `json:"classSelected"`
			Rewrite        bool `json:"rewrite"`
			NewStuff       bool `json:"newStuff"`
			ItemsEnabled   bool `json:"itemsEnabled"`
			DropsEnabled   bool `json:"dropsEnabled"`
			Tutorial       struct {
				Ios struct {
					ReorderTask bool `json:"reorderTask"`
					InviteParty bool `json:"inviteParty"`
					GroupPets   bool `json:"groupPets"`
					FilterTask  bool `json:"filterTask"`
					DeleteTask  bool `json:"deleteTask"`
					EditTask    bool `json:"editTask"`
					AddTask     bool `json:"addTask"`
				} `json:"ios"`
				Common struct {
					Inbox     bool `json:"inbox"`
					Mounts    bool `json:"mounts"`
					Items     bool `json:"items"`
					Equipment bool `json:"equipment"`
					Tavern    bool `json:"tavern"`
					Classes   bool `json:"classes"`
					Skills    bool `json:"skills"`
					Gems      bool `json:"gems"`
					Pets      bool `json:"pets"`
					Party     bool `json:"party"`
					Rewards   bool `json:"rewards"`
					Todos     bool `json:"todos"`
					Dailies   bool `json:"dailies"`
					Habits    bool `json:"habits"`
				} `json:"common"`
			} `json:"tutorial"`
			Tour struct {
				Equipment  int `json:"equipment"`
				Hall       int `json:"hall"`
				Mounts     int `json:"mounts"`
				Pets       int `json:"pets"`
				Market     int `json:"market"`
				Challenges int `json:"challenges"`
				Guilds     int `json:"guilds"`
				Party      int `json:"party"`
				Tavern     int `json:"tavern"`
				Stats      int `json:"stats"`
				Classes    int `json:"classes"`
				Intro      int `json:"intro"`
			} `json:"tour"`
			ShowTour                   bool `json:"showTour"`
			CustomizationsNotification bool `json:"customizationsNotification"`
		} `json:"flags"`
		Purchased struct {
			Plan struct {
				Consecutive struct {
					Trinkets    int `json:"trinkets"`
					GemCapExtra int `json:"gemCapExtra"`
					Offset      int `json:"offset"`
					Count       int `json:"count"`
				} `json:"consecutive"`
				MysteryItems []interface{} `json:"mysteryItems"`
				GemsBought   int           `json:"gemsBought"`
				ExtraMonths  int           `json:"extraMonths"`
				Quantity     int           `json:"quantity"`
			} `json:"plan"`
			TxnCount   int `json:"txnCount"`
			Background struct {
				Violet bool `json:"violet"`
			} `json:"background"`
			Shirt struct {
			} `json:"shirt"`
			Hair struct {
			} `json:"hair"`
			Skin struct {
			} `json:"skin"`
			Ads bool `json:"ads"`
		} `json:"purchased"`
		Balance     int `json:"balance"`
		Contributor struct {
		} `json:"contributor"`
		Backer struct {
		} `json:"backer"`
		Achievements struct {
			Perfect int `json:"perfect"`
			Quests  struct {
			} `json:"quests"`
			Challenges       []interface{} `json:"challenges"`
			UltimateGearSets struct {
				Warrior bool `json:"warrior"`
				Rogue   bool `json:"rogue"`
				Wizard  bool `json:"wizard"`
				Healer  bool `json:"healer"`
			} `json:"ultimateGearSets"`
		} `json:"achievements"`
		V    int `json:"_v"`
		Auth struct {
			Timestamps struct {
				Loggedin time.Time `json:"loggedin"`
				Created  time.Time `json:"created"`
			} `json:"timestamps"`
			Local struct {
				Username          string `json:"username"`
				LowerCaseUsername string `json:"lowerCaseUsername"`
				Email             string `json:"email"`
			} `json:"local"`
			Google struct {
			} `json:"google"`
			Facebook struct {
			} `json:"facebook"`
		} `json:"auth"`
		ID        string `json:"id"`
		NeedsCron bool   `json:"needsCron"`
	} `json:"data"`
	Notifications []interface{} `json:"notifications"`
}