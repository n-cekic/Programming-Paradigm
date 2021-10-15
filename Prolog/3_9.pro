ucenik(m1, milos, 1).
ucenik(i1, ivan, 2).
ucenik(m2, marko, 3).
ucenik(a1, ana, 1).
ucenik(m3, maja, 3).
ucenik(m4, mina, 2).

ocene(m1, srp, 5).
ocene(m2, srp, 4).
ocene(m3, srp, 4).
ocene(m4, srp, 3).
ocene(i1, srp, 4).
ocene(a1, srp, 5).
ocene(m1, eng, 3).
ocene(m2, eng, 4).
ocene(m3, eng, 5).
ocene(m4, eng, 5).
ocene(i1, eng, 4).
ocene(a1, eng, 3).
ocene(m1, mat, 3).
ocene(m2, mat, 3).
ocene(m3, mat, 2).
ocene(m4, mat, 4).
ocene(a1, mat, 5).
ocene(i1, mat, 4).

predmet(srp, srpski, 15).
predmet(eng, engleski, 15).
predmet(mat, matematika, 20).


barDvePeticeSifra(S) :- ocene(S, X, 5), ocene(S, Y, 5), X \= Y.

barDvePeticeIme(I) :- barDvePeticeSifra(S), ucenik(S, I, _).

barDvePeticePredmet(P) :- ocene(S1, P, 5), ocene(S2, P, 5), S1 \= S2.

% radi u oba smera: i za dati predmet daje odeljenje i za dato odeljenje daje predmet
odeljenjePetice(X, Y) :- ocene(S1, X, 5), ocene(S2, X, 5), S1 \= S2, ucenik(S1, _, Y), ucenik(S2, _, Y).