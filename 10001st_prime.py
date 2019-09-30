def check_prime(n):
    prime = True
    for i in range(2, n):
        if n % i == 0:
            prime = False
    return prime


def find_primes_to_n(n):
    primes = [2]
    number_to_check = 3
    while len(primes) < n:
        prime_indicator = check_prime(number_to_check)
        if prime_indicator:
            primes.append(number_to_check)
        number_to_check += 1
    return primes


def nth_prime(n, primes):
    this_prime = primes[n - 1]
    return this_prime


def main():
    x = nth_prime(10001, find_primes_to_n(10001))
    print(x)


main()

# CORRECT
