#include<stdlib.h>
#include<stdio.h>

// 计算区间内数字的多少
int count_range(int* numbers, int length, int start, int end)
{
    if(numbers == NULL)
        return 0;

    int count = 0;
    for(int i=0; i<length; i++)
    {
        if(numbers[i] >= start && numbers[i] <= end)
        {
            ++count;
        }
    }
    return count;
}

int find_duplication_number_in_array(int numbers[], int length)
{
    if (length <= 0)
        return -1;
    
    for(int i=0; i < length; ++i)
    {
        if(numbers[i] < 0 || numbers[i] > length - 1)
            return -1;
    }

    int start = 1;
    int end = length - 1;
    while(end >= start)
    {
        int middle = ((end - start) >> 1) + start;
        int count = count_range(numbers, length, start, middle);
        if(end == start)  // 区间只剩下一个数字
        {
            if(count > 1)
                return start;
            else
                break;
        }
        if(count > (middle - start + 1))
            end = middle;
        else
            start = middle + 1;
    }
    return -1;
}

void test_0()
{
    int numbers[] = {2, 3, 1, 6, 5, 2, 3};
    printf("duplication number: %d\n", find_duplication_number_in_array(numbers, sizeof(numbers)/sizeof(numbers[0])));
}

void test_1()
{
    int numbers[] = {2, 3, 1, 7, 5, 6, 4};
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

int main(int argc, char* argv[])
{
  run_all_test();
  return EXIT_SUCCESS;
}