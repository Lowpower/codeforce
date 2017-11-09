#include <stdio.h>

int main() {
	int num;
	scanf("%d", &num);
	int taxis[5]= {0};
	int temp;
	for(int i = 0; i < num; i++) {
		scanf("%d", &temp);
		taxis[temp] += 1;
	}
	int sum = taxis[4];
	int third = taxis[3];
	taxis[1] -= third;
	sum += third;
	sum += int(taxis[2] / 2);
	if (taxis[2]%2 != 0) {
		sum += 1;
		if (taxis[1] > 0) {
			taxis[1] -= 2;
		}
	}
	if (taxis[1] > 0) {
		sum += int(taxis[1] / 4);
		if (taxis[1]%4 != 0) {
			sum += 1;
		}
	}
	printf("%d\n",sum);
	return 0;
}
