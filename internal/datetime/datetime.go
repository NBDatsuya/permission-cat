package datetime

import "time"

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculatedTime(baseTime time.Time, durationStr string) (time.Time, error) {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return time.Time{}, err
	}

	return baseTime.Add(duration), nil
}
