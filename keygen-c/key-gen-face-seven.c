#include <stdio.h>
#include <time.h>

int main()
{
    clock_t clock0, clock1, clocki; // clock_t is defined in <time.h> as int
    int c1, c2, c3, c4, c5, c6, c7; // test faces running from 0 to 12
    int t, k, validKey;             // t=trial key, k=index searched key, validKey=boolean attribute to t (1=valid)
    int i, j;                       // indexes running through sums
    int sums[50000], sumKeys;       // array of all possible sums of key[c1-7]

    int key[13] = {0, 1, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}; // init keys - empirical
    k = 3;
    printf("bootstrap: keys = ");
    for (i = 0; i < k; i++)
        printf("%d ", key[i]);
    printf("\n");

    clock0 = clock();

    printf("searching keys from k =%d\n", k);
    while (k <= 12)
    { //------------------choose nb keys to search (<=12)
        t = key[k - 1] + 1;
        clocki = clock();
        do
        {
            key[k] = t;
            validKey = 1;
            sumKeys = 0;

            for (c1 = 0; c1 <= k; c1++)
            {
                for (c2 = c1; c2 <= k; c2++)
                {
                    for (c3 = c2; c3 <= k; c3++)
                    {
                        for (c4 = c3; c4 <= k; c4++)
                        {
                            for (c5 = c4; c5 <= k; c5++)
                            {
                                for (c6 = c5; c6 <= k; c6++)
                                {
                                    for (c7 = c6; c7 <= k; c7++)
                                    {
                                        if (c1 != c5 && c2 != c6 && c3 != c7)
                                        {
                                            sums[sumKeys] = key[c1] + key[c2] + key[c3] + key[c4] + key[c5] + key[c6] + key[c7];
                                            sumKeys++;
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }

            i = 0;
            do
            {
                j = i + 1;
                do
                {
                    if (sums[i] == sums[j])
                    {
                        validKey = 0;
                    }
                    j++;
                } while (validKey == 1 && j < sumKeys);
                i++;
            } while (validKey == 1 && i < sumKeys - 1);

            if (validKey == 1)
            {
                printf("\tkey[%d] = %d\n", k, t);
                clock1 = clock();
                printf("\t\trun time for key[%d] =\t%.2f s\n", k, (float)(clock1 - clocki) / CLOCKS_PER_SEC);
                printf("\t\trun time to key[%d] =\t%.2f s\n", k, (float)(clock1 - clock0) / CLOCKS_PER_SEC);
            }
            else
            {
                t++;
            }
        } while (validKey == 0);
        k++;
    }
    printf("done\n");
    return 0;
}

/*output
gcc -Wall -g -O3 key-gen-face-seven.c -o key-gen-face-seven-exec
./key-gen-face-seven-exec
bootstrap: keys = 0 1 5 
searching keys from k =3
	key[3] = 22
		run time for key[3] =	0.00 s
		run time to key[3] =	0.00 s
	key[4] = 98
		run time for key[4] =	0.00 s
		run time to key[4] =	0.00 s
	key[5] = 453
		run time for key[5] =	0.00 s
		run time to key[5] =	0.00 s
	key[6] = 2031
		run time for key[6] =	0.04 s
		run time to key[6] =	0.05 s
	key[7] = 8698
		run time for key[7] =	0.70 s
		run time to key[7] =	0.75 s
	key[8] = 22854
		run time for key[8] =	3.66 s
		run time to key[8] =	4.40 s
	key[9] = 83661
		run time for key[9] =	64.30 s
		run time to key[9] =	68.70 s
	key[10] = 262349
		run time for key[10] =	559.85 s
		run time to key[10] =	628.55 s
	key[11] = 636345
		run time for key[11] =	2493.26 s
		run time to key[11] =	3121.81 s
	key[12] = 1479181
		run time for key[12] =	12889.16 s
		run time to key[12] =	16010.97 s
done

