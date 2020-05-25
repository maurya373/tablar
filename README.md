# tablar

Tablar is a tool to generate random guitar/bass tablature given a key and a mode.

For example, if tablar is ran with key=G and mode=Phrygian, the following output is provided. (A count of 12 was defined, so thats the number of notes generated) 

If you aren't sure what modes are and how they relate to theory and tabs, I recommend this great [youtube video](https://www.youtube.com/watch?v=bwaeBUYcO5o) that covers it all in 15 minutes. Modes provide structure in terms of musical theory so we don't get horrible sounding notes in unison by pure randomness.

G Phrygian : G G# A# C D D# F

E-------------------------  
B---3-4-----6-------3-11--  
G-----------------1-------  
D-------------1-----------  
A-3-------3-----1---------  
E------1---------------10-  


The notes generated are truly random - the only constraint they have is being within the scale. What has been generated above is quite difficult to actually play, and would not sound good to a musician's ear, as its covering a wide range of octaves quickly.

Now, say the number of strings notes being generated on is limited down to the lowest 3 strings:

E------------------------  
B------------------------  
G------------------------  
D-6-------10---8---10-----  
A--10-8-------------------  
E------8-3---8---8---6-10-  

^ This one sounds heavy, trust me. **\m/** 

This is a lot more useful to spark ideas, as it is actually playable. Timing doesn't exist now, so the tempo of the notes is all customizable. 

If you want to unleash your inner Mark Morton, you should generate a riff using only the low string (oh and tune to Drop D).


There is quite some work to be done here for this to be actually useful, but this was mainly created in process of expirementing with [Go](https://golang.org).

Future considerations:
* constraints of how far the notes apart from each other
* variations of notes: bends, hammer-ons, pull-offs, slides, PALM MUTES
* multiple notes (chords) and chord progressions
