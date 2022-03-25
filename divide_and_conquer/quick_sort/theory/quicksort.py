# standard copy approach for quicksort - not in place
def sort_w_listcomp(xs):
  if not xs:
    return []
  else:
    pivot = xs[0]
    lesser = [x for x in xs if x < pivot]
    greater = [x for x in xs[1:] if x >= pivot]
    return sort_w_listcomp(lesser) + [pivot] + sort_w_listcomp(greater)

# standard recursive in place implementation
# neat way to avoid needing a swap function with double assignment
def sort_recursive_in_place(xs):
  _sort_recursive_w_swap(xs, 0, len(xs) - 1)

def _sort_recursive_w_swap(xs, start, end):
  if end - start > 0:
    pivot, i, j = xs[start], start, end
    while i <= j:
      while xs[i] < pivot:
        i += 1
      while xs[j] > pivot:
        j -= 1
      if i <= j:
        # swap
        xs[i], xs[j] = xs[j], xs[i]
        i += 1
        j -= 1
    _sort_recursive_w_swap(xs, start, j)
    _sort_recursive_w_swap(xs, i, end)