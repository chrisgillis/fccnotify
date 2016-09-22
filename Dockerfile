FROM golang:1.6-onbuild
CMD app -frn $frn -gmailaddr $gmailaddr -gmailpass $gmailpass
