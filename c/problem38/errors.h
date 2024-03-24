#ifndef ERRORS_H
#define ERRORS_H

typedef struct {
    void* Ok;
    char* Err;
} result;

result ok(void* value);
result error(const char* message); 
void free_result(result result);

#endif //ERRORS_H
