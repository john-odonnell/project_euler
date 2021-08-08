# Find the sum of all primes below two million

import sys
n = int(sys.argv[1])

# establish each number's prime-ness
is_prime = []
for i in range(0, n):
    is_prime.append(True)
for i in range(2, len(is_prime) / 2 + 1):
    multiplier = 2
    while multiplier * i < len(is_prime):
        is_prime[multiplier * i] = False
        multiplier += 1

sum = 0
for i in range(2, n):
    if is_prime[i]:
        sum += i

print(sum)

# CORRECT
