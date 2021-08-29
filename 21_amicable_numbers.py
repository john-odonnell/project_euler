# Let d(n) be defined as the sum of proper divisors of n
# (numbers less than n which divide evenly into n).
# If d(a) = b and d(b) = a, where a != b, then a and b are
# an amicable pair and each of a and b are called amicable nums

# Evaluate the sum of all amicable numbers under 10000

import sys, math
n = int(sys.argv[1])

def sum_of_proper_divisors(n: int) -> int:
    divs = []
    for i in range(1, math.ceil(math.sqrt(n))):
        if n % i == 0:
            divs.append(i)
            if i != 1:
                divs.append(n // i)
    if math.sqrt(n) ==  math.floor(math.sqrt(n)):
        divs.append(int(math.sqrt(n)))
    return sum(divs)

# Dictionary, where the key is the an integer n, and the value
# is a list of integers whose sum of proper divisors is n.
#
# To find amicable numbers:
# 1. Examine each key n in the dictionary
# 2. Check if each value m in the key's list is iteslf a key
# 3. If m is a key, and n is in m's list,
#    then m and n are amicable
sums_and_nums = {}

for i in range(1, n):
    sum_of_divs = sum_of_proper_divisors(i)

    if sum_of_divs < n:
        if sum_of_divs not in sums_and_nums:
            sums_and_nums[sum_of_divs] = [i]
        else:
            nums = sums_and_nums[sum_of_divs]
            nums.append(i)
            sums_and_nums[sum_of_divs] = nums

amicable_pairs = []
for n, nums in sums_and_nums.items():
    for m in nums:
        if m in sums_and_nums and m != n:
            if n in sums_and_nums[m] and {n, m} not in amicable_pairs:
                amicable_pairs.append({n, m})

sum_of_pairs = 0
for pair in amicable_pairs:
    sum_of_pairs += sum(pair)

print(sum_of_pairs)

# CORRECT
