#include <stdio.h>
#include <stdlib.h>
#include "utils.h"
#include "errors.h"

int_range_struct* int_range_alloc(int size) {
    int_range_struct* result = (int_range_struct*)malloc(sizeof(int_range_struct));
    if (result == NULL) {
        fprintf(stderr, "Memory allocation failed.\n");
        return NULL;
    }

    result->size = size;
    result->values = (int*)malloc(size * sizeof(int));
    if (result->values == NULL) {
        fprintf(stderr, "Memory allocation failed.\n");
        free(result);
        return NULL;
    }

    return result;
}

void free_int_range_result(int_range_struct* result) {
    if (result != NULL) {
        free(result->values);
        free(result);
    }
}

result int_range(int start, int stop, int step) {
    if (step == 0) {
        return error("Step value cannot be zero.");
    }

    if (start <= stop && step < 0) {
        return error("Step value must be positive for increasing range.");
    }

    if (start >= stop && step > 0) {
        return error("Step value must be negative for decreasing range.");
    }

    int array_size = (stop - start) / step;

    if (array_size <= 0) {
        return error("Invalid range: no elements in the range.");
    }

    int_range_struct* range_array = int_range_alloc(array_size);

    if (range_array->values == NULL) {
        return error("Memory allocation failed.");
    }

    int value = start;
    for (int i = 0; i < array_size; i++) {
        range_array->values[i] = value;
        value += step;
    }

    return ok(range_array);
}

int* multiply_by_each(int* multipliers, int multiplers_size, int multiplicand){
    int* product = malloc(multiplers_size * sizeof(int));
    
    for (int i=0; i<multiplers_size; i++){
        product[i] = multipliers[i] * multiplicand;
    }

    return product;
}

void print_int_array(int arr[], int size) {
    printf("[");
    for (int i = 0; i < size; i++) {
        printf("%d", arr[i]);
        if (i < size - 1) {
            printf(", ");
        }
    }
    printf("]\n");
}
