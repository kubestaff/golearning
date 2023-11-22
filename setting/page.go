package setting

import (
	"os"
	"strconv"

	"github.com/kubestaff/golearning/helper"
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
		return helper.HandleErrorText("Invalid user id")
	}

	if inputs.Get("amountOfUsersOnMainPage") != "" {
		return h.HandleFormSetting(inputs)
	}

	settingsProvider := Provider{
		DbConnection: h.DbConnection,
	}
	setting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)

	//if there is an error but this error is not about non existing file
	if err != nil && !os.IsNotExist(err) {
		return helper.HandleErr(err)
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
		return helper.HandleErrorText("Invalid amount of users on main page")
	}
	userIdStr := inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helper.HandleErrorText("Invalid user id")
	}

	settingsProvider := Provider{
		DbConnection: h.DbConnection,
	}
	existingSetting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)
	if err != nil && !os.IsNotExist(err) {
		return helper.HandleErr(err)
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
		return helper.HandleErr(err)
	}

	output := map[string]string{
		"%amountOfUsersOnMainPage%": strconv.Itoa(amountOfUsersOnMainPageSubmittedInt),
		"%id%":                      userIdStr,
	}

	return "html/setting.html", output
}
func (h Handler) HandleDeleteUserSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr := inputs.Values.Get("id")

	userIdInt := 0
	var err error
	if userIdStr != "" {
		userIdInt, err = strconv.Atoi(userIdStr)
		if err != nil {
			return helper.HandleErrorText("Invalid user id")
		}
	}

	userProvider := Provider{
		DbConnection: h.DbConnection,
	}
	usr, isFound, err := userProvider.GetSettingByUserId(userIdInt)
	if err != nil {
		return helper.HandleErr(err)
	}

	if !isFound {
		return helper.HandleErrorText("User id is not found")
	}

	err = userProvider.DeleteSetting(&usr)
	if err != nil {
		return helper.HandleErr(err)
	}

	return "html/success.html", map[string]string{"%success%": "Successfully deleted user"}
}
