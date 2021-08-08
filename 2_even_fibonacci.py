# find the sum of the even numbers of the fibonacci sequence
# that do not exceed 4e6

import sys
n = int(sys.argv[1])

def gen_fib(n):
    fib = [1, 1]
    next_num = 0
    while next_num <= n:
        next_num = fib[len(fib) - 1] + fib[len(fib) - 2]
        fib.append(next_num)
    return fib

arr = gen_fib(n)

total = 0
for number in arr:
    if number % 2 == 0:
        total += number

print(total)

# CORRECT
