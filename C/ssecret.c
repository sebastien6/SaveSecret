#include <stdio.h>
#include <stdlib.h>
#include "ssecret"

void WSaveSecret(GoString p0, GoString p1) { SaveSecret(p0, p1); }

char* WGetSecret(GoString p0, GoString p1, GoString p2) { 
	char* c;
	c = GetSecret(p0, p1, p2);
	return c;
}

int main () {}
