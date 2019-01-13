#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h> 
#include<assert.h>

bool find(int* matrix, int rows, int columns, int number)
{
    bool found = false;

    if(matrix != NULL && rows > 0 && columns > 0)
    {
        int row = 0;
        int column = columns -1;
        while(row < rows && column >=0)
        {
            if(matrix[row * columns + column] == number)
            {
                found = true;
                break;
            }
            else if(matrix[row * columns + column] > number) 
            {
                --column;
            }
            else
            {
                ++row;
            }
        }
    }
    return found;
}

void test_in_array()
{
    int matrix[][4] = {{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}};
    assert(find(matrix, 4, 4, 7));
}

void test_not_in_array()
{
    int matrix[][4] = {{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}};
    assert(!find(matrix, 4, 4, 5));
}

void test_find_min_num()
{
    int matrix[][4] = {{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}};
    assert(find(matrix, 4, 4, 1));
}

void test_find_max_num()
{
    int matrix[][4] = {{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}};
    assert(find(matrix, 4, 4, 15));
}

void run_all_test()
{
    test_in_array();
    test_not_in_array();
    test_find_min_num();
    test_find_max_num();
}

int main(int argc, char* argv[]) {
  run_all_test();
  return EXIT_SUCCESS;
}