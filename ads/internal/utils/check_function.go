package utils

import "ads/domain"

func CheckEmptyString(ad string, updateAd string) string {
	if updateAd != "" {
		return updateAd
	}
	return ad
}

func CheckEmptyNumber(ad float32, updateAd float32) float32 {
	if updateAd != 0 {
		return updateAd
	}
	return ad
}

func CheckEmptyPicture(ad domain.Picture, updateAd domain.Picture) domain.Picture {
	var tmp domain.Picture

	if updateAd.Title != "" {
		tmp.Title = updateAd.Title
	} else {
		tmp.Title = ad.Title
	}

	if updateAd.Description != "" {
		tmp.Description = updateAd.Description
	} else {
		tmp.Description = ad.Description
	}

	if updateAd.Url != "" {
		tmp.Url = updateAd.Url
	} else {
		tmp.Url = ad.Url
	}

	return tmp
}
