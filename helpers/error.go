package helpers

func HandleErrorText(err string) (filename string, placeholders map[string]string) {

	output := map[string]string {
		"%error%": err,
	}

	return "html/error.html", output
}

func HandleErr(err error) (filename string, placeholders map[string]string) {
	return HandleErrorText(err.Error())
}
