#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>

int* getProperDivisors(int x, int* count){
    int* divisors = (int*)malloc(x * sizeof(int));
    if (divisors == NULL){
        printf("Memory allocation failed.\n");
        exit(EXIT_FAILURE);
    }

    divisors[0] = 1;
    int divisor_count = 1;
    int root_of_x = (int)sqrt((double)x);
    for(int i=2; i<=root_of_x; i++) {
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

int aliquotSum(int x){
    int count;
    int* divisors = getProperDivisors(x, &count);

    int sum = 0;
    for(int i=0;i<count;i++){
        sum += divisors[i];
    }

    free(divisors);
    return sum;
}


typedef enum AliquotSumClass{
    Perfect, 
    Deficient, 
    Abundent
} AliquotSumClass;

AliquotSumClass GetAliquotSumClass(int x) {
    int sum = aliquotSum(x);
    if (x == sum){
        return Perfect;
    } else if (x < sum){
        return Abundent;
    } else {       
        return Deficient;
    } 
}

int main(int argc, char *argv[]){
    clock_t start, end;
    double elapsed_time;
    start = clock();

    int* abundentNumbers = (int*)malloc(sizeof(int)*28123);
    int abundentCount = 0;
    /* int* results = (int*)malloc(sizeof(int)*28); */
    /* int resultsCount = 0; */
    long long int sumOfNonsummableByAbundantNumbers = 0;

    for (int i=1; i<=28123; i++){
        if(GetAliquotSumClass(i)==Abundent){
            abundentNumbers[abundentCount] = i;
            abundentCount++;
        }

        int break_outer_loop = 0;
        for (int j=0;j<abundentCount;j++){
            int difference = i - abundentNumbers[j];
            for (int k=0; k<abundentCount;k++){
                if(difference==abundentNumbers[k]){
                    break_outer_loop = 1;
                    break;
                } 
            }
            if (break_outer_loop==1){
                break;
            }
        }
        if (break_outer_loop==0){
            sumOfNonsummableByAbundantNumbers += i;
            printf("i=%d\n", i);
        }
    }

    end = clock();
    elapsed_time = ((double)end-start)/CLOCKS_PER_SEC;
    printf(
            "sum of nonsummable by abundant numbers = %lld (elapsed time=%f)\n",
            sumOfNonsummableByAbundantNumbers,
            elapsed_time
    );
    return EXIT_SUCCESS; 
}
