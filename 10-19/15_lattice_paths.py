# Starting in the top left corner of a 2x2 grid, and only being
# able to move to the right and down, there are exactly six
# routes to the bottom right corner.
# How many such routes are there through a 20x20 grid?

# input: dimensions of an n x m grid
import sys
n = int(sys.argv[1])
m = int(sys.argv[2])

# At first, my solution was going to be a stack-based solution:
# start at (0,0), and propogate through the grid (first right,
# then down) to the bottom right, then backtrack and produce
# a different path. A better way to do it:

"""
_  _  _     _  _  1     _  _  1     _  3  1     6  3  1
_  _  _     _  _  1     _  2  1     3  2  1     3  2  1
_  _  0     1  1  0     1  1  0     1  1  0     1  1  0
"""

# Points in the bottom-most row and right-most column always
# have only 1 possible path (p), and p for each other point
# is the sum of p of the points directly to the right and down
# from that point.

matrix = []
for i in range(0, n + 1):
    row = []
    for j in range(0, m + 1):
        row.append(0)
    matrix.append(row)

for i in range(0, n):
    matrix[i][m] = 1
for j in range(0, m):
    matrix[n][j] = 1

for i in range(n - 1, -1, -1):
    for j in range(m - 1, -1, -1):
        matrix[i][j] = matrix[i+1][j] + matrix[i][j+1]

print(matrix[0][0])

# CORRECT
