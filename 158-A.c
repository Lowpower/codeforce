#include <iostream>
#include <cstdio>
#include <cstring>
using namespace std;
int main()
{
	int n,k,i,sum;
	int a[55];
	while(scanf("%d%d",&n,&k)!=EOF)
	{
		for(i=1; i<=n; i++)
			scanf("%d",&a[i]);
		int sum=0;
		if(a[k]!=0)
		{
			for(i=1; i<=n; i++)
			{
				if(a[i]>=a[k])
					sum++;
			}
		}
		else
		{
			for(i=1; i<=n; i++)
			{
				if(a[i]>a[k])
					sum++;
			}
		}
		printf("%d\n",sum);
	}
	return 0;
}
