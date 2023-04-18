package gptj

/*
// https://eli.thegreenplace.net/2019/passing-callbacks-and-pointers-to-cgo

extern int concatString(const char*);

int cgo_concatString(const char* s) {
	return concatString(s);
}
*/
import "C"
