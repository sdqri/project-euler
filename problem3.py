#Largest prime factor
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
	for i in range(2, n):
		if n%i==0:
			return False
	return True

def find_prime_factors(n):
	l = []
	for i in range(2, n):
		if(is_prime(i)):
			l.append(i)
	return l
		
number = 600851475143
logging.debug(find_prime_factors(number))
		
