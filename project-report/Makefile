PANDOC=scholdoc
LATEX=pdflatex
NAME=project-report
TEMPLATE=scholdoc-template.latex

all: journal-pdf

journal:
	$(PANDOC) -Ss -N --template=$(TEMPLATE) -o $(NAME).tex $(NAME).md

pdf:
	$(LATEX) $(NAME).tex

journal-pdf: journal pdf
