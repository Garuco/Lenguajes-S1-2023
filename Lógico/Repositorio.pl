
%Ejercicio 1: Subcadenas
substring(SS,S) :-
    string_to_list(SS, SSL),
    string_to_list(S, SL),
    append(_,T,SL),
    append(SSL,_,T),
    SS \= [],!.

subCadenas(_,[],_).
subCadenas(SS,[H|T],F):-
    substring(SS,H),
    append(F,[H],F2),
    subCadenas(SS,T,F2),!.
subCadenas(SS,[_|T],F):-
    subCadenas(SS,T,F),!.
%---------------------------------

%Ejercicio 2: Subconjunto
sub_conjunto([],_).
sub_conjunto([],[]).
sub_conjunto([H1|T1],[H2|T2]):-
    H1 = H2,
    sub_conjunto(T1,T2),!.
sub_conjunto(SC,[_|T2]):-
    sub_conjunto(SC,T2).

%Una manera mas facil.
%subtract(+Set, +Delete, -Result)
% SUBSTRACT borra todos los elementos de +Set que esten en +Delete,
% -Result es el resultado.
sub_conjunto2(SC,C) :-
    subtract(SC,C,[]),!.
%---------------------------------

%Ejercicio 3: Aplanar
aplanar([], []).
aplanar(X,[X]) :- not(is_list(X)).
aplanar([H|T],R) :-
    aplanar(H,R1),
    aplanar(T,R2),
    append(R1,R2,R),!.
%---------------------------------


