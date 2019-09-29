# find the sum of the even numbers of the fibonacci sequence
# that do not exceed 4e6


def gen_fib(n):
    fib = [1, 1]
    next_num = 0
    while next_num <= n:
        next_num = fib[len(fib) - 1] + fib[len(fib) - 2]
        fib.append(next_num)
    return fib


def sum_of_even(arr):
    total = 0
    for number in arr:
        if number % 2 == 0:
            total += number
    return total


def main():
    print(sum_of_even(gen_fib(4000000)))


main()

# CORRECT
