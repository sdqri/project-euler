# Largest palindrome product
# Problem 4
'''

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

'''

import logging
import math
from functools import reduce

logging.basicConfig(level = logging.DEBUG)

def is_palindrome(n: int)->bool:
    if(n==int(str(n)[::-1])):
        return True
    else:
        return False

def find_largest_palindrome():
    palindromes = []
    for i in range(100, 1000):
        for j in range(999, i, -1):
            if(is_palindrome(i*j)):
                palindromes.append((i, j, i*j))
    return palindromes

palindromes = find_largest_palindrome()
lp = sorted(palindromes, key=lambda x:x[2], reverse=True)[0]
logging.debug(lp)
