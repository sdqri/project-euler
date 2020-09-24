#largest prime factor
#Problem 3
'''

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

'''
import logging
import math
from functools import reduce

logging.basicConfig(level = logging.DEBUG)

def is_prime(n):
    for i in range(2, int(math.sqrt(n))):
        if n%i==0:
            return False
    return True

def find_largest_prime_factor(n):
    for i in range(int(math.sqrt(n)), 1, -1):
        if(n%i==0 and is_prime(i)):
            return i
    return None
        
number = 600851475143
logging.debug(find_largest_prime_factor(number))
