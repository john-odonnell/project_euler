# The following iterative sequence is defined for the set of positive integers:
#    n -> n/2    (n is even)
#    n -> 3n + 1 (n is odd)
# Which starting number, under one million, produces the longest chain?
# Once the chain starts the terms are allowed to go above one million

import sys
n = int(sys.argv[1])

def brute_force(n):
    chain_lengths = [-1]
    for i in range(1, n):
        # only compute chains for odd numbers, as even numbered
        # chains will always be log(n) long, and won't compete
        # for the longest length chain
        if i % 2 == 0:
            chain_lengths.append(0)
        else:
            j = i
            count = 1
            while j != 1:
                if j % 2 == 0:
                    j = int(j / 2)
                else:
                    j = (3 * j) + 1
                count += 1
            chain_lengths.append(count)
    max_chain_length = max(chain_lengths)
    max_chain_index = chain_lengths.index(max_chain_length)
    print(max_chain_index)
    print(max_chain_length)

brute_force(n)

# CORRECT
