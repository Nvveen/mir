set title "Comparison of containers and storage"
set xlabel "Method"
set ylabel "Time (s)"
set boxwidth 0.5
set style fill solid
set terminal postscript eps color font 'Helvetica,10'
set output 'bar.eps'
plot "bar.txt" using 1:3:xtic(2) with boxes title ""
