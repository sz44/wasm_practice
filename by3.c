#include <stdio.h>
#include <emscripten/emscripten.h>

int main() {
	return 100;
}

EMSCRIPTEN_KEEPALIVE int myFunction(int argc, char ** argv) {
	if (argc > 1) {
		int num = atoi(argv[1]);
		printf("multiplying by 3...\n");
		return num * 3;
	} else {
		printf("no argument provided");
		return 0;
	}
}

