def check_prime(n):
    prime = True
    for i in range(2, n):
        if n % i == 0:
            prime = False
    return prime


def sum_of_primes_below_n(n):
    primes_below_n = [2]
    for i in range(3, n, 2):
        if check_prime(i):
            primes_below_n.append(i)
            print(i)

    total = 0
    for prime in primes_below_n:
        total += prime

    return total


def main():
    print(sum_of_primes_below_n(2000000))


main()

# I assume I got it correct
