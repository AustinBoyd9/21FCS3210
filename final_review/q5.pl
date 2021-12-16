subs(_, _, [], []).
subs(X, Y, [X | T1], [Y | T2]) :- subs(X, Y, T1, T2).
subs(X, Y, [Z | T1], [Z | T2]) :- X \= Z, subs(X, Y, T1, T2).