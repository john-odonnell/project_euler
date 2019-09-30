def special_pythagorean_triplet():
    product = 0
    c = 3
    triple_found = False
    while not triple_found:
        for b in range(1, c):
            for a in range(1, b):
                if ((a**2) + (b**2)) == (c**2):
                    if a + b + c == 1000:
                        product = a * b * c
                        triple_found = True
        c += 1
    return product


def main():
    print(special_pythagorean_triplet())


main()

# CORRECT
