PANDOC=pandoc
LATEX=pdflatex
NAME=project-progress
TEMPLATE=pandoc-template.latex

all: journal-pdf

journal:
	$(PANDOC) -t beamer -Ss -N --template=pandoc-template-beamer.latex -o $(NAME).tex $(NAME).md

pdf:
	$(LATEX) $(NAME).tex

journal-pdf: journal pdf
