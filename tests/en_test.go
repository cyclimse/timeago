package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
	. "github.com/SerhiiCho/timeago/utils"
)

func TestConvEn(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{SmallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{SmallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{SmallSubTime(-5 * time.Minute), "5 minutes ago", "en"},
		{SmallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{SmallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{SmallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{SmallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{SmallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{SmallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{SmallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{SmallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{SmallSubTime(-59 * time.Minute), "59 minutes ago", "en"},
		{SmallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{SmallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{SmallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{SmallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{SmallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{SmallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{SmallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{SmallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{SmallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{SmallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{SmallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{BigSubTime(0, 1, 1), "1 month ago", "en"},
		{BigSubTime(0, 2, 1), "2 months ago", "en"},
		{BigSubTime(0, 9, 1), "9 months ago", "en"},
		{BigSubTime(0, 11, 1), "11 months ago", "en"},
		{BigSubTime(0, 12, 1), "1 year ago", "en"},
		{BigSubTime(1, 0, 1), "1 year ago", "en"},
		{BigSubTime(2, 0, 1), "2 years ago", "en"},
		{BigSubTime(21, 0, 1), "21 years ago", "en"},
		{BigSubTime(31, 0, 1), "31 years ago", "en"},
		{BigSubTime(100, 0, 1), "100 years ago", "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Conv(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestConvEnWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), "Online", "en"},
		{SmallSubTime(time.Second), "Online", "en"},
		{SmallSubTime(-1 * time.Second), "Online", "en"},
		{SmallSubTime(-2 * time.Second), "Online", "en"},
		{SmallSubTime(-9 * time.Second), "Online", "en"},
		{SmallSubTime(-10 * time.Second), "Online", "en"},
		{SmallSubTime(-11 * time.Second), "Online", "en"},
		{SmallSubTime(-20 * time.Second), "Online", "en"},
		{SmallSubTime(-21 * time.Second), "Online", "en"},
		{SmallSubTime(-22 * time.Second), "Online", "en"},
		{SmallSubTime(-30 * time.Second), "Online", "en"},
		{SmallSubTime(-31 * time.Second), "Online", "en"},
		{SmallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{SmallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{SmallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{SmallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{SmallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{SmallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{SmallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{SmallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{SmallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{SmallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{SmallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{SmallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{SmallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{SmallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{SmallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{SmallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{SmallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{SmallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{SmallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{SmallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{SmallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{SmallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{BigSubTime(0, 1, 1), "1 month ago", "en"},
		{BigSubTime(0, 2, 1), "2 months ago", "en"},
		{BigSubTime(0, 9, 1), "9 months ago", "en"},
		{BigSubTime(0, 11, 1), "11 months ago", "en"},
		{BigSubTime(0, 12, 1), "1 year ago", "en"},
		{BigSubTime(1, 0, 1), "1 year ago", "en"},
		{BigSubTime(2, 0, 1), "2 years ago", "en"},
		{BigSubTime(21, 0, 1), "21 years ago", "en"},
		{BigSubTime(31, 0, 1), "31 years ago", "en"},
		{BigSubTime(100, 0, 1), "100 years ago", "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date+"|online", func(test *testing.T) {
			timeago.Set("language", tc.lang)

			if res := timeago.Conv(tc.date + "|online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestConvEnWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), []string{"0 seconds ago", "1 second ago"}, "en"},
		{SmallSubTime(time.Second), []string{"0 seconds ago", "1 second ago"}, "en"},
		{SmallSubTime(-1 * time.Second), []string{"1 second ago", "2 seconds ago"}, "en"},
		{SmallSubTime(-2 * time.Second), []string{"2 seconds ago", "3 seconds ago"}, "en"},
		{SmallSubTime(-9 * time.Second), []string{"9 seconds ago", "10 seconds ago"}, "en"},
		{SmallSubTime(-10 * time.Second), []string{"10 seconds ago", "11 seconds ago"}, "en"},
		{SmallSubTime(-11 * time.Second), []string{"11 seconds ago", "12 seconds ago"}, "en"},
		{SmallSubTime(-20 * time.Second), []string{"20 seconds ago", "21 seconds ago"}, "en"},
		{SmallSubTime(-21 * time.Second), []string{"21 seconds ago", "22 seconds ago"}, "en"},
		{SmallSubTime(-22 * time.Second), []string{"22 seconds ago", "23 seconds ago"}, "en"},
		{SmallSubTime(-30 * time.Second), []string{"30 seconds ago", "31 seconds ago"}, "en"},
		{SmallSubTime(-31 * time.Second), []string{"31 seconds ago", "32 seconds ago"}, "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Conv(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
