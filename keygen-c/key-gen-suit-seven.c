#include <stdio.h>
#include <time.h>

int main()
{
    clock_t clock0, clock1;         // clock_t is defined in <time.h> as int
    int c1, c2, c3, c4, c5, c6, c7; // test faces running from 0 to 12
    int k1, k2, k3, k4;             // candidate keys
    int bound, validKey;            // bound=max value of a key; validKey=boolean (1=valid); c=number of 4 key sets
    int i, j;                       // indexes running through sums
    int sums[50000], c;             // array of all possible sums of key[c1-7]

    int key[4] = {0, 0, 0, 0}; // keys
    bound = 38;                // empirical

    clock0 = clock();

    printf("searching keys up to %d\n", bound);

    for (k1 = 0; k1 <= bound; k1++)
    {
        for (k2 = k1; k2 <= bound; k2++)
        {
            for (k3 = k2; k3 <= bound; k3++)
            {
                for (k4 = k3; k4 <= bound; k4++)
                {
                    key[0] = k1;
                    key[1] = k2;
                    key[2] = k3;
                    key[3] = k4;

                    validKey = 1;
                    c = 0;
                    for (c1 = 0; c1 <= 3; c1++)
                    {
                        for (c2 = c1; c2 <= 3; c2++)
                        {
                            for (c3 = c2; c3 <= 3; c3++)
                            {
                                for (c4 = c3; c4 <= 3; c4++)
                                {
                                    for (c5 = c4; c5 <= 3; c5++)
                                    {
                                        for (c6 = c5; c6 <= 3; c6++)
                                        {
                                            for (c7 = c6; c7 <= 3; c7++)
                                            {
                                                sums[c] = key[c1] + key[c2] + key[c3] + key[c4] + key[c5] + key[c6] + key[c7];
                                                c++;
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
                                validKey = 0;
                            j++;
                        } while (validKey == 1 && j < c);
                        i++;
                    } while (validKey == 1 && i < c - 1);
                    if (validKey == 1)
                    {
                        clock1 = clock();
                        printf("\tkeys = %d %d %d %d - t=%.2f s\n", key[0], key[1], key[2], key[3], (float)(clock1 - clock0) / CLOCKS_PER_SEC);
                    }
                }
            }
        }
    }

    printf("done\n");
    return 0;
}

/*output
gcc -Wall -g -O3 key-gen-suit-seven.c -o key-gen-suit-seven-exec
./key-gen-suit-seven-exec
searching keys up to 38
	keys = 0 1 29 37 - t=0.00 s
	keys = 0 2 31 38 - t=0.00 s
	keys = 0 3 32 37 - t=0.01 s
	keys = 0 5 34 37 - t=0.01 s
	keys = 0 7 36 38 - t=0.01 s
	keys = 0 8 36 37 - t=0.02 s
	keys = 1 2 30 38 - t=0.04 s
	keys = 1 4 33 38 - t=0.04 s
	keys = 1 6 35 38 - t=0.04 s
	keys = 1 9 37 38 - t=0.05 s
done
*/
