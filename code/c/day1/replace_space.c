#include<stdio.h>
#include <stdlib.h>
#include<assert.h>
#include<string.h>

void replace_blank(char str[], int length)
{
    if(str == NULL && length <= 0)
        return;
    
    int original_length = 0;  // 当前字符串长度
    int number_of_blank = 0;
    int i = 0;

    while(str[i] != '\0')
    {
         ++original_length;
         if(str[i] == ' ')
            ++number_of_blank;
        ++i;
    }

    int new_length = original_length + number_of_blank * 2;  // 替换空格后长度
    if(new_length > length)
        return;

    int index_of_original = original_length;
    int index_of_new = new_length;
    while(index_of_original >= 0 && index_of_new > index_of_original)
    {
        if(str[index_of_original] == ' ')
        {
            str[index_of_new--] = '0';
            str[index_of_new--] = '2';
            str[index_of_new--] = '%';
        }
        else
        {
            str[index_of_new--] = str[index_of_original];
        }
        --index_of_original;
    }
}

void test_0()
{
    char str[100] = "We Are Happy";
    replace_blank(str, 100);
    assert(strcmp(str, "We%20Are%20Happy") == 0);
}

void test_1()
{
    char str[100] = "We  Are Happy";
    replace_blank(str, 100);
    assert(strcmp(str, "We%20%20Are%20Happy") == 0);
}

void test_2()
{
    char str[100] = "We  Are Happy ";
    replace_blank(str, 100);
    assert(strcmp(str, "We%20%20Are%20Happy%20") == 0);
}

void test_3()
{
    char str[100] = " We  Are Happy ";
    replace_blank(str, 100);
    assert(strcmp(str, "%20We%20%20Are%20Happy%20") == 0);
}

void test_4()
{
    char str[100] = "WeAreHappy";
    replace_blank(str, 100);
    assert(strcmp(str, "WeAreHappy") == 0);
}


void run_all_test()
{
    test_0();
    test_1();
    test_2();
    test_3();
    test_4();
}

int main(int argc, char* argv[]) {
  run_all_test();
  return EXIT_SUCCESS;
}