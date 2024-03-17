#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

void printArray(int* arr, int length) {
    printf("[ ");
    for (int i = 0; i < length; i++) {
        printf("%d ", arr[i]);
    }
    printf("]\n");
}

int factorial(int x){
    if (x==0||x==1){
        return 1;
    }
    return x * factorial(x-1);
}

void get_all_permutatins_inner(
    int* digits, int digits_length, 
    int* rotation, int rotation_length,
    int** rotations, int* rotations_length)
{
    if (digits_length == 1) {
        rotation[rotation_length] = digits[0];
        rotations[*rotations_length] = malloc((rotation_length + 1) * sizeof(int));
        memcpy(rotations[*rotations_length], rotation, (rotation_length + 1) * sizeof(int));
        (*rotations_length)++;
        return;
    }

    for (int i = 0; i < digits_length; i++) {
        int* new_rotation = malloc((rotation_length + 1) * sizeof(int));
        memcpy(new_rotation, rotation, rotation_length * sizeof(int));
        new_rotation[rotation_length] = digits[i];

        int* new_digits = malloc((digits_length - 1) * sizeof(int));
        int index = 0;
        for (int j = 0; j < digits_length; j++) {
            if (j != i) {
                new_digits[index++] = digits[j];
            }
        } 

        get_all_permutatins_inner(new_digits, digits_length - 1,
            new_rotation, rotation_length + 1,
            rotations, rotations_length);

        free(new_rotation);
        free(new_digits);
    }
}

void get_all_permutations(int* digits_in, int digits_length_in, int*** rotations_out, int* rotations_length_out) {
    int** rotations = malloc(factorial(digits_length_in) * sizeof(int*));

    get_all_permutatins_inner(digits_in, digits_length_in, NULL, 0, rotations, rotations_length_out);


    *rotations_out = rotations;
}

void get_all_circulations(int* digits_in, int digits_length_in, int*** rotations_out, int* rotations_length_out) {
    int** rotations = malloc(digits_length_in * sizeof(int*));
    
    for (int i = 0; i < digits_length_in; i++) {
        int* rotation = malloc(digits_length_in * sizeof(int));
        for (int j = 0; j < digits_length_in; j++) {
            int index = (j + i) % digits_length_in;
            rotation[j] = digits_in[index];
        }
        rotations[i] = rotation;
    }

    *rotations_out = rotations;
    *rotations_length_out = digits_length_in;
}

int count_digits(int num) {
    int count = 0;

    if (num == 0)
        return 1;

    while (num != 0) {
        num /= 10;
        count++;
    }

    return count;
}

int intPow(int base, int exponent) {
    if (exponent==0){
        return 1;
    }
    int result = 1;
    while (exponent > 0) {
        result *= base;
        exponent--;
    }
    return result;
}

int array_to_number(int* arr, int arr_length){
    int result = 0;
    for(int i=arr_length-1;i>=0;i--){
        result += arr[i] * (intPow(10, arr_length-1-i));
    }
    return result;
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

bool is_circular_prime(int number){
    int digits_count = count_digits(number);
    int* digits = malloc(digits_count*sizeof(int));
    
    // Convert number to array of digits
    int i = digits_count-1;
    while (number != 0) {
        digits[i] = number % 10;    
        number = number / 10;
        i--;
    }
    
    int** rotations;
    int rotations_length;
    
    get_all_circulations(digits, digits_count, &rotations, &rotations_length);
    bool results = true;
    
    for (int i=0;i<rotations_length;i++){
        if (is_prime(array_to_number(rotations[i], digits_count)) != true){
            results = false;
            break; 
        }
    }

    free(digits);
    for (int i = 0; i < rotations_length; i++) {
        free(rotations[i]);
    }
    free(rotations);

    return results;
}

int main(int argc, char* argv[]) {
    clock_t start = clock();
    int count_circular_primes = 0;
    for(int i=0; i<1000000;i++){
       if (is_circular_prime(i)){
            count_circular_primes++;
       }
    } 
    clock_t stop = clock();
    double elapsed = (double)(stop-start)/CLOCKS_PER_SEC;
    /* long elapsed = (stop-start)CL */
    printf("number of circular primes<1000000 = %d (elapsed=%fs)", count_circular_primes, elapsed);
}
