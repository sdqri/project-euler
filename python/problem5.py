# Smallest multiple
# Problem 5
'''

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

'''

import logging
import operator
from functools import reduce
from collections import Counter

logging.basicConfig(level = logging.DEBUG)

def factor(x):
    f = Counter()
    f[1] += 1
    i = 2
    while(x != 1):
        if(x%i==0):
            x = x/i
            f[i] += 1
        else:
            i += 1
    return f

all_factors = [factor(n) for n in range(1, 21)]
smallest_factors = {}
for factors in all_factors:
    for k,v in factors.items():
        if k not in smallest_factors or (k in smallest_factors and smallest_factors[k]<v):
            smallest_factors[k] = v

smallest_multiple = reduce(operator.mul, [k**v for k,v in smallest_factors.items()])
logging.debug(smallest_multiple)
