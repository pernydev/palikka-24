package cooldown

import "time"

var Cooldowns = make(map[string]time.Time)

func SetCooldown(ip string, id string) {
	setCooldown(ip, 7)
	setCooldown(id, 7)
}

func setCooldown(key string, seconds int) {
	Cooldowns[key] = time.Now().Add(time.Duration(seconds) * time.Second)
}

func IsOnCooldown(ip string, id string) bool {
	return isOnCooldown(ip) || isOnCooldown(id)
}

func isOnCooldown(key string) bool {
	if time.Now().Before(Cooldowns[key]) {
		return true
	}
	delete(Cooldowns, key)
	return false
}

func CleanUp() {
	for key, value := range Cooldowns {
		if time.Now().After(value) {
			delete(Cooldowns, key)
		}
	}
}
