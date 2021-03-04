# extractor
A (useless) binary to extract text between two boundary from stdin to stdout

## Install 
```
go get github.com/shoxxdj/extractor
```

## Man 

```
Usage of extractor:
  -start string
    	the first string (default "defaultValue_Th4tn3v3r_Ex1st")
  -start-like string
    	part the first string (default "defaultValue_Th4tn3v3r_Ex1st")
  -end string
    	the end string (default "defaultValue_Th4tn3v3r_Ex1st")
  -end-like string
    	part of end string (default "defaultValue_Th4tn3v3r_Ex1st")
  -continue
    	if not set, stop after first block

```

## Example 
```
>> cat beers.txt
Delirium
Chouffe
Triple karmeliet 
Effet Papillon
Azimut
K8
Mort subite
Paix Dieu
BBT
Woolf
Kanaha beer
Olympic white
Terko

>> cat beers.txt | extractor -start="Azimut" -end="Woolf"
Azimut
K8
Mort subite
Paix Dieu
BBT
Woolf
```

## Tips 

- Use ```-start-like``` and ```-end-like``` if you have an indented file.


Feel free to contribute 
