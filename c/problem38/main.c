#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "errors.h"
#include "utils.h"

bool is_pandigital(int number){
    int digit_count[10] = {0};
    int total_count = 0;

    while(number > 0){
        int digit = number % 10;
        digit_count[digit]++;
        total_count++;
        number /= 10;
    }

    if (digit_count[0]>0){
        return false;
    }

    for (int i=1; i<=9; i++){
        if(digit_count[i]!=1){
            return false;
        }
    }
    
    return total_count == 9;
}

bool is_array_pandigital(int* numbers_array, int array_size){
    int digit_count[10] = {0};
    int total_count = 0;

    for(int i=0; i<array_size;i++){
        int number = numbers_array[i];
        while(number > 0){
            int digit = number % 10;
            digit_count[digit]++;
            total_count++;
            number /= 10;
        }
    }

    if (digit_count[0]>0){
        return false;
    }

    for (int i=1; i<=9; i++){
        if(digit_count[i]!=1){
            return false;
        }
    }
    
    return total_count == 9;
}

int digit_length(int number) {
    int length = 0;

    if (number < 0) {
        number = -number;
    }

    do {
        length++;
        number /= 10;
    } while (number != 0);

    return length;
}

int array_digit_length(int arr[], int size) {
    int cumulative_length = 0;

    for (int i = 0; i < size; i++) {
        cumulative_length += digit_length(arr[i]);
    }

    return cumulative_length;
}

int concatenate_array(int arr[], int size) {
    int concatenated = 0;
    for (int i = 0; i < size; i++) {
        concatenated *= pow(10, digit_length(arr[i]));
        concatenated += arr[i];
    }
    return concatenated;
}

int main(int argc, char* argv[]){
    clock_t start = clock();

    int max = 0;

    for(int j=3; j<=10;j++){
        for (int i=0;i<1000000000;i++){
            result range_result = int_range(1, j, 1);
            if(range_result.Err != NULL){
                fprintf(stderr, "Error: %s\n", range_result.Err);     
                goto fail;
            }
            int_range_struct* range_struct = range_result.Ok;
            int* products = multiply_by_each(range_struct->values, range_struct->size, i);
            bool break_flag = false;
            if (array_digit_length(products, range_struct->size) > 9){
                break_flag = true;
            }

            if(is_array_pandigital(products, range_struct->size)){
                int concatenated_pandigital = concatenate_array(products, range_struct->size);
                if (concatenated_pandigital > max){
                    max = concatenated_pandigital;
                }
            }
            
            free(products);
            free_int_range_result(range_struct);
            free_result(range_result);
            if (break_flag == true){
                break;
            }
        }
    }
    
    clock_t stop = clock();
    double elapsed = (double)(stop - start)/CLOCKS_PER_SEC;
    printf("maximum pandigital=%d (elapsed = %fs)\n", max, elapsed);
    return EXIT_SUCCESS;
fail:
    return EXIT_FAILURE;
}
