# Find the sum of the digits in the number 100!

import sys, math
n = int(sys.argv[1])

sum = 0
digits = str(math.factorial(n))
for i in range(0, len(digits)):
    sum += int(digits[i])

print(sum)

# CORRECT
