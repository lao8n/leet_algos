-module(quicksort).
-export([quicksort/1, lc_quicksort/1]).

%% basic quicksort function
quicksort([]) -> []; % base case 
quicksort([Pivot|Rest]) ->
  {Smaller, Larger} = partition(Pivot, Rest, [], []),
  quicksort(Smaller) ++ [Pivot] ++ quicksort(Larger).

partition(_, [], Smaller, Larger) -> {Smaller, Larger};
partition(Pivot, [H|T], Smaller, Larger) ->
  if H =< Pivot -> partition(Pivot, T, [H|Smaller], Larger);
     H >  Pivot -> partition(Pivot, T, Smaller, [H|Larger])
  end.

%% quicksort build with list comprehensions rather than partition function
lc_quicksort([]) -> [];
lc_quicksort([Pivot|Rest]) ->
  lc_quicksort([Smaller || Smaller <- Rest, Smaller =< Pivot]) % || means given 
  ++ [Pivot] ++
  lc_quicksort([Larger || Larger <- Rest, Larger > Pivot]).

%% you could also avoid processing values equal to the pivot more than once
%% we could make this more efficient by returning 3 lists: smaller, equal and larger
%% to the pivot, and concatenating them as partition. see below 
bestest_qsort([]) -> [];
bestest_qsort(L=[_|_]) ->
    bestest_qsort(L, []).

bestest_qsort([], Acc) -> Acc;
bestest_qsort([Pivot|Rest], Acc) ->
    bestest_partition(Pivot, Rest, {[], [Pivot], []}, Acc).

bestest_partition(_, [], {Smaller, Equal, Larger}, Acc) ->
    bestest_qsort(Smaller, Equal ++ bestest_qsort(Larger, Acc));
bestest_partition(Pivot, [H|T], {Smaller, Equal, Larger}, Acc) ->
    if H < Pivot ->
           bestest_partition(Pivot, T, {[H|Smaller], Equal, Larger}, Acc);
       H > Pivot ->
           bestest_partition(Pivot, T, {Smaller, Equal, [H|Larger]}, Acc);
       H == Pivot ->
           bestest_partition(Pivot, T, {Smaller, [H|Equal], Larger}, Acc)
    end.

%% to improve this further you would want to choose pivot more carefully
