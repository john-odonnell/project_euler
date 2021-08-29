# You are given the following information, but may prefer to do
# some research for yourself.
# - 1 Jan 1900 was a Monday
# - 30 days has September, April, June and November
#   All the rest have 31,
#   saving February which has 28, rain or shine
#   and on leap years, 29
# - A leap year occurs on any year evenly divisible by 4, but
#   not on a century unless it is divisible by 400
# How many Sundays fell on the 1st of the month during the
# twentieth century (1 Jan 1901 to 31 Dec 2000)?

from calendar import isleap

days_in_month_stnd = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
days_in_month_leap = [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

count = 0

current_day = 1
current_month = 0
current_year = 1900

days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"]
months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]

while current_year != 2001:
    if current_day == 0 and current_year >= 1901:
        count += 1

    if isleap(current_year):
        current_day = (current_day + days_in_month_leap[current_month]) % 7
    else:
        current_day = (current_day + days_in_month_stnd[current_month]) % 7

    if current_month == 11:
        current_year += 1
        current_month = 0
    else:
        current_month += 1

print(count)

# CORRECT
