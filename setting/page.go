package setting

import (
	"errors"
	"os"
	"strconv"

	"github.com/kubestaff/golearning/helpers"
	"github.com/kubestaff/web-helper/server"
)

func HandleReadSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helpers.HandleErrorText("Invalid user id")
	}

	if inputs.Get("amountOfUsersOnMainPage") != "" {
		return HandleformSetting(inputs)
	}

	userProvider := Provider{}
	setting, isFound, err := userProvider.GetSettingByUserId(userIdInt)
	if err != nil {
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

func HandleformSetting(inputs server.Input) (filename string, placeholders map[string]string) {
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

	settingsProvider := Provider{}
	existingSetting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return helpers.HandleErr(err)
	}

	newSetting := UserSetting{
		AmountOfUsersOnMainPage: amountOfUsersOnMainPageSubmittedInt,
		UserId:                  userIdInt,
	}
	if isFound {
		newSetting.Id = existingSetting.Id
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
