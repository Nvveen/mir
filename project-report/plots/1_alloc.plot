set title "Allocations per set of requests"
set xlabel "Concurrent requests"
set ylabel "Allocations (in bytes)"
set terminal postscript eps color font 'Helvetica,10'
set output '1_alloc.eps'
plot "1_alloc.txt" title "" with lines
