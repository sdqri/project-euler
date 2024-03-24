#include <stdlib.h>
#include <string.h>
#include "errors.h"

result ok(void* value) {
    result result;
    result.Ok = value;
    result.Err = NULL;
    return result;
}

result error(const char* message) {
    result result;
    result.Ok = NULL;
    result.Err = strdup(message);
    return result;
}

void free_result(result result) {
    if (result.Err != NULL) {
        free(result.Err);
    }
}
