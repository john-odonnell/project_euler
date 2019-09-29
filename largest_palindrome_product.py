# find the largest palindrome made from the product of
# two three digit numbers


def int_to_rev_array(n):
    rev_array = []
    while n > 0:
        rev_array.append(n % 10)
        n = int(n / 10)
    return rev_array


def reverse_list(array):
    rev_array = []
    i = len(array) - 1
    while i >= 0:
        rev_array.append(array[i])
        i -= 1
    return rev_array


def find_palindrome_product(n):
    highest_palindrome_product = 0
    for i in range((10**(n-1)), (10**n)):
        for j in range((10**(n-1)), (10**n)):
            product = i * j
            rev_product = int_to_rev_array(product)
            for_product = reverse_list(rev_product)
            if product > highest_palindrome_product and rev_product == for_product:
                highest_palindrome_product = product
    return highest_palindrome_product


def main():
    print(find_palindrome_product(3))


main()

# CORRECT
