set title "Operation time"
set xlabel "Concurrent requests"
set ylabel "Time (s)"
set terminal postscript eps color font 'Helvetica,10'
set output '1.eps'
plot "1.txt" title "" with lines
