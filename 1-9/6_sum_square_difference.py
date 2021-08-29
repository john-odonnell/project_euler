# Find the difference between the sum of the squares of the first hundred
# natural numbers and the square of the sum


def sum_of_squares(n):
    total = 0
    for i in range(1, n + 1):
        total += i**2
    return total


def square_of_sum(n):
    total = 0
    for i in range(1, n + 1):
        total += i
    total = total**2
    return total


def main():
    difference = square_of_sum(100) - sum_of_squares(100)
    print(difference)


main()

# CORRECT
