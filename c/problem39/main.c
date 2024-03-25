#include <stdbool.h>
#include <stdio.h>
#include <time.h>

typedef struct {
    int a;
    int b;
    int c;
} pythogorean_triplet;

int power(int base, int exponent) {
    int result = 1;
    for (int i = 0; i < exponent; i++) {
        result *= base;
    }
    return result;
}

bool is_pythogorean_triplet_valid(pythogorean_triplet triplet){
    return power(triplet.a, 2) + power(triplet.b, 2) == power(triplet.c, 2);
}

int count_right_triangles(int primeter){
    if(primeter<12){
        return 0;
    }
    int count_triangles = 0;
    for (int hypotenuse=primeter-2;hypotenuse>4;hypotenuse--){
        int legs = primeter - hypotenuse;
        for(int a=1; a<legs/2; a++){
            int b = legs - a;
            pythogorean_triplet triplet;
            triplet.a = a;
            triplet.b = b;
            triplet.c = hypotenuse;
            if(is_pythogorean_triplet_valid(triplet)==true){
                count_triangles++;
            }
        }
    }
    
    return count_triangles;
}

int main(int argc, char* argv[]){
    clock_t start = clock();

    int max_p = 0;
    int max_solutions = 0;
    for (int p=12;p<=1000;p++){
        int number_of_solutions = count_right_triangles(p);
        if(number_of_solutions>max_solutions){
            max_p = p;
            max_solutions = number_of_solutions;
        }
    }

    clock_t stop = clock();
    double elapsed = (double)(stop-start)/CLOCKS_PER_SEC;
    printf("max soultions = %d for p = %d (elapsed=%fs)", max_solutions, max_p, elapsed);
}
