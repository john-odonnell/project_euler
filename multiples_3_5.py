# find the sum of the multiples of 3 and 5 that are less than 1000


def sum_of_multiples(n):
    total = 0
    for i in range(0, n):
        if i % 3 == 0 or i % 5 == 0:
            total += i
    return total


def main():
    print(sum_of_multiples(1000))


main()

# CORRECT
