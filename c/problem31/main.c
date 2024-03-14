#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <time.h>

const int COINS[8] = {1, 2, 5, 10, 20, 50, 100, 200};

void change(int n, int *coins, int coinsSize, int *result){
    if (n==0){
        (*result)++;
         return;
    } else if(n<0 || coinsSize==0){
        return;
    }

    for(int i=coinsSize-1; i>=0;i--){
        int* newCoins = (int*)malloc(i*sizeof(int));
        memcpy(newCoins, COINS, i*sizeof(int));
        int c = 1;
        while (true) {
            int newN = n-(c*coins[i]);
            if (newN<0){
                break;
            } 
            change(newN, newCoins, i, result);
            c++;
        }
    } 
}


int main(int argc, char *argv[]){
    clock_t start, end; 
    double elapsed;
    start = clock();
    int* newCoins = (int*)malloc(8*sizeof(int));
    memcpy(newCoins, COINS, 8*sizeof(int));
    int result = 0;
    change(200, newCoins, 8, &result);
    end = clock();
    elapsed = ((double)end-start)/CLOCKS_PER_SEC;
    printf("result=%d (elapsed=%fs)", result, elapsed);
    return EXIT_SUCCESS;
}
