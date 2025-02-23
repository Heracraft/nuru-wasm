export const defaultCode = `// Welcome to the Nuru playground
// Write your code here and click the play button to run it

// Kielelezo cha programu ya Fibonacci

urefu = namba(jaza("ingiza urefu wa mlolongo"))

fanya fibo = unda(x) {
	kama (x == 0) {
		rudisha 0;
	} au kama (x == 1) {
		rudisha 1;
	} sivyo {
		rudisha fibo(x - 1) + fibo(x - 2);
	}
}
andika(fibo(urefu));`;
