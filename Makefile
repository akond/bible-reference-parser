antlr4=java -Xmx500M -cp "/var/www/html/antlr-4.7.2-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool

bible_base_listener.go Bible.tokens: Bible.g4
	$(antlr4) -Dlanguage=Go $<
	