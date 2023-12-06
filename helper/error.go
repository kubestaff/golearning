package helper

func HandleErrorText(errText string) (flename string, placeholders map[string]string) {
	output := map[string]string{
		"%error%": 	errText,	
	}

	return "html/error.html", output
}

func HandleErr(err error) (flename string, placeholders map[string]string) {
	return HandleErrorText(err.Error())
}