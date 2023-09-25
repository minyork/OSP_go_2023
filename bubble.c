#include <stdio.h>

int main(){

    int a[] = {10,50,30,60,70};
    int i;
    int space;
    for(i=0; i<5 ;i++)
    {
        if(a[i] > a[i+1]){
            space = a[i];
            a[i] = a[i+1];
            a[i+1] = space;
        } 
    }

    printf(a);

    return 0;
}