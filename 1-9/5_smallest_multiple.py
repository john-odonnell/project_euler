# find the smallest positive number that is evenly divisible by
# all of the numbers from 1 to 20


def find_smallest_multiple(n):
    i = 1
    found = False
    while not found:
        passed = 0
        for j in range(1, n + 1):
            if i % j == 0:
                passed += 1
        if passed == n:
            found = True
        else:
            i += 1
    return i


def main():
    print(find_smallest_multiple(20))


main()

# CORRECT
