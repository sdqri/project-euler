# Maximum Path Sum I
# Problem 18

# triangle = [
#     [75],
#     [95, 64],
#     [17, 47, 82],
#     [18, 35, 87, 10],
#     [20, 4, 82, 47, 65],
#     [19, 1, 23, 75, 3, 34],
#     [88, 2, 77, 73, 7, 63, 67],
#     [99, 65, 4, 28, 6, 16, 70, 92],
#     [41, 41, 26, 56, 83, 40, 80, 70, 33],
#     [41, 48, 72, 33, 47, 32, 37, 16, 94, 29],
#     [53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14],
#     [70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57],
#     [91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48],
#     [63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31],
#     [4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23],
# ]
#
# smol_triangle = [
#     [3],
#     [7, 4],
#     [2, 4, 6],
#     [8, 5, 9, 3],
# ]
#
from functools import lru_cache

triangle: list[int] = None

@lru_cache
def find_best_score(target_index: int, line_number: int) -> int:
    global triangle
    line = triangle[line_number]
    if len(line) == 1:
        return line[0]
    if target_index < len(line) - 1:
        p1 = target_index
    else:
        p1 = target_index - 1
    if target_index - 1 >= 0:
        p2 = target_index - 1
    else:
        p2 = target_index
    return line[target_index] + max(
        [
            find_best_score(p1, line_number - 1),
            find_best_score(p2, line_number - 1),
        ]
    )


if __name__ == "__main__":
    triangle = []
    with open("./0067_triangle.txt") as fp:
        lines = fp.readlines()
        for line in lines:
            triangle.append([int(x) for x in line.split(" ")])

    # print(triangle)

    results = []
    for i, _ in enumerate(triangle[-1]):
        results.append(find_best_score(i, len(triangle) - 1))
    print(f"best score = {max(results)}")
