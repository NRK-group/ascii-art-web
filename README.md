# ASCII-ART-WEB

## Authors

kkamil, Adriell, G.Orlandi

## USAGE

Execute "go run ." in the terminal
Visit "localhost:5050" on your browser
Type the text you would like to convert
To clear the text input use the backspace key or click the clear button
Select the font of your choice
Click the "Generate Art" button

## IMPLEMENTATION DETAILS

Server executes GET request when the app is launched
A form is implemented in HTML to pass the text and font as inputs for the POST request
The POST request generates the output using the AsciiArt function
The output is passed to the HTML file to be displayed as Ascii Art
