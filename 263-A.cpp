#include<iostream>
using namespace std;

int main(){
    int row,col,x;
    for(int i=0;i<5;++i){
        for(int j=0;j<5;++j){
            cin>>x;
            if(x){row=i;col=j;}
        }
    }
    cout<<abs(row-2)+abs(col-2)<<endl;
    return 0;
}