# Sum square difference
# Problem 6
'''

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

'''

import logging

logging.basicConfig(level=logging.DEBUG)

n = int(input("Enter n = "))
sum_of_the_squares = sum(n ** 2 for n in range(1, n + 1))
square_of_the_sum = sum(n for n in range(1, n + 1)) ** 2
logging.debug(square_of_the_sum - sum_of_the_squares)
