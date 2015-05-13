set title "Comparison of allocations (100 requests)"
set xlabel "Method"
set ylabel "Allocations"
set boxwidth 0.5
set style fill solid
set terminal postscript eps color font 'Helvetica,10'
set output 'alloc.eps'
plot "alloc.txt" using 1:2:xtic(3) with boxes title ""
