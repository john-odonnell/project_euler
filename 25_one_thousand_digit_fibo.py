# The 12th term of the Fibonacci sequence is the first to
# contain 3 digits.
#
# What is the index of the first term in the sequence
# to contain 1000 digits?

import sys
n = int(sys.argv[1])

a = 1
b = 1
c = a + b
count = 3

while len(str(c)) < n:
    a = b
    b = c
    c = a + b
    count += 1

print(count)

# CORRECT
