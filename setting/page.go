package setting

import (
	"os"
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/web-helper/server"
	"gorm.io/gorm"
)

type Handler struct {
	DbConnection *gorm.DB
}

func (h Handler) HandleReadSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("Invalid user id")
	}

	if inputs.Get("amountOfUsersOnMainPage") != "" {
		return h.HandleFormSetting(inputs)
	}

	settingsProvider := Provider{
		DbConnection: h.DbConnection,
	}
	setting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)
	if err != nil && !os.IsNotExist(err) {
		return helpers.HandleErr(err)
	}

	output := map[string]string{
		"%amountOfUsersOnMainPage%": "10",
		"%id%":                      userIdStr,
	}

	if isFound {
		output["%amountOfUsersOnMainPage%"] = strconv.Itoa(setting.AmountOfUsersOnMainPage)
	}

	return "html/setting.html", output
}

func (h Handler) HandleFormSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	amountOfUsersOnMainPageSubmittedStr := inputs.Get("amountOfUsersOnMainPage")
	amountOfUsersOnMainPageSubmittedInt, err := strconv.Atoi(amountOfUsersOnMainPageSubmittedStr)
	if err != nil {
		return helpers.HandleErrorText("invalid amount of users on main page")
	}

	userIdStr := inputs.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("invalid user id")
	}

	settingsProvider := Provider{
		DbConnection: h.DbConnection,
	}
	existingSetting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)
	if err != nil && !os.IsNotExist(err) {
		return helpers.HandleErr(err)
	}

	newSetting := UserSetting{
		AmountOfUsersOnMainPage: amountOfUsersOnMainPageSubmittedInt,
		UserId:                  userIdInt,
	}
	if isFound {
		newSetting.ID = existingSetting.ID
	}

	err = settingsProvider.SaveSetting(&newSetting)
	if err != nil {
		helpers.HandleErr(err)
	}

	output := map[string]string{
		"%amountOfUsersOnMainPage%": amountOfUsersOnMainPageSubmittedStr,
		"%id%":                      userIdStr,
	}

	return "html/setting.html", output
}
