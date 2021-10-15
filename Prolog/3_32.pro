%Kata, Lazar, Marko, Nevenka i Ognjen
%Filipović, Grbović, Hadžić, Ivanović i Janković

clan(X, [X|_]).
clan(X, [_|R]) :- clan(X, R).


deca(L) :- L = [d(_,_,2),d(_,_,3),d(_,_,4),d(_,_,5),d(_,_,6)],
    clan(d(lazar, jankovic, _),L),
    clan(d(kata, _, G1),L),
    clan(d(_, ivanovic, G2),L),
    clan(d(nevenka, _, G3),L),
    clan(d(_,filipovic, G4),L),
    clan(d(marko, _, G5),L),
    clan(d(ognjen,_,G6),L),
    clan(d(_, hadzic, G7),L),
    clan(d(_,grbovic,_), L),
    G1 =:= G2+1, G2 =:= G3+1, G4 =:= G5+3, G6 =:= 2*G7.


    