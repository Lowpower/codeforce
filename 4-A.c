#include <cstdio>
int main()
{
	int n;
	scanf("%d",&n);
	puts((n%2==0 && n != 2)? "YES" : "NO");
	return 0;
}
