#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>

int* getDivisors(int x, int* count){
    int* divisors = (int*)malloc(x * sizeof(int));
    if (divisors == NULL){
        printf("Memory allocation failed.\n");
        exit(EXIT_FAILURE);
    }

    int divisor_count = 0;
    int root_of_x = (int)sqrt((double)x);
    for(int i=1; i<=root_of_x; i++) {
        if (x%i == 0) {
            divisors[divisor_count] = i;
            divisor_count++;
            if (i != x / i) {
                divisors[divisor_count] = x/i;
                divisor_count++;
            }
        } 
    }
    
    *count = divisor_count;
    return divisors;
}

int sumOfProperDivisors(int x){
    int count;
    int* divisors = getDivisors(x, &count);

    int sum = 0;
    for(int i=0;i<count;i++){
        sum += divisors[i];
    }
    // The proper divisors of a positive integer N are those numbers, other than N itself, that divide N without remainder.
    sum = sum - x;

    free(divisors);
    return sum;
}

bool is_amicable(int x){
    int sum = sumOfProperDivisors(x);
    if (x!=sum && x == sumOfProperDivisors(sum)) {
        return true;
    }
    return false;
}

int main(int argc, char *argv[]){
    clock_t start, end;
    double elapsed_time;
    start = clock();
    
    int sum_of_amicable_numbers = 0;
    for (int i=1; i<10000; i++){
        if (is_amicable(i)) {
            sum_of_amicable_numbers += i;
        }
    }
    
    end = clock();
    elapsed_time = ((double)end-start)/CLOCKS_PER_SEC;
    printf(
            "sum of amicable numbers less than 10000 = %d (elapsed time=%f)\n", 
            sum_of_amicable_numbers, 
            elapsed_time);
    return EXIT_SUCCESS; 
}
