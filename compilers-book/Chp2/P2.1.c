#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include <string.h>

struct conversion_for_calcuation {
    int convert_value;
    int convert_value_next;
    int base10;
    int digit;
};

struct roman_numeral {
    char name;
    int value;
};

struct conversion_for_calcuation calculate(int convert_value) {
    // next number to process
    int convert_value_next = floor(convert_value / 10);

    // the digited result
    int digit = convert_value - (floor(convert_value / 10) * 10);

    // the base10 part of result
    int base10 = convert_value - digit;


    return (struct conversion_for_calcuation){
        .convert_value = convert_value,
        .convert_value_next = convert_value_next,
        .base10 = base10,
        .digit = digit,
    };
}

struct roman_numeral romans[7] = {
    { .name = 'I', .value = 1 },
    { .name = 'V', .value = 5 },
    { .name = 'X', .value = 10 },
    { .name = 'L', .value = 50 },
    { .name = 'C', .value = 100 },
    { .name = 'D', .value = 500 },
    { .name = 'M', .value = 1000 }
};

int main() {
    int convert_value;

    printf("Integer value to convert to roman numerals: ");
    int input_result = scanf("%d", &convert_value);

    /*
    if (input_result == 1) {
        struct conversion_for_calcuation result = calculate(convert_value);

        while (result.convert_value > 0) {

            printf("base10 + digit = value to convert: %d + %d = %d\n", result.base10, result.digit, result.convert_value);

            result = calculate(result.convert_value_next);
        }
    }
    */

    if (input_result == 1) {
        printf("value converting = %d\n", convert_value);

        int numerical = -1;
        char name = '\0';
        int next_value = convert_value;
        int number_multiples = -1;
        int remainder = -1;
        bool previous_digit = false;
        char *result;
        char *temp;
        for (int i = 6; i >= 0; i--) {
            numerical = romans[i].value;
            name = romans[i].name;
            if (numerical > next_value) continue;

            remainder = (next_value % numerical);

            number_multiples = (next_value - remainder) / numerical;
            printf("number multiples = %d, numerical roman value = %d\n", number_multiples, numerical);

            next_value = next_value - (number_multiples * numerical);

            printf("next value = %d, previous digit: %b, roman numeral: %c\n", next_value, previous_digit, name);

            // if number_multiples == 1 could have a previous digit
            if (number_multiples == 1)
                previous_digit = true;
            else
                previous_digit = false;

            // keep track of the roman numeral string
            *temp = *result;
            char temp2;
            result = (char *) malloc((sizeof(temp) + 1) * sizeof(char));
            // copy temp back into result after allowcating new result
            for (int i = 0; i < strlen(temp) - 1; i++) {
                *result++ = temp[i];
                *result = '\0';
            }
            // cycle back through the result string to adjust for new roman numerals
            for (int i = 0; i < strlen(result) - 1; i++) {
                // NULL is at end of char string so 2
                if (i == strlen(result) - 2)
                {
                    if (previous_digit) {
                        temp2 = result[i];

                        --(*result);
                        *result = name;
                        *result++ = temp2;
                        *result = '\0';

                        break;
                    } else {
                        *result++ = name;
                        *result = '\0';
                    }
                }
            }
        }
        free(result);
    }

    return 0;
}
