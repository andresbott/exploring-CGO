#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "hello.h"

const char* getHello() {
    return "hello world";
}

char* return_string(const char* input) {
    // Allocate memory for the new string
    char* result = malloc(strlen(input) + 1);

    // Check if memory allocation was successful
    if (result == NULL) {
        perror("Memory allocation failed");
        exit(EXIT_FAILURE);
    }

    // Copy the input string to the result
    strcpy(result, input);

    return result;
}

// Function to return a dynamically allocated array of strings
char** sample_string_array() {
    // Number of strings
    int numStrings = 3;

    // Allocate memory for the array of strings
    char** stringArray = (char**)malloc(numStrings * sizeof(char*));

    if (stringArray == NULL) {
        perror("Memory allocation failed");
        exit(EXIT_FAILURE);
    }

    // Initialize each string
    for (int i = 0; i < numStrings; ++i) {
        // Adjust the size based on your needs
        stringArray[i] = (char*)malloc(20 * sizeof(char));
        sprintf(stringArray[i], "String %d", i + 1);
    }

    return stringArray;
}



// this returns a populated instance of the struct
struct sampleStruct get_struct(){
    // Example usage
    const char *sampleName = "Sample";
    int sampleNumber = 42;
    void *sampleDataPtr = malloc(10);
    return create_sampleStruct(sampleName, sampleNumber, sampleDataPtr);
}
// this factory creates a struct
struct sampleStruct create_sampleStruct(const char *name, int number, void *data) {
    sampleStruct instance;

    strncpy(instance.name, name, sizeof(instance.name) - 1);
    instance.name[sizeof(instance.name) - 1] = '\0'; // Ensure null-termination

    instance.number = number;
    instance.data = data;
    return instance;
}