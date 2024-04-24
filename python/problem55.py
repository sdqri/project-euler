import time

def reverse_int(x: int)-> int:
    return int(str(x)[::-1])

def is_palindromic(x: int)-> bool:
    if len(str(x))<2:
        return False
    return str(x) == str(x)[::-1]

def is_lychrel_number(x: int)-> bool:
    for i in range(50):
        x = x + reverse_int(x)
        if is_palindromic(x):
            return False
    return True

if __name__ == "__main__":
    start = time.time()
    count_lychrel_numbers = 0
    n = 10000
    for i in range(1, 10000):
        if is_lychrel_number(i):
            count_lychrel_numbers += 1
    elapsed = time.time() - start
    print(f"lychrel numbers less than {n} = {count_lychrel_numbers} (elapsed = {elapsed:.6f}s)")
