# Simple Go Web Server

This project is a simple web server implemented in Go (Golang). It serves static resources such as images, CSS, and JavaScript from a designated directory, while HTML files are served dynamically. This structure allows for a clear separation of static and dynamic content.

## Features
- Static Content Handling: Serves static resources from a Resources folder.
- Dynamic Content Handling: Serves HTML content dynamically.
- Simple Setup: Minimal setup required, easy to start and test.

## Security Notice
I know you can probably tell but this server implementation is intended for educational purposes and not designed with security features suitable for a production environment. It lacks measures such as HTTPS, input sanitization, and authentication, which are critical for real-world applications. Use this server only in a safe, controlled setting. 

## Troubleshooting Rendering Issues
If CSS or JavaScript is not rendering correctly, make sure:

- Files are in the Resources directory.
- HTML references are correct (e.g., /Resources/filename.css).
- For additional help on rendering issues, refer to this  blog post that helped resolve similar challenges:
  https://dazspace.codes/blog/How+to+render+CSS+in+a+Golang+server-2021-October-10 
