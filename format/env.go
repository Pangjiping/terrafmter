package format

import (
	"fmt"
	"os"
)

func getAccessKey() (string, error) {
	accKey := os.Getenv(ACCESS_KEY)
	if accKey == "" {
		return "", fmt.Errorf("cannot get %s, please check env", ACCESS_KEY)
	}
	return accKey, nil
}

func getSecretKey() (string, error) {
	secKey := os.Getenv(SECRET_KEY)
	if secKey == "" {
		return "", fmt.Errorf("cannot get %s, please check env", SECRET_KEY)
	}
	return secKey, nil
}

func getRegion() (string, error) {
	region := os.Getenv(REGION)
	if region == "" {
		if err := os.Setenv(REGION, DEFAULT_REGION); err != nil {
			return "", err
		}
		return DEFAULT_REGION, nil
	}
	return region, nil
}
