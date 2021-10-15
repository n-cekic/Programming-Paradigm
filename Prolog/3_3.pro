op1(X, Y) :- X is Y. %x = y, ako je x promenljiva, x= := y ako nije promenljiva

op2(X, Y) :- X =:= Y. %poredjenje po jednakosti

abs(X, X) :- X > 0.
abs(X, Y) :- X < 0, Y is -X.

abs2(X, X) :- X > 0, !.