#ifndef HELLO_H
#define HELLO_H

const char* getHello();
char* return_string(const char *in);

typedef struct sampleStruct {
    char name[20];
    int number;
    void *data;
} sampleStruct;

struct sampleStruct get_struct();
struct sampleStruct create_sampleStruct(const char *name, int number, void *data);


#endif // HELLO_H