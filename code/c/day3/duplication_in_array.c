#include<stdio.h>
#include<stdlib.h>

int find_duplication_number_in_array(int numbers[], int length)
{
    if (length <= 0)
        return -1;
    
    for(int i=0; i < length; ++i)
    {
        if(numbers[i] < 0 || numbers[i] > length - 1)
            return -1;
    }

    for(int i=0; i< length; ++i)
    {
        while(numbers[i] != i) 
        {
            if(numbers[i] == numbers[numbers[i]])
            {
                return numbers[i];
            }
            int temp = numbers[i];
            numbers[i] = numbers[temp];
            numbers[temp] = temp;
        }
    }

    return -1;
}

void test_0()
{
    int numbers[] = {2, 3, 1, 0, 5, 2, 3};
    printf("duplication number: %d\n", find_duplication_number_in_array(numbers, sizeof(numbers)/sizeof(numbers[0])));
}

void test_1()
{
    int numbers[] = {2, 3, 1, 0, 5, 6, 4};
    printf("duplication number: %d\n", find_duplication_number_in_array(numbers, sizeof(numbers)/sizeof(numbers[0])));
}

void test_2()
{
    int numbers[] = {1, 5, 6};
    printf("duplication number: %d\n", find_duplication_number_in_array(numbers, sizeof(numbers)/sizeof(numbers[0])));
}

void run_all_test()
{
    test_0();
    test_1();
    test_2();
}

int main(int argc, char* argv[]) {
  run_all_test();
  return EXIT_SUCCESS;
}