//Package leap provides a function to calculate if a year is a Leap Year in the Gregorian Calendar
package leap

// IsLeapYear will return if a given year in Gregorian calendar is a leap year
func IsLeapYear(i int) bool {
	return (i%4 == 0 && i%100 != 0) || i%400 == 0
}
