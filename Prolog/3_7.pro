sumaCifara(N, SC) :- N < 10, SC is N, !.
sumaCifara(N, SC) :- C is N mod 10, N1 is N // 10, sumaCifara(N1, SC1),  SC is C + SC1.


brojCifara(N, BC) :- N < 10, BC is 1, !.
brojCifara(N, BC) :- N1 is N // 10, brojCifara(N1, BC1), BC is 1 + BC1.


vecaCifra(A, B, A) :- A > B.
vecaCifra(A, B, B) :- not A > B.


maxCifra(N, MC) :- C is N mod 10, N1 is N // 10, maxCifraPom(N1, C, MC), !.

maxCifraPom(N, T, MC) :- N < 10, vecaCifra(N, T, MC), !.
maxCifraPom(N, T, MC) :- C is N mod 10, vecaCifra(T, C, M), N1 is N // 10, maxCifraPom(N1, M, MC).