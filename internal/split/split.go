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

func Do(left map[string]string, right map[string]string) (map[string]string, map[string]string, map[string]string, error) {
	baseCfg := hideSecrets(left)
	targetCfg := hideSecrets(right)

	rightOnly := map[string]string{}
	leftOnly := map[string]string{}
	overwrites := map[string]string{}

	for k, v := range targetCfg {
		if _, ok := baseCfg[k]; !ok {
			rightOnly[k] = v
			continue
		}

		if baseCfg[k] != targetCfg[k] {
			overwrites[k] = v
		}
	}

	for k, v := range baseCfg {
		if _, ok := targetCfg[k]; !ok {
			leftOnly[k] = v
		}
	}

	return rightOnly, leftOnly, overwrites, nil
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
