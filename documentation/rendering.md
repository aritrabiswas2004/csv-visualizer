# Table Rendering

Table rendering is simple. There are 2 ways

1. **Text rendering**: This is what the program was built for, with
text being the MySQL-like output that is described everywhere.

```text
+---------------------+----------------+------+-----------+----------+
| Name                | Location       | Age  | Elevation | Good     |
+---------------------+----------------+------+-----------+----------+
| Taj Mahal           | Agra           | 500  | 40.56     | yes      |
| Pyramids of Giza    | Giza           | 4000 | 46.77     | no       |
| Colosseum           | Rome           | 5000 | 10.02     | bearable |
| Christ The Redeemer | Rio de Janeiro | 400  | 100.22    | great    |
| Machu Picchu        | Peru           | 1000 | 1932.22   | nice     |
+---------------------+----------------+------+-----------+----------+
```

2. **Markdown Rendering**: Renders into tables that markdown accepts. LaTeX compatibility is possibly coming soon.

```text
| Name                | Location       | Age  | Elevation | Good     |
|---------------------|----------------|------|-----------|----------|
| Taj Mahal           | Agra           | 500  | 40.56     | yes      |
| Pyramids of Giza    | Giza           | 4000 | 46.77     | no       |
| Colosseum           | Rome           | 5000 | 10.02     | bearable |
| Christ The Redeemer | Rio de Janeiro | 400  | 100.22    | great    |
| Machu Picchu        | Peru           | 1000 | 1932.22   | nice     |
```