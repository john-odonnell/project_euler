# A permutation is an ordered arrangement of objects.
# For example, 3124 is one possible permutation of the digits
# 1, 2, 3 and 4. If all of the permutations are listed
# numerically or alphabetically, we call it lexicographic
# order. The lexicographic permutations of 0, 1 and 2 are:
#
# 012   021   102   120   201   210
#
# What is the millionth lexicographic permutation of the digits
# 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

# There are 10! (or 10 * 9!) permutations of these digits

# There are 9! permutations of the back 9 digits, and 
# 2 * 9! = 725760 is the closest it can get without
# passing 1000000, so we know the first digit is 2

# As 8! = 40320, and given that 2 * 9! = 725760,
# (2 * 9!) + (6 * 8!) = 967680 is the closest we can get without
# passing 1000000. Since 2 is spoken for, the permutation is
# 2 7 _ _ _ _ _ _ _ _

# This will then be applied to 7! and so on to zero in on the
# millionth permutation.

import sys, math, datetime
nth_permutation = int(sys.argv[1]) - 1

digits = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
permutation = []
counter = 0

begin = datetime.datetime.now()
while len(permutation) < 9:
    n = 0
    while (counter + ((n+1) * math.factorial(len(digits) - 1))) <= nth_permutation:
        n += 1
    permutation.append(str(digits[n]))
    counter += (n * math.factorial(len(digits) - 1))
    digits.pop(digits.index(digits[n]))

permutation += str(digits[0])
end = datetime.datetime.now()

print("".join(permutation))
print(end - begin)

# CORRECT
# ~30 microseconds on my laptop
