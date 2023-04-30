%Ejercicio 1: Subcadenas
member(E,[E|_]):-!.
member(E,[HEAD|TAIL]):-
    E \= HEAD,
    member(E,TAIL).

substring(SS,S):-
    split_string(S," "," ",SL),
    member(SS,SL).

subCadenas(_,[],[]).
subCadenas(SS,[H|T],F):-
    substring(SS,H),!,
    subCadenas(SS,T,F2),
    append(F2,[H],F),!.
subCadenas(SS,[_|T],F):-
    subCadenas(SS,T,F),!.
%---------------------------------

%Ejercicio 2: Subconjunto
sub_conjunto([],_):-!.
sub_conjunto([],[]):-!.
sub_conjunto([H1|T1],[H1|T2]):-
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
aplanar([],[]).
aplanar(X,[X]) :- not(is_list(X)).
aplanar([H|T],R) :-
    aplanar(H,R1),
    aplanar(T,R2),
    append(R1,R2,R),!.
%---------------------------------

%Ejercicio 4: Laberinto
conectado(i,2).
conectado(2,3).
conectado(2,8).
conectado(8,9).
conectado(9,3).
conectado(3,4).
conectado(4,10).
conectado(10,16).
conectado(16,22).
conectado(22,21).
conectado(21,15).
conectado(15,14).
conectado(14,13).
conectado(13,7).
conectado(7,1).
conectado(14,20).
conectado(20,26).
conectado(22,28).
conectado(26,27).
conectado(27,28).
conectado(28,29).
conectado(29,23).
conectado(23,17).
conectado(17,11).
conectado(11,5).
%conectado(1,7). %por eliminar
conectado(5,6).
conectado(28,34).
conectado(34,35).
conectado(35,36).
conectado(36,30).
conectado(30,24).
conectado(24,18).
conectado(18,12).
conectado(32,31).
conectado(31,25).
conectado(25,19).
conectado(34,33).
conectado(33,32).
conectado(32,f).

conectados(X,Y):- conectado(X,Y).
conectados(X,Y):- conectado(Y,X).

ruta(Fin,Fin,LRuta,Res):-
    reverse([Fin|LRuta],Res). %% reversa al resultado obtenido
ruta(Ini,Fin,LRuta,Res):-
    conectados(Ini,Z),
    not(member(Z,LRuta)),
    ruta(Z,Fin,[Ini|LRuta],Res).

% :- dynamic rutacorta/1.

rutas(Ini,Fin,LRutas):-
    findall(Ruta, ruta(Ini,Fin,[],Ruta), LRutas),!.
