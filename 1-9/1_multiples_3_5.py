# find the sum of the multiples of 3 and 5 that are less than 1000

import sys
n = int(sys.argv[1])

total = 0
for i in range(0, n):
    if i % 3 == 0 or i % 5 == 0:
        total += i

print(total)

# CORRECT
