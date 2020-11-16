# Invoke-WebRequest -Uri 'https://www.antlr.org/download/antlr-4.12.0-complete.jar' -OutFile antlr.jar
Start-Process java -ArgumentList '-jar', 'antlr.jar', '-Dlanguage=Go', '-visitor', '-o pkg/dbc/parser', 'grammar/DBC.g4'
Start-Process java -ArgumentList '-jar', 'antlr.jar', '-Dlanguage=Go', '-visitor', '-o pkg/a2l/parser', 'grammar/A2L.g4'
