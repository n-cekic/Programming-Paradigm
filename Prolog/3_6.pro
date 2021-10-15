maksimum(A, B, M) :- A >= B, M is A.
maksimum(A, B, M) :- A < B, M is B.

maksimum2(A, B, A) :- A >= B, !.
maksimum2(A, B, B) :- not A >= B, !.

suma(1, 1) :- !.
suma(N, S) :- N > 0, M is N - 1, suma(M, Z), S is Z + N.

sumaParnih(2, 2) :- !.
sumaParnih(N, S) :- N > 2, N1 is N - 2, sumaParnih(N1, S1), S is S1 + N. 

proizvod(1, 1) :- !.
proizvod(N, P) :- N > 1, N1 is N - 1, proizvod(N1, P1), P is P1 * N.

proizvodNeparnih(1, 1) :- !.
proizvodNeparnih(N, P) :- N > 1, N1 is N - 2, proizvodNeparnih(N1, P1), P is P1 * N.

cifra(1,jedan).
cifra(2,dva).
cifra(3,tri).
cifra(4,cetiri).
cifra(5,pet).
cifra(6,sest).
cifra(7,sedam).
cifra(8,osam).
cifra(9,devet).
cifra(0,nula).

cifre(N) :- N < 1, !.
cifre(N) :- N < 10, cifra(N, X), write(X), nl, !.
cifre(N) :- N > 0, C is N mod 10,cifra(C, X), N1 is N // 10, cifre(N1), write(X).