package slack

type Slack struct {
	config *Config
}

type Config struct {
	App string
	Url string
}

type SlackMessage struct {
	Username    string             `json:"username"`
	Text        *Format            `json:"text"`
	Attachments []*SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Fallback   string `json:"fallback"`
	Color      string `json:"color"`
	Pretext    string `json:"pretext"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	AuthorIcon string `json:"author_icon"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
	Fields     []struct {
		Title string `json:"title"`
		Value string `json:"value"`
		Short bool   `json:"short"`
	} `json:"fields"`
	ImageURL   string `json:"image_url"`
	ThumbURL   string `json:"thumb_url"`
	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	Ts         int    `json:"ts"`
}

type Format string
type option string
