# What is the largest prime factor of the number 600851475143?

import sys
n = int(sys.argv[1])

# establish each possible factor's prime-ness
is_prime = []
for i in range(0, (n//2) + 1):
    is_prime.append(True)
for i in range(2, len(is_prime)+1):
    multiplier = 2
    while multiplier * i < len(is_prime):
        is_prime[multiplier * i] = False
        multiplier += 1

highest_prime_factor = -1
for i in range(2, len(is_prime)):
    if n % i == 0 and is_prime[i]:
        highest_prime_factor = i

print(highest_prime_factor)

# CORRECT, but works very slowly for values upwards of 10^7
