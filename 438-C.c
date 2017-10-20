#include <iostream>
#include <cstring>
using namespace std;

int main() {
	int n, k, temp;
	int vis[20];
	memset(vis, 0, sizeof(vis));
	cin>> n >> k;
	for(int i = 0; i < n; i++) {
		int ans = 0;
		for(int j = 0; j < k; j++){
			cin >> temp;
			ans = ans *2 + temp;
		}
		vis[ans]++;
	}
	for(int i = 0; i < (1<<k); i++) {
		for(int j = 0; j < (1<<k); j++) {
			if(vis[i] > 0 && vis[j] > 0 && (i&j) == 0) {
				cout<< "YES" << endl;
				return 0;
			}
		}
	}
	cout << "NO" << endl;
	return 0;
}
