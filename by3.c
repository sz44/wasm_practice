#include <stdio.h>
#include <emscripten/emscripten.h>

EMSCRIPTEN_KEEPALIVE int myFunction(int num) {
	printf("multiplying by 3...\n");
	return num * 3;
}

