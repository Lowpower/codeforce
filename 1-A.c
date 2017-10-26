#include <cstdio>
#include <iostream>
#include <climits>
using namespace std;

int main()
{
	int n, m, a;
	long long sum;
	scanf("%d %d %d",&n, &m, &a);
	sum = (n - (n / a) * a) > 0 ? (n / a)+ 1 : (n / a);
	sum += (((m - a) - ((m - a) / a) * a) > 0 ? (((m - a) / a) + 1) * sum : (m - a) / a * sum);
	cout << sum <<endl;
	return 0;
}
