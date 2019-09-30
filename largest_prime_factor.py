def check_prime(n):
    prime = True
    for i in range(2, n):
        if n % i == 0:
            prime = False
    return prime


def find_highest_prime_factor(n):
    highest_prime_factor = -1
    i = 2
    n_top_range = n
    while i < n_top_range:
        n_top_range = n / i
        if n % i == 0:
            if check_prime(int(n_top_range)):
                highest_prime_factor = n_top_range
                break
            elif check_prime(i):
                highest_prime_factor = i
                print(highest_prime_factor)
        i += 1
    return highest_prime_factor


def main():
    x = find_highest_prime_factor(4444444444)
    print(x)


main()

# CORRECT, but works very slowly for values upwards of 10^9
