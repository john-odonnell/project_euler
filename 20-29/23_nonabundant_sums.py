# A perfect number is a number for which the sum of its proper
# divisors is exactly equal to the number. For example, the sum
# of the proper divisors of 28 would be 1+2+4+7+14 = 28,
# which means that 28 is a perfect number.
#
# A number n is called deficient if the sum of its proper
# divisors is less than n and it is called abundant if this sum
# exceeds n.
#
# As 12 is the smallest abundant number, 1+2+3+4+6 = 16,
# the smallest number that can be written as the sum of two
# abundant numbers is 24. By mathematical analysis, it can be
# shown that all integers greater than 28123 can be written as
# the sum of two abundant numbers. However, this upper limit
# cannot be reduced any further by analysis even though it is
# known that the greatest number that cannot be expressed as
# the sum of two abundant numbers is less than this limit.
#
# Find the sum of all the positive integers which cannot be
# written as the sum of two abundant numbers.

import math

def is_abundant(n: int) -> bool:
    divs = []
    for i in range(1, math.ceil(math.sqrt(n))):
        if n % i == 0:
            divs.append(i)
            if i != 1:
                divs.append(n // i)
    if math.sqrt(n) ==  math.floor(math.sqrt(n)):
        divs.append(int(math.sqrt(n)))
    
    if sum(divs) > n:
        return True
    else:
        return False

# booleans representing each integer's abundant-ness
abundants = []
# booleans representing if each integer is a sum of abundants
sum_of_two_abundants = []
for i in range(0, 28123 + 1):
    sum_of_two_abundants.append("False")


for i in range(1, 28123 + 1):
    if is_abundant(i):
        abundants.append(i)
        for abundant in abundants:
            if (i + abundant) < 28123 + 1:
                sum_of_two_abundants[i + abundant] = "True"

nonabundant_sums = 0
for i in range(0, len(sum_of_two_abundants)):
    if sum_of_two_abundants[i] == "False":
        nonabundant_sums += i

print(nonabundant_sums)

# CORRECT
