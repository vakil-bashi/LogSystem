package date

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/timeutil"
	"github.com/vakil-bashi/log-system/logger"
	"github.com/vakil-bashi/log-system/utils/responses"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

var (
	DateTimeService dateTimeServiceInterface = &dateTimesService{}
)

type dateTimesService struct{}

type dateTimeServiceInterface interface {
	GetNow() time.Time
	GetNowString() string
	GetNowDBFormat() string
	ConvertToDate(string) time.Time
	Date(int, int, int) time.Time
	DeltaTime(string, string) (*int, *responses.Response)
	DateIsPassed(string) *responses.Response
	IsExpired(string) *responses.Response
	IsNotExpired(string) *responses.Response
}

func (s *dateTimesService) GetNow() time.Time {
	return time.Now().UTC()
}

func (s *dateTimesService) GetNowString() string {
	return s.GetNow().Format(apiDateLayout)
}

func (s *dateTimesService) Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (s *dateTimesService) DeltaTime(checkInDate string, checkOutDate string) (*int, *responses.Response) {

	checkIn := strings.Split(checkInDate, "-")

	checkInDay, _ := strconv.ParseInt(checkIn[2], 10, 64)
	checkInMonth, _ := strconv.ParseInt(checkIn[1], 10, 64)
	checkInYear, _ := strconv.ParseInt(checkIn[0], 10, 64)

	checkOut := strings.Split(checkOutDate, "-")

	checkOutDay, _ := strconv.ParseInt(checkOut[2], 10, 64)
	checkOutMonth, _ := strconv.ParseInt(checkOut[1], 10, 64)
	checkOutYear, _ := strconv.ParseInt(checkOut[0], 10, 64)

	t1 := s.Date(int(checkInYear), int(checkInMonth), int(checkInDay))
	t2 := s.Date(int(checkOutYear), int(checkOutMonth), int(checkOutDay))
	days := int(t2.Sub(t1).Hours() / 24)

	return &days, nil
}

func (s *dateTimesService) DateIsPassed(date string) *responses.Response {
	base, err := time.Parse("2006-01-02", date)
	if err != nil {
		logger.Error("error when trying to parse date", err)
		return responses.NewInternalServerError("Internal server error", "Please Try again later...")
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         1,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      0,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result >= 1 {
		logger.Error("date is passed", err)
		return responses.NewBadRequestError("Date is Passed", "Check-in Date at least must be today!", http.StatusBadRequest)
	}
	return nil
}

func (s *dateTimesService) IsExpired(dateVerified string) *responses.Response {
	logger.Info("Enter to Is-Expired function")

	dateVerified = strings.Replace(dateVerified, " ", "T", -1)
	dateVerified += "Z"

	base, err := time.Parse(apiDateLayout, dateVerified)
	if err != nil {
		logger.Error("error when trying to parse date", err)
		return responses.NewBadRequestError("datetime format is invalid", "please make sure to enter correct format of datetime (2006-01-02T15:04:05Z)", http.StatusBadRequest)
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         0,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      1,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result >= 1 {
		logger.Error("verification code is expired", err)
		return responses.NewBadRequestError("verification code is expired", "Please try again", http.StatusBadRequest)
	}
	logger.Info("Close from Is-Expired function successfully!")
	return nil
}

func (s *dateTimesService) IsNotExpired(dateVerified string) *responses.Response {
	logger.Info("Enter to Is-Not-Expired function")

	dateVerified = strings.Replace(dateVerified, " ", "T", -1)
	dateVerified += "Z"

	base, err := time.Parse(apiDateLayout, dateVerified)
	if err != nil {
		logger.Error("error when trying to parse date ", err)
		return responses.NewBadRequestError("datetime format is invalid", "please make sure to enter correct format of datetime (2006-01-02T15:04:05Z)", http.StatusBadRequest)
	}

	oneMinutesLater := timeutil.Timedelta{
		Days:         0,
		Seconds:      0,
		Microseconds: 0,
		Milliseconds: 0,
		Minutes:      1,
		Hours:        0,
		Weeks:        0,
	}

	base = base.Add(oneMinutesLater.Duration())
	current := time.Now().UTC()
	result := current.Sub(base).Minutes()

	if result < 1 {
		logger.Error("verification code is not expired", err)
		return responses.NewBadRequestError("verification code is not expired yet", "please try again after at least 2 minutes.", http.StatusBadRequest)
	}
	logger.Info("Close from Is-Not-Expired function successfully")
	return nil
}

func (s *dateTimesService) GetNowDBFormat() string {
	ret := s.GetNow().Format(apiDbLayout)
	return ret
}

func (s *dateTimesService) ConvertToDate(dateTime string) time.Time {
	t, err := time.Parse(apiDateLayout, dateTime)

	if err != nil {
		fmt.Println(err)
	}
	return t
}
