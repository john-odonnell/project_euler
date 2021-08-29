# Using names.txt, a 46K text file containing over 5000
# first names, begin by sorting it into alphabetical order.
# Then working out the alphabetical value for each name,
# multiply this value by its alphabetical position in the list
# to obtain a name score.
#
# What is the total of all the name scores in the file?

names_file = open("names.txt", "r")
names_str = names_file.readline()
names = names_str.replace('"', "").split(",")
names.sort()

alphabetical_values = []
for name in names:
    name_sum = 0
    for letter in name:
        name_sum += (ord(letter) - 64)
    alphabetical_values.append(name_sum)


alphabetical_scores = []
for i in range(0, len(names)):
    alphabetical_scores.append(alphabetical_values[i] * (i+1))

print(sum(alphabetical_scores))

# CORRECT
