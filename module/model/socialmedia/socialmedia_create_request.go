package socialmedia

type SocialmediaCreateRequest struct {
	Name           string `json:"name" gorm:"column:name;type:varchar(255)"`
	SocialMediaUrl string `json:"social_media_url"`
}
