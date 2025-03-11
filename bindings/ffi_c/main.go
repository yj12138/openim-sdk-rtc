package main

/*
#include<stdio.h>
#include<stdlib.h>
#include<stdint.h>

typedef void (*CallBack)(char* dataPtr,int len);


CallBack eventCallBack;

void InvokeCallBack(CallBack cb,char* dataPtr,int len){
   if(cb == NULL){
       printf("IMSDK: not set message handler");
   }else{
       cb(dataPtr,len);
   }
}

void CPrint(char* dataPtr,int len){
    if (dataPtr == NULL || len <= 0) {
        printf("Invalid input\n");
        return;
    }

    printf("CPrint Output: ");
    for (int i = 0; i < len; i++) {
        printf(" %d ",(unsigned char)dataPtr[i]); // 逐个字符打印
    }
    printf("\n"); // 换行
}



*/
import "C"

func main() {
}
