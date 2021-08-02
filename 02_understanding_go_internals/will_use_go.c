#include <stdio.h>
#include "used_by_c.h"

int main(int argc, char **argv) {
    GoInt x = 12;
    GoInt y = 23;

    printf("About to call a Go function!\n");
    PrintMessage();

    GoInt p = Multiply(x,y);
    printf("Product: %d\n", (int)p);
    printf("It worked!\n");
    return 0;
}

//  gcc -o will_use_go will_use_go.c ./used_by_c.o