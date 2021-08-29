# What is the first triangular number to have over five hundred divisors?
# general case: find the first triangular number with over n divisors

import sys
import math
n = int(sys.argv[1])

def num_divisors(n):
    """find the number of divisors of an integer"""
    count = 0
    for i in range(1, math.floor(math.sqrt(n))):
        if n % i == 0:
            count += 2
    if math.sqrt(n) - math.floor(math.sqrt(n)) == 0:
        count += 1
    return count

i = 1
tri = 1
count = num_divisors(tri)
while count <= n:
    i += 1
    tri += i
    # for our purposes, we can eliminate all odd numbers
    # from consideration. the number of divisors of even
    # numbers increases faster than that of odd numbers.
    if tri % 2 == 0:
        count = num_divisors(tri)

print(tri)

# CORRECT
