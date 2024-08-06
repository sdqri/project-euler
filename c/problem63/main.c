#include <stdbool.h>
#include <stdio.h>
#include <time.h>

#include "BigInt.h" /*https://github.com/dandclark/BigInt/tree/master*/

bool big_pow(BigInt *base, const int p) {
    BigInt *base_clone = BigInt_clone(base, base->num_digits);
    for (int i = 0; i < p - 1; i++) {
        BOOL flag = BigInt_multiply(base, base_clone);
        if (flag == 0) { /* failure case*/
            printf("failure with flag = %d\n", flag);
            BigInt_free(base_clone);
            return false;
        }
    }
    BigInt_free(base_clone);
    return true;
}

int main(int argc, char *argv[]) {
    clock_t start = clock();
    int count = 0;
    for (int i = 1; i < 10; i++) {
        int p = 1;
        bool flag = true;
        while (flag) {
            BigInt *base = BigInt_construct(i);
            big_pow(base, p);
            if (p > base->num_digits) {
                flag = false;
            } else if (p == base->num_digits) {
                count++;
            } 
            p++;
            BigInt_free(base);
        }
    }
    printf("number of n-digit positive integers which are also an nth power =  = %d", count);
    clock_t stop = clock();
    double elapsed = (double)(stop - start) / CLOCKS_PER_SEC;
    printf(" ( elapsed = %fs )\n", elapsed);
}
