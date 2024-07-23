#  Sum square difference
# Problem 60
'''

Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
'''
import math
import time
from itertools import permutations, combinations
from tqdm import tqdm

def is_prime(num: int) -> bool:
    if num < 2:
        return False
    for i in range(2, int(math.sqrt(num)+1)):
        if num % i == 0:
            return False
    return True

def concat_nums(*args: int) -> int:
    str_num = ""
    for arg in args:
        str_num += str(arg)
    return int(str_num)

def is_prime_set(*args: int) -> bool:
    for arg in args:
        if not is_prime(arg):
            return False
    cmbs = permutations(args, 2)
    for cmb in cmbs:
        if not is_prime(concat_nums(*cmb)):
            return False
    return True

if __name__ == "__main__":
    start = time.time()
    primes: list[int] = []
    print("finding primes < 10000...")
    for i in range(1, 10000):
        if is_prime(i):
            primes.append(i)

    primes.remove(2)
    primes.remove(5)


    pairs: set[tuple[int]] = set()
    nums = set()
    print("finding prime pairs...")
    for i in tqdm(primes):
        for j in primes:
            if i>=j:
                continue
            if is_prime_set(i, j):
                pairs.add((i, j))
                nums.add(i)
                nums.add(j)


    cmbs = pairs
    for i in range(3):
        if i == 0:
            print("finding prime triples...")
        elif i== 1:
            print("finding prime quadruples...")
        elif i== 2:
            print("finding prime quintuple...")
        new_cmbs: set[tuple[int]] = set()
        for cmb in tqdm(cmbs):
            for num in nums:
                t = None
                failure = False
                for x in cmb:
                    t = tuple(sorted([x, num]))
                    if t not in pairs:
                        failure = True
                        break
                if failure == True:
                    continue
                new_cmbs.add(tuple(sorted([*cmb, num])))
        cmbs = new_cmbs
    elapsed = time.time() - start
    print(f"lowest sum for a primes quintuple = {sum(list(cmbs)[0])} (elapsed={elapsed:.6f}s)")
    print(f"combination={list(cmbs)[0]}")
 
