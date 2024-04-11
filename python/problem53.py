import math
import time


def combination(n: int, r: int):
    return math.factorial(n) / (math.factorial(r) * math.factorial(n - r))


if __name__ == "__main__":
    start = time.time() 
    count_combinations: int = 0
    for n in range(1, 101):
        half_n = n // 2
        if combination(n, half_n) > 1_000_000:
            for r in range(1, n + 1):
                if combination(n, r) > 1_000_000:
                    count_combinations += 1
    end = time.time()
    elapsed = end - start
    print(f"combinations = {count_combinations} (elapsed={elapsed:.6f}s)")
