package split

import (
	"fmt"
	"strings"
)

var secretKeys = []string{
	"_key",
	"private",
	"secret",
	"password",
	"account_id",
	"token",
	"_api_",
	"auth_token",
	"auth_password",
	"servers",
	"memcached_url",
	"memcachedcloud_username",
	"database_url",
	"redis_url",
	"flipper_url",
	"webhook_url",
}

func Do(dfault map[string]string, override map[string]string) (map[string]string, map[string]string, map[string]string, error) {
	dfault = hideSecrets(dfault)
	override = hideSecrets(override)

	defaultOnly := map[string]string{}
	overrideOnly := map[string]string{}
	overrides := map[string]string{}

	for k, v := range override {
		if _, ok := dfault[k]; !ok {
			overrideOnly[k] = v
			continue
		}

		// if value appears in both sets, it goes in overrides
		if _, ok := dfault[k]; ok {
			if _, ok := override[k]; ok {
				overrides[k] = v
			}
		}
	}

	for k, v := range dfault {
		if k == "BASIC_AUTH_USERNAME" {
			fmt.Println(k)
		}
		if _, ok := override[k]; !ok {
			defaultOnly[k] = v
		}
	}

	return overrideOnly, defaultOnly, overrides, nil
}

func hideSecrets(out map[string]string) map[string]string {
	for k := range out {
		if containsSecret(k) {
			out[k] = fmt.Sprintf("var.%s", strings.ToUpper(k))
		}
	}

	return out
}

func containsSecret(str string) bool {
	isSecret := false

	for _, key := range secretKeys {
		if strings.Contains(strings.ToLower(str), key) {
			isSecret = true
		}
	}

	return isSecret
}
