#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void printArray(int *arr, int length) {
    printf("[ ");
    for (int i = 0; i < length; i++) {
        printf("%d ", arr[i]);
    }
    printf("]\n");
}

int factorial(int x) {
    if (x == 0 || x == 1) {
        return 1;
    }
    return x * factorial(x - 1);
}

void get_all_permutatins_inner(int *digits, int digits_length, int *perumtation,
                               int permutation_length, int **permutations,
                               int *permutations_length) {
    if (digits_length == 1) {
        perumtation[permutation_length] = digits[0];
        permutations[*permutations_length] =
            malloc((permutation_length + 1) * sizeof(int));
        memcpy(permutations[*permutations_length], perumtation,
               (permutation_length + 1) * sizeof(int));
        (*permutations_length)++;
        return;
    }

    for (int i = 0; i < digits_length; i++) {
        int *new_rotation = malloc((permutation_length + 1) * sizeof(int));
        memcpy(new_rotation, perumtation, permutation_length * sizeof(int));
        new_rotation[permutation_length] = digits[i];

        int *new_digits = malloc((digits_length - 1) * sizeof(int));
        int index = 0;
        for (int j = 0; j < digits_length; j++) {
            if (j != i) {
                new_digits[index++] = digits[j];
            }
        }

        get_all_permutatins_inner(new_digits, digits_length - 1, new_rotation,
                                  permutation_length + 1, permutations,
                                  permutations_length);

        free(new_rotation);
        free(new_digits);
    }
}

void get_all_permutations(int *digits_in, int digits_length_in,
                          int ***permutations_out,
                          int *permutations_length_out) {
    int **permutations = malloc(factorial(digits_length_in) * sizeof(int *));

    get_all_permutatins_inner(digits_in, digits_length_in, NULL, 0,
                              permutations, permutations_length_out);

    *permutations_out = permutations;
}

void get_all_combinations(int *digits_in, int digits_length_in,
                          int selection_length_in, int ***combinations_out,
                          int *combinations_length_out) {
    *combinations_length_out =
        factorial(digits_length_in) /
        (factorial(selection_length_in) *
         factorial(digits_length_in - selection_length_in));

    *combinations_out = malloc(*combinations_length_out * sizeof(int *));

    int combinations_index = 0;
    for (int bitmask = 0; bitmask < (1 << digits_length_in); bitmask++) {
        if (__builtin_popcount(bitmask) == selection_length_in) {
            int *combination = malloc(selection_length_in * sizeof(int));
            int index = 0;
            for (int i = 0; i < digits_length_in; i++) {
                if (bitmask & (1 << i)) {
                    combination[index++] = digits_in[i];
                }
            }
            (*combinations_out)[combinations_index++] = combination;
        }
    }
}

void get_all_combinations_with_repetition_inner(
    int *digits_in, int digits_length_in, int selection_length_in,
    int *combination, int combination_length, int **combinations,
    int *combinations_length, int min_i) {
    if (combination_length == selection_length_in) {
        combinations[*combinations_length] =
            malloc((combination_length + 1) * sizeof(int));
        memcpy(combinations[*combinations_length], combination,
               (combination_length + 1) * sizeof(int));
        (*combinations_length)++;
        return;
    }

    if (combinations == NULL) {
        combinations = malloc(selection_length_in * sizeof(int));
        combinations_length = 0;
    }

    for (int i = min_i; i < digits_length_in; i++) {

        int *new_combination = malloc(selection_length_in * sizeof(int));
        memcpy(new_combination, combination, combination_length * sizeof(int));
        new_combination[combination_length] = digits_in[i];

        get_all_combinations_with_repetition_inner(
            digits_in, digits_length_in, selection_length_in, new_combination,
            combination_length + 1, combinations, combinations_length, i);
    }
}

void get_all_combinations_with_repetition(int *digits_in, int digits_length_in,
                                          int selection_length_in,
                                          int ***combinations_out,
                                          int *combinations_length_out) {
    int combinations_length =
        factorial(digits_length_in + selection_length_in - 1) /
        (factorial(selection_length_in) * factorial(digits_length_in - 1));

    int **combinations = malloc(combinations_length * sizeof(int *));

    get_all_combinations_with_repetition_inner(
        digits_in, digits_length_in, selection_length_in, NULL, 0, combinations,
        combinations_length_out, 0);

    *combinations_out = combinations;
    *combinations_length_out = combinations_length;
}

int intPow(int base, int exponent) {
    if (exponent == 0) {
        return 1;
    }
    int result = 1;
    while (exponent > 0) {
        result *= base;
        exponent--;
    }
    return result;
}

int array_to_number(int *arr, int arr_length) {
    int result = 0;
    for (int i = arr_length - 1; i >= 0; i--) {
        result += arr[i] * (intPow(10, arr_length - 1 - i));
    }
    return result;
}

int count_three_equally_distant_numbers(int arr[], int n) {
    int count = 0;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            for (int k = j + 1; k < n; k++) {
                if (arr[i] - arr[j] == arr[j] - arr[k] &&
                    arr[i] - arr[j] != 0) {
                    printf("%d, %d, %d\n", arr[i], arr[j], arr[k]);
                    count++;
                }
            }
        }
    }
   return count;
}

bool is_prime(int num) {
    if (num <= 1) {
        return false;
    }
    if (num <= 3) {
        return true;
    }
    if (num % 2 == 0 || num % 3 == 0) {
        return false;
    }
    for (int i = 5; i * i <= num; i += 6) {
        if (num % i == 0 || num % (i + 2) == 0) {
            return false;
        }
    }
    return true;
}

bool in_array(int element, int *arr, int len) {
    for (int i = 0; i < len; i++) {
        if ((arr[i]) == element) {
            return true;
        }
    }
    return false;
}

int main(int argc, char *argv[]) {
    clock_t start = clock();
    int digits[] = {1, 2, 3, 4, 5, 6, 7, 8, 9};
    int digits_length = sizeof(digits) / sizeof(digits[0]);
    int **combinations;
    int combinations_length = 0;

    int r = 4;
    get_all_combinations_with_repetition(digits, digits_length, r,
                                         &combinations, &combinations_length);

    for (int i = 0; i < combinations_length; i++) {
        int **permutations;
        int permutations_length = 0;
        get_all_permutations(combinations[i], r, &permutations,
                             &permutations_length);

        int *prime_numbers = NULL;
        int prime_count = 0;
        for (int j = 0; j < permutations_length; j++) {
            int num = array_to_number(permutations[j], r);
            if (!in_array(num, prime_numbers, prime_count) && is_prime(num)) {
                prime_numbers =
                    realloc(prime_numbers, (prime_count + 1) * sizeof(int));
                prime_numbers[prime_count++] = num;
            }
        }
        int count_sequence =
            count_three_equally_distant_numbers(prime_numbers, prime_count);
        if (prime_count >= 3 && count_sequence > 0) {
            printf("found %d sequence in this combination : ", count_sequence);
            printArray(combinations[i], r);
        }

        for (int j = 0; j < permutations_length; j++) {
            free(permutations[j]);
        }
        free(permutations);

        free(prime_numbers);
    }

    for (int i = 0; i < combinations_length; i++) {
        free(combinations[i]);
    }
    free(combinations);

    clock_t stop = clock();
    double elapsed = (double)(stop - start) / CLOCKS_PER_SEC;
    printf("( elapsed = %fs )\n", elapsed);
}
