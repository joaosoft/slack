package slack

var (
	options = map[option]string{
		optionQuote:      "`%s`",
		optionBold:       "*%s*",
		optionNewLine:    "%s\n",
		optionBlockQuote: ">%s",
		optionCode:       "```%s```",
	}
)
