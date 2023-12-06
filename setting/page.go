package setting

import (
	"strconv"
	"os"
	"errors"

	"github.com/kubestaff/golearning/helper"
	"github.com/kubestaff/web-helper/server"

)

func HandleReadSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	userIdStr :=inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	If err != nil {
		return helper.HandleErrorText("Invalid user id")
	}

	if input.Get("amountOfUsersOnMainPage") != "" {
		return HandleFormSetting(inputs)
	}

	settingsProvider := Provider{}
	setting, isFound, err:= settingssProvider.GetSettingByUserId(userIdInt)

	//if there is an error but this error is not about non existing file 
	if err != nil && !os.IsNotExist(err) {
		return helper.HandleErr(err)
	}

	output := map[string]string{
		"%amountOfUsersOnMainPage%": "10",
		"%id%": userIdStr,
	}
	If isFound {
		output["%amountOfUsersOnMainPage%"] = strconv.Itoa(setting.AmountOfUsersOnMainPage),		
		
	}
	return "html/setting.html", output
}

func HandleFormSetting(inputs server.Input) (filename string, placeholders map[string]string) {
	amountOfUsersOnMainPageSubmittedStr := inputs.Get("amountOfUsersOnMainPage")
	amountOfUsersOnMainPageSubmittedInt, err := strconv.Atoi(amountOfUsersOnMainPageSubmittedStr)
	If err != nil {
		return helper.HandleErrorText("Invalid amount of user on main page")
	}
	userIdStr :=inputs.Values.Get("id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return helper.HandleErrorText("Invalid user id")
	}

	settingsProvider := Provider{}
	existingSetting, isFound, err := settingssProvider.GetSettingByUserId(userIdInt)
	if err != nil && !os.IsNotExist(err) {
		return helper.HandleErr(err)
	}

	setting := UserSetting{
		AmountOfUsersOnMainPage: amountOfUsersOnMainPageSubmittedInt,
		UserId: userIdInt,
	}

	if isFound {
		newSetting.Id = existingSetting.Id

	}

	err = settingsProvider.SaveSetting(&newSetting)
	if err != nil {
		return helper.HandleErr(err)
	}

	output := map[string]string{
		"%amountOfUsersOnMainPage%": strconv.Itoa(amountOfUsersOnMainPageSubmittedInt),
		"%id%": userIdStr,
	}
	
	return "html/setting.html", output
}