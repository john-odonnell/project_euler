# 2^15 = 32768 and the sum of its digits is 26
# What is the sum of digits of the number 2^1000?

import sys
n = int(sys.argv[1])

sum = 0
two_pow_n_str = str(pow(2, n))
for i in range(0, len(two_pow_n_str)):
    sum += int(two_pow_n_str[i])

print(sum)

# CORRECT
