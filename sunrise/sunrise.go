package sunrise

import (
	//"fmt"
	"time"
	"math"
	//"strconv"
)

func JulianDayNumber(t time.Time) int {
  // JDN = (1461 × (Y + 4800 + (M − 14)/12))/4 +(367 × (M − 2 − 12 × ((M − 14)/12)))/12 − (3 × ((Y + 4900 + (M - 14)/12)/100))/4 + D − 32075
  year := t.Year()
  month := int(t.Month())
  day := t.Day()
  return (1461*(year + 4800 + (month - 14)/12))/4 + (367*(month - 2 - 12*((month -14)/12)))/12 - (3*((year + 4900 + (month-14)/12)/100))/4 + day - 32075
}

func JulianDate(t time.Time) float64 {
  return float64(JulianDayNumber(t)) + float64(t.Hour() - 12)/24 + float64(t.Minute())/1440 + float64(t.Second())/86400
}

func JulianDateToTime(jd float64) (t time.Time) {
 return t
}

func CurrentJulianDate(t time.Time) float64 {
  return math.Round(JulianDate(t) - 2451545.00 + 0.0008)
}

func MeanSolarNoon(t time.Time, longitudeWest float64) float64 {
  return CurrentJulianDate(t) - longitudeWest/360
}

func SolarMeanAnomly(t time.Time, longitudeWest float64) float64 {
 return math.Mod(357.5291 + 0.98560028*MeanSolarNoon(t, longitudeWest), 360)
}

func EquationOfCenter(t time.Time, longitudeWest float64) float64 {
  ma := SolarMeanAnomly(t, longitudeWest)
  return 1.9148*math.Sin(ma) + 0.02*math.Sin(2*ma) + 0.0003*math.Sin(3*ma)
}

func EclipticLongitude(t time.Time, longitudeWest float64) float64 {
  return math.Mod(SolarMeanAnomly(t, longitudeWest) + EquationOfCenter(t, longitudeWest) + 180 + 102.9372, 360)
}

func SolarTransit(t time.Time, longitudeWest float64) float64 {
  return 2451545 + MeanSolarNoon(t, longitudeWest) + 0.0053*math.Sin(SolarMeanAnomly(t, longitudeWest)) - 0.0069*math.Sin(2*EclipticLongitude(t, longitudeWest))
}

func DeclinationOfTheSun(t time.Time, longitudeWest float64) float64 {
  return math.Asin(math.Sin(EclipticLongitude(t, longitudeWest))*math.Sin(23.44))
}

func HourAngle(t time.Time, longitudeWest, latitudeNorth float64) float64 {
  d := DeclinationOfTheSun(t, longitudeWest)
  return math.Acos((math.Sin(-0.83)-math.Sin(latitudeNorth)*math.Sin(d))/math.Cos(latitudeNorth)*math.Cos(d))
}

func SunRise(t time.Time, longitudeWest, latitudeNorth float64) float64 {
	ha := HourAngle(t, longitudeWest, latitudeNorth)
	transit := SolarTransit(t, longitudeWest)
	return transit - ha/360
}

func SunSet(t time.Time, longitudeWest, latitudeNorth float64) float64 {
	ha := HourAngle(t, longitudeWest, latitudeNorth)
	transit := SolarTransit(t, longitudeWest)
	return transit + ha/360
}
