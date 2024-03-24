#ifndef UTILS_H
#define UTILS_H

#include "errors.h"

typedef struct {
    int* values; 
    int size;   
} int_range_struct;

int_range_struct* int_range_alloc(int size);
void free_int_range_result(int_range_struct* result);
result int_range(int start, int stop, int step);
int* multiply_by_each(int* multipliers, int multiplers_size, int multiplicand);
void print_int_array(int arr[], int size);


#endif //UTILS_H
