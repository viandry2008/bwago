package campaign

import "strings"

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.ID = campaign.ID
	campaignFormater.UserID = campaign.UserID
	campaignFormater.Name = campaign.Name
	campaignFormater.ShortDescription = campaign.ShortDescription
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	campaignFormater.Slug = campaign.Slug
	campaignFormater.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormater.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormater
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormater {
	campaignsFormater := []CampaignFormater{}

	for _, campaign := range campaigns {
		campaignFormater := FormatCampaign(campaign)
		campaignsFormater = append(campaignsFormater, campaignFormater)
	}
	return campaignsFormater
}

type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	CampaignDetailFormatter := CampaignDetailFormatter{}
	CampaignDetailFormatter.ID = campaign.ID
	CampaignDetailFormatter.Name = campaign.Name
	CampaignDetailFormatter.ShortDescription = campaign.ShortDescription
	CampaignDetailFormatter.Description = campaign.Description
	CampaignDetailFormatter.GoalAmount = campaign.GoalAmount
	CampaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignDetailFormatter.UserID = campaign.UserID
	CampaignDetailFormatter.Slug = campaign.Slug
	CampaignDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	CampaignDetailFormatter.Perks = perks

	user := campaign.User

	CampaignUserFormatter := CampaignUserFormatter{}
	CampaignUserFormatter.Name = user.Name
	CampaignUserFormatter.ImageURL = user.AvatarFileName

	CampaignDetailFormatter.User = CampaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		CampaignImageFormatter := CampaignImageFormatter{}

		CampaignImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		CampaignImageFormatter.IsPrimary = isPrimary

		images = append(images, CampaignImageFormatter)
	}

	CampaignDetailFormatter.Images = images

	return CampaignDetailFormatter
}
