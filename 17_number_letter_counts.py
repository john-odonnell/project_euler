# If the numbers 1 to 5 are written out in words, then there
# are 19 letters used in total

# If all the numbers from 1 to 1000 (inclusive) were written
# out in words, how many letters would be used?

import sys
n = int(sys.argv[1])

digits     = [4, 3, 3, 5, 4, 4, 3, 5, 5, 4]
teens      = [3, 6, 6, 8, 8, 7, 7, 9, 8, 8]
tens       = [0, 3, 6, 6, 5, 5, 5, 7, 6, 6]
pow_of_ten = [0, 7, 8]

sum_letters = 0
for i in range(0, n + 1):
    written = []
    thousands_place = False
    hundreds_place = False
    tens_place = False
    ones_place = False

    n = (i // 1000) % 10
    if n > 0:
        sum_letters += digits[n]
        sum_letters += pow_of_ten[2]
        thousands_place = True
    n = (i //  100) % 10
    if n > 0:
        sum_letters += digits[n]
        sum_letters += pow_of_ten[1]
        hundreds_place = True
    n = (i //   10) % 10
    if n > 0:
        tens_place = True
    if n >= 0:
        if n == 1:
            m = (i //    1) % 10
            sum_letters += teens[m]
            if m > 0:
                ones_place = True
        else:
            sum_letters += tens[n]
            n = (i //    1) % 10
            if n > 0:
                sum_letters += digits[n]
                ones_place = True

    if (hundreds_place or thousands_place) and \
       (tens_place or ones_place):
        sum_letters += 3

print(sum_letters)

# CORRECT
