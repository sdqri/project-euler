import math
import time

def is_prime(number: int) -> bool:
    if number <= 1:
        return False
    max_factor = math.floor(math.sqrt(number))
    for i in range(2, max_factor+1):
        if number % i == 0:
            return False
    return True

def is_truncatable_prime(number: int) -> bool:
    if not is_prime(number) or number<10:
        return False
    str_number = str(number)
    str_length = len(str_number)
    for i in range(str_length):
        num1 = int(str_number[0:str_length-i])                
        num2 = int(str_number[i:str_length+1])                
        if not is_prime(num1) or not is_prime(num2):
            return False
    return True

if __name__ == "__main__":
    start = time.time() 
    primes: [int] = []
    for i in range(1_000_000 + 1):
        if is_truncatable_prime(i):
            primes.append(i)
    end = time.time() 
    elapsed = end - start
    print(f"primes = f{primes}, sum={sum(primes)} (elapsed={elapsed}s)")
