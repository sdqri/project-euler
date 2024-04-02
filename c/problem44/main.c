#include<math.h>
#include <stdbool.h>
#include <stdio.h>
#include <threads.h>
#include <time.h>

typedef unsigned long long int i64;

i64 P(i64 n) {
    return (n * (3*n -1))/2;
}

// 3n^2 - n - 2Pn = 0 => discriminant = 1 + 4 * 3 * 2Pn
bool is_pentagon(i64 Pn) {
    i64 discriminant = 1 + 24 * Pn;
    double d_root = sqrt((double)discriminant);
    i64 greater_root = ((1 + d_root) / 6);
    return (greater_root>0) && P(greater_root) == Pn; 
}


int main(int argc, char* argv[]){
    clock_t start = clock();

    i64 n = 1;
    while(1){ 
        i64 D = P(n);
    
        i64 d = 0;
        i64 i = 1;
        i64 k = i;
        bool found_answer = false;
    
        while(1){
            if(P(k)-P(i)==D && is_pentagon(P(k)+P(i))){
                printf("P%lld=%lld, P%lld=%lld, D=%lld sum=%lld ", i, P(i), k, P(k), D, P(k)+P(i));
                found_answer = true;
                break;
            }

            if (P(k)-P(i)>=D) {
                if (d==1){
                    found_answer = false;
                    break;
                }
                i++;
                k = i;
                d = 0;
            } else {
                d++;
                k = i + d;
            }

    
        }
    
        if(found_answer){
            break;
        }
        
        n++;
    }
    clock_t end = clock();
    double elapsed = (double)(end - start)/CLOCKS_PER_SEC;
    printf("(elapsed=%fs)\n", elapsed);
}
