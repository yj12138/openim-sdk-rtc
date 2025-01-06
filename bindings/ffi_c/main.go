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
*/
import "C"

func main() {
}
