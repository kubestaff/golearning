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
		return helpers.HandleErrorText("invalid user id")
	}

	if inputs.Get("amountOfUsersOnMainPage") != "" {
		return HandleFormSetting(inputs)
	}

	settingsProvider := Provider{}
	setting, isFound, err := settingsProvider.GetSettingByUserId(userIdInt)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
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
