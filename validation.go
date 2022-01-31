package main

import (
	"fmt"
	"net/url"
	"regexp"
)

func Jwt(token string) error {
	err := NotBlank(token)

	if err != nil {
		return fmt.Errorf("invalid token: %s", err.Error())
	}

	// @see https://github.com/symfony/mercure/blob/7546092e654f9bb22e554819fb614e462075065a/src/Hub.php#L108-L123
	expression, err := regexp.Compile("^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$")

	if err != nil {
		return fmt.Errorf("invalid token: %s", err.Error())
	}

	if !expression.MatchString(token) {
		return fmt.Errorf("token is not valid")
	}

	return nil
}

func NotBlank(str string) error {
	if str == "" {
		return fmt.Errorf("this value cannot be blank")
	}

	return nil
}

func All(items []string, validator func(string) error) error {
	for _, item := range items {
		if err := validator(item); err != nil {
			return err
		}
	}

	return nil
}

func Url(s string) error {
	_, err := url.Parse(s)

	if err != nil {
		return err
	}

	return nil
}
