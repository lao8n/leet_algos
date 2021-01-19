-module(quicksort).
-export([quicksort/1, sort_w_listcomp/1, sort_w_acc/1]).

%% basic quicksort function
quicksort([]) -> []; % base case 
quicksort([Pivot|Tail]) ->
  {Lesser, Greater} = partition(Pivot, Tail, [], []),
  quicksort(Lesser) ++ [Pivot] ++ quicksort(Greater).

partition(_, [], Lesser, Greater) -> {Lesser, Greater};
partition(Pivot, [H|T], Lesser, Greater) ->
  if H =< Pivot -> partition(Pivot, T, [H|Lesser], Greater);
     H >  Pivot -> partition(Pivot, T, Lesser, [H|Greater])
  end.

%% quicksort build with list comprehensions rather than partition function
sort_w_listcomp([]) -> [];
sort_w_listcomp([Pivot|Tail]) ->
  sort_w_listcomp([Lesser || Lesser <- Tail, Lesser =< Pivot]) % || means given 
  ++ [Pivot] ++
  sort_w_listcomp([Greater || Greater <- Tail, Greater > Pivot]).

%% you could also avoid processing values equal to the pivot more than once
%% we could make this more efficient by returning 3 lists: Lesser, equal and Greater
%% to the pivot, and concatenating them as partition. see below 
sort_w_acc([]) -> [];
sort_w_acc(L=[_|_]) ->
    sort_w_acc(L, []).

sort_w_acc([], Acc) -> Acc;
sort_w_acc([Pivot|Tail], Acc) ->
    partition_w_acc(Pivot, Tail, {[], [Pivot], []}, Acc).

partition_w_acc(_, [], {Lesser, Equal, Greater}, Acc) ->
    sort_w_acc(Lesser, Equal ++ sort_w_acc(Greater, Acc));
partition_w_acc(Pivot, [H|T], {Lesser, Equal, Greater}, Acc) ->
    if H < Pivot ->
           partition_w_acc(Pivot, T, {[H|Lesser], Equal, Greater}, Acc);
       H > Pivot ->
           partition_w_acc(Pivot, T, {Lesser, Equal, [H|Greater]}, Acc);
       H == Pivot ->
           partition_w_acc(Pivot, T, {Lesser, [H|Equal], Greater}, Acc)
    end.

%% to improve this further you would want to choose pivot more carefully
%% cannot do in place because erlang is functional and only has immutable 
%% variables
